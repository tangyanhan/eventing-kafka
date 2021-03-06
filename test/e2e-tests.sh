#!/usr/bin/env bash

# Copyright 2020 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script runs the end-to-end tests against eventing-contrib built
# from source.

# If you already have the *_OVERRIDE environment variables set, call
# this script with the --run-tests arguments and it will use the cluster
# and run the tests.
# Note that local clusters often do not have the resources to run 12 parallel
# tests (the default) as the tests each tend to create their own namespaces and
# dispatchers.  For example, a local Docker cluster with 4 CPUs and 8 GB RAM will
# probably be able to handle 6 at maximum.  Be sure to adequately set the
# MAX_PARALLEL_TESTS variable before running this script, with the caveat that
# lowering it too much might make the tests run over the timeout that
# the go_test_e2e commands are using below.

# Calling this script without arguments will create a new cluster in
# project $PROJECT_ID, start Knative eventing system, install resources
# in eventing-contrib, run all of the test suites and delete the cluster.

# To run an individual integration test suite, set one of these environment variables:
#   TEST_CONSOLIDATED_CHANNEL=1    To install and run the consolidated channel
#   TEST_DISTRIBUTED_CHANNEL=1     To install and run the distributed channel

# Variables supported by this script:
#   PROJECT_ID:  the GKR project in which to create the new cluster (unless using "--run-tests")
#   MAX_PARALLEL_TESTS:  The maximum number of go tests to run in parallel (via "-test.parallel", default 12)

TEST_PARALLEL=${MAX_PARALLEL_TESTS:-12}

source $(dirname $0)/../vendor/knative.dev/hack/e2e-tests.sh

# If gcloud is not available make it a no-op, not an error.
which gcloud &> /dev/null || gcloud() { echo "[ignore-gcloud $*]" 1>&2; }

# Use GNU Tools on MacOS (Requires the 'grep' and 'gnu-sed' Homebrew formulae)
if [ "$(uname)" == "Darwin" ]; then
  sed=gsed
  grep=ggrep
fi

# Eventing main config path from HEAD.
readonly EVENTING_CONFIG="./config/"
readonly EVENTING_MT_CHANNEL_BROKER_CONFIG="./config/brokers/mt-channel-broker"
readonly EVENTING_IN_MEMORY_CHANNEL_CONFIG="./config/channels/in-memory-channel"

# Vendored Eventing Test Images.
readonly VENDOR_EVENTING_TEST_IMAGES="vendor/knative.dev/eventing/test/test_images/"
# HEAD eventing test images.
readonly HEAD_EVENTING_TEST_IMAGES="${GOPATH}/src/knative.dev/eventing/test/test_images/"

# Config tracing config.
readonly CONFIG_TRACING_CONFIG="test/config/config-tracing.yaml"

# Strimzi Kafka Cluster Brokers URL (base64 encoded value for k8s secret)
readonly STRIMZI_KAFKA_NAMESPACE="kafka" # Installation Namespace
readonly STRIMZI_KAFKA_CLUSTER_BROKERS="my-cluster-kafka-bootstrap.kafka.svc:9092"
readonly STRIMZI_KAFKA_CLUSTER_BROKERS_ENCODED="bXktY2x1c3Rlci1rYWZrYS1ib290c3RyYXAua2Fma2Euc3ZjOjkwOTI=" # Simple base64 encoding of STRIMZI_KAFKA_CLUSTER_BROKERS

# Eventing Kafka main config path from HEAD.
readonly KAFKA_CRD_CONFIG_TEMPLATE_DIR="./config/channel"
readonly DISTRIBUTED_TEMPLATE_DIR="${KAFKA_CRD_CONFIG_TEMPLATE_DIR}/distributed"
readonly CONSOLIDATED_TEMPLATE_DIR="${KAFKA_CRD_CONFIG_TEMPLATE_DIR}/consolidated"

# Eventing Kafka Channel CRD Secret (Will be modified with Strimzi Cluster Brokers - No Authentication)
readonly EVENTING_KAFKA_SECRET_TEMPLATE="300-kafka-secret.yaml"

# Eventing Kafka Channel CRD Config Map (Will be modified with SASL/TLS disabled)
readonly EVENTING_KAFKA_CONFIG_TEMPLATE="300-eventing-kafka-configmap.yaml"

# Strimzi installation config template used for starting up Kafka clusters.
readonly STRIMZI_INSTALLATION_CONFIG_TEMPLATE="test/config/100-strimzi-cluster-operator-0.20.0.yaml"
# Strimzi installation config.
readonly STRIMZI_INSTALLATION_CONFIG="$(mktemp)"
# Kafka cluster CR config file.
readonly KAFKA_INSTALLATION_CONFIG="test/config/100-kafka-ephemeral-triple-2.6.0.yaml"
# Kafka TLS ConfigMap.
readonly KAFKA_TLS_CONFIG="test/config/config-kafka-tls.yaml"
# Kafka SASL ConfigMap.
readonly KAFKA_SASL_CONFIG="test/config/config-kafka-sasl.yaml"
# Kafka Users CR config file.
readonly KAFKA_USERS_CONFIG="test/config/100-strimzi-users-0.20.0.yaml"
# Kafka PLAIN cluster URL
readonly KAFKA_PLAIN_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9092"
# Kafka TLS cluster URL
readonly KAFKA_TLS_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9093"
# Kafka SASL cluster URL
readonly KAFKA_SASL_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9094"
# Kafka cluster URL for our installation, during tests
KAFKA_CLUSTER_URL=${KAFKA_PLAIN_CLUSTER_URL}
# Kafka channel CRD config template file. It needs to be modified to be the real config file.
readonly KAFKA_CRD_CONFIG_TEMPLATE="400-kafka-config.yaml"
# Real Kafka channel CRD config , generated from the template directory and modified template file.
readonly KAFKA_CRD_CONFIG_DIR="$(mktemp -d)"
# Kafka channel CRD config template directory.
readonly KAFKA_SOURCE_CRD_CONFIG_DIR="config/source"

# Namespace where we install Eventing components
readonly SYSTEM_NAMESPACE="knative-eventing"

# Zipkin setup
readonly KNATIVE_EVENTING_MONITORING_YAML="test/config/monitoring.yaml"

#
# TODO - Consider adding this function to the test-infra library.sh utilities ?
#
# Add The kn-eventing-test-pull-secret To Specified ServiceAccount & Restart Pods
#
# If the default namespace contains a Secret named 'kn-eventing-test-pull-secret',
# then copy it into the specified Namespace and add it to the specified ServiceAccount,
# and restart the specified Pods.
#
# This utility function exists to support local cluster testing with a private Docker
# repository, and is based on the CopySecret() functionality in eventing/pkg/utils.
#
function add_kn_eventing_test_pull_secret() {

  # Local Constants
  local secret="kn-eventing-test-pull-secret"

  # Get The Function Arguments
  local namespace="$1"
  local account="$2"
  local deployment="$3"

  # If The Eventing Test Pull Secret Is Present & The Namespace Was Specified
  if [[ $(kubectl get secret $secret -n default --ignore-not-found --no-headers=true | wc -l) -eq 1 && -n "$namespace" ]]; then

      # If The Secret Is Not Already In The Specified Namespace Then Copy It In
      if [[ $(kubectl get secret $secret -n "$namespace" --ignore-not-found --no-headers=true | wc -l) -lt 1 ]]; then
        kubectl get secret $secret -n default -o yaml | sed "s/namespace: default/namespace: $namespace/" | kubectl create -f -
      fi

      # If Specified Then Patch The ServiceAccount To Include The Image Pull Secret
      if [[ -n "$account" ]]; then
        kubectl patch serviceaccount -n "$namespace" "$account" -p "{\"imagePullSecrets\": [{\"name\": \"$secret\"}]}"
      fi

      # If Specified Then Restart The Pods Of The Deployment
      if [[ -n "$deployment" ]]; then
        kubectl rollout restart -n "$namespace" deployment "$deployment"
      fi
  fi
}

function knative_setup() {
  if is_release_branch; then
    echo ">> Install Knative Eventing from ${KNATIVE_EVENTING_RELEASE}"
    kubectl apply -f ${KNATIVE_EVENTING_RELEASE}
  else
    echo ">> Install Knative Eventing from HEAD"
    pushd .
    cd ${GOPATH} && mkdir -p src/knative.dev && cd src/knative.dev
    git clone https://github.com/knative/eventing
    cd eventing
    ko apply -f "${EVENTING_CONFIG}"
    # Install MT Channel Based Broker
    ko apply -f "${EVENTING_MT_CHANNEL_BROKER_CONFIG}"
    # Install IMC
    ko apply -f "${EVENTING_IN_MEMORY_CHANNEL_CONFIG}"
    popd
  fi
  wait_until_pods_running knative-eventing || fail_test "Knative Eventing did not come up"

  install_zipkin
}

# Setup zipkin
function install_zipkin() {
  echo "Installing Zipkin..."
  kubectl apply -f "${KNATIVE_EVENTING_MONITORING_YAML}"
  wait_until_pods_running knative-eventing || fail_test "Zipkin inside eventing did not come up"
  # Setup config tracing for tracing tests
  kubectl apply -f "${CONFIG_TRACING_CONFIG}"
}

# Remove zipkin
function uninstall_zipkin() {
  echo "Uninstalling Zipkin..."
  kubectl delete -f "${KNATIVE_EVENTING_MONITORING_YAML}"
  wait_until_object_does_not_exist deployment zipkin knative-eventing || fail_test "Zipkin deployment was unable to be deleted"
  kubectl delete -n knative-eventing configmap config-tracing
}

function knative_teardown() {
  echo ">> Stopping Knative Eventing"
  if is_release_branch; then
    echo ">> Uninstalling Knative Eventing from ${KNATIVE_EVENTING_RELEASE}"
    kubectl delete -f "${KNATIVE_EVENTING_RELEASE}"
  else
    echo ">> Uninstalling Knative Eventing from HEAD"
    pushd .
    cd ${GOPATH}/src/knative.dev/eventing
    # Remove IMC
    ko delete -f "${EVENTING_IN_MEMORY_CHANNEL_CONFIG}"
    # Remove MT Channel Based Broker
    ko delete -f "${EVENTING_MT_CHANNEL_BROKER_CONFIG}"
    # Remove eventing
    ko delete -f "${EVENTING_CONFIG}"
    popd
  fi
  wait_until_object_does_not_exist namespaces knative-eventing
}

# Add function call to trap
# Parameters: $1 - Function to call
#             $2...$n - Signals for trap
function add_trap() {
  local cmd=$1
  shift
  for trap_signal in $@; do
    local current_trap="$(trap -p $trap_signal | cut -d\' -f2)"
    local new_cmd="($cmd)"
    [[ -n "${current_trap}" ]] && new_cmd="${current_trap};${new_cmd}"
    trap -- "${new_cmd}" $trap_signal
  done
}

function test_setup() {
  kafka_setup || return 1

  # Install kail if needed.
  if ! which kail > /dev/null; then
    bash <( curl -sfL https://raw.githubusercontent.com/boz/kail/master/godownloader.sh) -b "$GOPATH/bin"
  fi

  # Capture all logs.
  kail > "${ARTIFACTS}/k8s.log.txt" &
  local kail_pid=$!
  # Clean up kail so it doesn't interfere with job shutting down
  add_trap "kill $kail_pid || true" EXIT

  # Publish test images.
  echo ">> Publishing test images from eventing"
  # We vendor test image code from eventing, in order to use ko to resolve them into Docker images, the
  # path has to be a GOPATH.  The two slashes at the beginning are to anchor the match so that running the test
  # twice doesn't re-parse the yaml and cause errors.
  sed -i 's@//knative.dev/eventing/test/test_images@//knative.dev/eventing-kafka/vendor/knative.dev/eventing/test/test_images@g' "${VENDOR_EVENTING_TEST_IMAGES}"*/*.yaml
  $(dirname $0)/upload-test-images.sh "${VENDOR_EVENTING_TEST_IMAGES}" e2e || fail_test "Error uploading test images"
  $(dirname $0)/upload-test-images.sh "test/test_images" e2e || fail_test "Error uploading test images"
}

function test_teardown() {
  kafka_teardown
}

function install_consolidated_channel_crds() {
  echo "Installing consolidated Kafka Channel CRD"
  rm "${KAFKA_CRD_CONFIG_DIR}/"*yaml
  cp "${CONSOLIDATED_TEMPLATE_DIR}/"*yaml "${KAFKA_CRD_CONFIG_DIR}"
  sed -i "s/REPLACE_WITH_CLUSTER_URL/${KAFKA_CLUSTER_URL}/" ${KAFKA_CRD_CONFIG_DIR}/${KAFKA_CRD_CONFIG_TEMPLATE}
  ko apply -f "${KAFKA_CRD_CONFIG_DIR}" || return 1
  wait_until_pods_running knative-eventing || fail_test "Failed to install the consolidated Kafka Channel CRD"
}

function install_consolidated_sources_crds() {
  echo "Installing consolidated Kafka Source CRD"
  ko apply -f "${KAFKA_SOURCE_CRD_CONFIG_DIR}" || return 1
  wait_until_pods_running knative-eventing || fail_test "Failed to install the consolidated Kafka Source CRD"
}

# Uninstall The eventing-kafka KafkaChannel Implementation Via Ko
function uninstall_channel_crds() {
  echo "Uninstalling Kafka Channel CRD"
  kubectl delete secret -n knative-eventing kafka-cluster
  sleep 10 # Give Controller Time To React To Kafka Secret Deletion ; )
  echo "Current namespaces:"
  kubectl get namespaces
  echo "Current kafkachannels:"
  kubectl get kafkachannel -A
  ko delete --ignore-not-found=true --now --timeout 120s -f "${KAFKA_CRD_CONFIG_DIR}"
}

function uninstall_sources_crds() {
  echo "Uninstalling Kafka Source CRD"
  ko delete --ignore-not-found=true --now --timeout 120s -f "${KAFKA_SOURCE_CRD_CONFIG_DIR}"
}

function install_distributed_channel_crds() {
  echo "Installing distributed Kafka Channel CRD"
  rm "${KAFKA_CRD_CONFIG_DIR}/"*yaml
  cp "${DISTRIBUTED_TEMPLATE_DIR}/"*yaml "${KAFKA_CRD_CONFIG_DIR}"

  # Update The Kafka Secret With Strimzi Kafka Cluster Brokers (No Authentication)
  sed -i "s/brokers: \"\"/brokers: ${STRIMZI_KAFKA_CLUSTER_BROKERS_ENCODED}/" "${KAFKA_CRD_CONFIG_DIR}/${EVENTING_KAFKA_SECRET_TEMPLATE}"

  # Update the config-kafka configmap to disable SASL/TLS for the tests
  sed -i '/^ *TLS:/{n;s/true/false/};/^ *SASL:/{n;s/true/false/}' "${KAFKA_CRD_CONFIG_DIR}/${EVENTING_KAFKA_CONFIG_TEMPLATE}"

  # Install The eventing-kafka KafkaChannel Implementation
  ko apply -f "${KAFKA_CRD_CONFIG_DIR}" || return 1

   # Add The kn-eventing-test-pull-secret (If Present) To ServiceAccount & Restart eventing-kafka Deployment
  add_kn_eventing_test_pull_secret knative-eventing eventing-kafka-channel-controller eventing-kafka-channel-controller

  wait_until_pods_running knative-eventing || fail_test "Failed to install the distributed Kafka Channel CRD"
}

function kafka_setup() {
  # Create The Namespace Where Strimzi Kafka Will Be Installed
  echo "Installing Kafka Cluster"
  kubectl get -o name namespace ${STRIMZI_KAFKA_NAMESPACE} || kubectl create namespace ${STRIMZI_KAFKA_NAMESPACE} || return 1

  # Install Strimzi Into The Desired Namespace (Dynamically Changing The Namespace)
  sed "s/namespace: .*/namespace: ${STRIMZI_KAFKA_NAMESPACE}/" ${STRIMZI_INSTALLATION_CONFIG_TEMPLATE} > "${STRIMZI_INSTALLATION_CONFIG}"

  # Create The Actual Kafka Cluster Instance For The Cluster Operator To Setup
  kubectl apply -f "${STRIMZI_INSTALLATION_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"
  kubectl apply -f "${KAFKA_INSTALLATION_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"

  # Delay Pod Running Check Until All Pods Are Created To Prevent Race Condition (Strimzi Kafka Instance Can Take A Bit To Spin Up)
  local iterations=0
  local progress="Waiting for Kafka Pods to be created..."
  while [[ $(kubectl get pods --no-headers=true -n ${STRIMZI_KAFKA_NAMESPACE} | wc -l) -lt 6 && $iterations -lt 60 ]] # 1 ClusterOperator, 3 Zookeeper, 1 Kafka, 1 EntityOperator
  do
    echo -ne "${progress}\r"
    progress="${progress}."
    iterations=$((iterations + 1))
    sleep 3
  done
  echo "${progress}"

  # Wait For The Strimzi Kafka Cluster Operator To Be Ready (Forcing Delay To Ensure CRDs Are Installed To Prevent Race Condition)
  wait_until_pods_running "${STRIMZI_KAFKA_NAMESPACE}" || fail_test "Failed to start up a Strimzi Kafka Instance"

  # Create some Strimzi Kafka Users
  kubectl apply -f "${KAFKA_USERS_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"
}

function kafka_teardown() {
  echo "Uninstalling Kafka cluster"
  kubectl delete -f ${KAFKA_INSTALLATION_CONFIG} -n "${STRIMZI_KAFKA_NAMESPACE}"
  kubectl delete -f "${STRIMZI_INSTALLATION_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"
  kubectl delete namespace "${STRIMZI_KAFKA_NAMESPACE}"
}

function create_tls_secrets() {
  echo "Creating TLS Kafka secret"
  STRIMZI_CRT=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-cluster-cluster-ca-cert --template='{{index .data "ca.crt"}}' | base64 --decode )
  TLSUSER_CRT=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-tls-user --template='{{index .data "user.crt"}}' | base64 --decode )
  TLSUSER_KEY=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-tls-user --template='{{index .data "user.key"}}' | base64 --decode )

  kubectl create secret --namespace knative-eventing generic strimzi-tls-secret \
    --from-literal=ca.crt="$STRIMZI_CRT" \
    --from-literal=user.crt="$TLSUSER_CRT" \
    --from-literal=user.key="$TLSUSER_KEY"
}

function create_sasl_secrets() {
  echo "Creating SASL Kafka secret"
  STRIMZI_CRT=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-cluster-cluster-ca-cert --template='{{index .data "ca.crt"}}' | base64 --decode )
  SASL_PASSWD=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-sasl-user --template='{{index .data "password"}}' | base64 --decode )

  kubectl create secret --namespace knative-eventing generic strimzi-sasl-secret \
    --from-literal=ca.crt="$STRIMZI_CRT" \
    --from-literal=password="$SASL_PASSWD" \
    --from-literal=saslType="SCRAM-SHA-512" \
    --from-literal=user="my-sasl-user"
}

# Installs the resources necessary to test the consolidated channel, runs those tests, and then cleans up those resources
function test_consolidated_channel_plain() {
  # Test the consolidated channel with no auth
  echo "Testing the consolidated channel and source"
  install_consolidated_channel_crds || return 1
  install_consolidated_sources_crds || return 1

  go_test_e2e -tags=e2e,source -timeout=40m -test.parallel=${TEST_PARALLEL} ./test/e2e -channels=messaging.knative.dev/v1alpha1:KafkaChannel,messaging.knative.dev/v1beta1:KafkaChannel  || fail_test
  go_test_e2e -tags=e2e,source -timeout=5m -test.parallel=${TEST_PARALLEL} ./test/conformance -channels=messaging.knative.dev/v1beta1:KafkaChannel -sources=sources.knative.dev/v1beta1:KafkaSource || fail_test

  uninstall_sources_crds || return 1
  uninstall_channel_crds || return 1
}

function test_consolidated_channel_tls() {
  # Test the consolidated channel with TLS
  echo "Testing the consolidated channel with TLS"
  # Set the URL to the TLS listeners config
  cp ${KAFKA_TLS_CONFIG} "${CONSOLIDATED_TEMPLATE_DIR}/configmaps/kafka-config.yaml"
  KAFKA_CLUSTER_URL=${KAFKA_TLS_CLUSTER_URL}

  install_consolidated_channel_crds || return 1

  go_test_e2e -tags=e2e -timeout=40m -test.parallel=${TEST_PARALLEL} ./test/e2e -channels=messaging.knative.dev/v1beta1:KafkaChannel  || fail_test

  uninstall_channel_crds || return 1
}

function test_consolidated_channel_sasl() {
  # Test the consolidated channel with SASL
  echo "Testing the consolidated channel with SASL"
  # Set the URL to the SASL listeners config
  cp ${KAFKA_SASL_CONFIG} "${CONSOLIDATED_TEMPLATE_DIR}/configmaps/kafka-config.yaml"
  KAFKA_CLUSTER_URL=${KAFKA_SASL_CLUSTER_URL}

  install_consolidated_channel_crds || return 1

  go_test_e2e -tags=e2e -timeout=40m -test.parallel=${TEST_PARALLEL} ./test/e2e -channels=messaging.knative.dev/v1beta1:KafkaChannel  || fail_test

  uninstall_channel_crds || return 1
}

# Installs the resources necessary to test the distributed channel, runs those tests, and then cleans up those resources
function test_distributed_channel() {
  # Test the distributed channel
  echo "Testing the distributed channel"
  install_distributed_channel_crds || return 1

  # TODO: Enable v1alpha1 testing once we have the auto-converting webhook in our config/yaml
  go_test_e2e -tags=e2e -timeout=40m -test.parallel=${TEST_PARALLEL} ./test/e2e -channels=messaging.knative.dev/v1beta1:KafkaChannel  || fail_test
  go_test_e2e -tags=e2e -timeout=5m -test.parallel=${TEST_PARALLEL} ./test/conformance -channels=messaging.knative.dev/v1beta1:KafkaChannel || fail_test

  uninstall_channel_crds || return 1
}

function parse_flags() {
  # This function will be called repeatedly by initialize() with one fewer
  # argument each time and expects a return value of "the number of arguments to skip"
  # so we can just check the first argument and return 1 (to have it redirected to the
  # test container) or 0 (to have initialize() parse it normally).
  case $1 in
    --distributed)
      TEST_DISTRIBUTED_CHANNEL=1
      return 1
      ;;
    --consolidated)
      TEST_CONSOLIDATED_CHANNEL=1
      return 1
      ;;
    --consolidated-tls)
      TEST_CONSOLIDATED_CHANNEL_TLS=1
      return 1
      ;;
    --consolidated-sasl)
      TEST_CONSOLIDATED_CHANNEL_SASL=1
      return 1
      ;;
  esac
  return 0
}

TEST_CONSOLIDATED_CHANNEL=${TEST_CONSOLIDATED_CHANNEL:-0}
TEST_CONSOLIDATED_CHANNEL_TLS=${TEST_CONSOLIDATED_CHANNEL_TLS:-0}
TEST_CONSOLIDATED_CHANNEL_SASL=${TEST_CONSOLIDATED_CHANNEL_SASL:-0}
TEST_DISTRIBUTED_CHANNEL=${TEST_DISTRIBUTED_CHANNEL:-0}

echo "e2e-tests.sh command line: $@"

# Note:  The setting of gcp-project-id option here has no effect when testing locally; it is only for the kubetest2 utility
# If you wish to use this script just as test setup, *without* teardown, add "--skip-teardowns" to the initialize command
initialize $@ --skip-istio-addon

echo "e2e-tests.sh environment:"
echo "TEST_CONSOLIDATED_CHANNEL: ${TEST_CONSOLIDATED_CHANNEL}"
echo "TEST_DISTRIBUTED_CHANNEL: ${TEST_DISTRIBUTED_CHANNEL}"
echo "TEST_CONSOLIDATED_CHANNEL_TLS: ${TEST_CONSOLIDATED_CHANNEL_TLS}"
echo "TEST_CONSOLIDATED_CHANNEL_SASL: ${TEST_CONSOLIDATED_CHANNEL_SASL}"

# If neither of the four test was specified, run both plain tests
if [[ $TEST_CONSOLIDATED_CHANNEL != 1 ]] && [[ $TEST_CONSOLIDATED_CHANNEL_TLS != 1 ]] && [[ $TEST_CONSOLIDATED_CHANNEL_SASL != 1 ]] && [[ $TEST_DISTRIBUTED_CHANNEL != 1 ]]; then
  TEST_DISTRIBUTED_CHANNEL=1
  TEST_CONSOLIDATED_CHANNEL=1
fi

export SYSTEM_NAMESPACE

if [[ $TEST_CONSOLIDATED_CHANNEL == 1 ]]; then
  echo "Launching the PLAIN TESTS:"
  test_consolidated_channel_plain || exit 1
fi

if [[ $TEST_CONSOLIDATED_CHANNEL_TLS == 1 ]]; then
  echo "Launching the TLS TESTS:"
  create_tls_secrets
  test_consolidated_channel_tls || exit 1
fi

if [[ $TEST_CONSOLIDATED_CHANNEL_SASL == 1 ]]; then
  echo "Launching the SASL TESTS:"
  create_sasl_secrets
  test_consolidated_channel_sasl || exit 1
fi

# Terminate any zipkin port-forward processes that are still present on the system
ps -e -o pid,command | grep 'kubectl port-forward zipkin[^*]*9411:9411 -n' | sed 's/^ *\([0-9][0-9]*\) .*/\1/' | xargs kill

if [[ $TEST_DISTRIBUTED_CHANNEL == 1 ]]; then
  test_distributed_channel || exit 1
fi

# If you wish to use this script just as test setup, *without* teardown, just uncomment this line and comment all go_test_e2e commands
# trap - SIGINT SIGQUIT SIGTSTP EXIT

success
