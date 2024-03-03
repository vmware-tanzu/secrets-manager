#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

go run ./ci/test/main.go ./ci/test/run.go

## Enable strict error checking.
#set -euo pipefail
#
#ORIGIN=${1:-"local"}
#if [[ "$ORIGIN" != "remote" && "$ORIGIN" != "eks" ]]; then
#  ORIGIN="local"
#fi
#
#CI="$2"
#
#echo "---- VSecM Integration Tests ----"
#echo "Running tests for $ORIGIN origin"
#
#if [[ -z "$CI" ]]; then
#  printf "\n"
#  printf "This script assumes that you have a local minikube cluster running,\n"
#  printf "and you have already installed SPIRE and VMware Secrets Manager.\n"
#  printf "Also, make sure you have executed 'eval \$(minikube docker-env)\'\n"
#  printf "before running this script.\n"
#  printf "\n"
#  read -n 1 -s -r -p "Press any key to proceed…"
#  printf "\n\n"
#fi
#
## ----- Helper Functions -------------------------------------------------------
#
#### Cleanup and Exit ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_
#
## Output a success message.
#happy_exit() {
#  printf "\n"
#  printf "Everything is awesome!\n"
#  printf "\n"
#}
#
## A k8s problem occurred.
#sad_cuddle() {
#  local msg
#  readonly msg=$1
#
#  if [[ -z "$msg" ]]; then
#    printf "Called sad_cuddle() without a message.\n"
#    exit 1
#  fi
#
#  printf "\n"
#  printf "Something went wrong. :(\n"
#  printf "%s\n" "$msg"
#  printf "\n"
#  exit 1
#}
#
## Unit Test
#
##  go test -json ./...
#
## Removes the secret and the demo workload deployment.
#cleanup() {
#  printf "Cleanup…\n"
#
#  local sentinel
#  readonly sentinel=$(define_sentinel)
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -d || sad_cuddle "Cleanup: Failed to delete secret."
#
#  if kubectl get deployment example -n default >/dev/null 2>&1; then
#    kubectl delete deployment example -n default \
#      || sad_cuddle "Cleanup: Failed to delete deployment."
#  else
#    printf "Deployment does not exist, skipping delete step.\n"
#  fi
#
#  # Wait for the workload to be gone.
#  wait_for_example_workload_deletion &
#  wait $!
#}
#
## Deletes the secret associated with the 'example' workload.
#delete_secret() {
#  printf "Deleting secret…\n"
#
#  local sentinel
#  readonly sentinel=$(define_sentinel)
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -d || sad_cuddle "delete_secret: Failed to delete secret."
#
#  printf "Deleted secret.\n"
#}
#
#### Definitions ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
#
## Retrieves the name of the 'example' pod.
#define_example_workload() {
#  local workload
#  readonly workload=$(kubectl get po -n default \
#    | grep "example-" | awk '{print $1}'; exit $?)
#  if [ $? -ne 0 ]; then
#    sad_cuddle "define_example_workload: Failed to define workload."
#  fi
#
#  printf "%s" "$workload"
#}
#
## Retrieves the name of the 'vsecm-sentinel' pod.
#define_sentinel() {
#  local sentinel
#  local max_retries=5
#  local retry_count=0
#
#  while [ $retry_count -lt $max_retries ]; do
#    sentinel=$(kubectl get po -n vsecm-system \
#      | grep "vsecm-sentinel-" | awk '{print $1}'; exit $?)
#
#    # shellcheck disable=SC2181
#    if [ $? -ne 0 ]; then
#      sad_cuddle "define_sentinel: Failed to define sentinel."
#    fi
#
#    if [[ "$sentinel" != *$'\n'* ]]; then
#      break
#    else
#      sleep 10
#      retry_count=$((retry_count + 1))
#    fi
#  done
#
#  # If the maximum number of retries has been reached
#  if [ $retry_count -eq $max_retries ]; then
#    sad_cuddle "define_sentinel: Maximum retries reached."
#    return 1
#  fi
#
#  printf "%s" "$sentinel"
#}
#
## Retrieves the name of the 'vsecm-safe' pod.
#define_safe() {
#  local safe
#  local max_retries=5
#  local retry_count=0
#
#  while [ $retry_count -lt $max_retries ]; do
#    safe=$(kubectl get po -n vsecm-system \
#      | grep "vsecm-safe-" | awk '{print $1}'; exit $?)
#
#    # shellcheck disable=SC2181
#    if [ $? -ne 0 ]; then
#      sad_cuddle "define_safe: Failed to define safe."
#    fi
#
#    if [[ "safe" != *$'\n'* ]]; then
#      break
#    else
#      sleep 10
#      retry_count=$((retry_count + 1))
#    fi
#  done
#
#  # If the maximum number of retries has been reached
#  if [ $retry_count -eq $max_retries ]; then
#    sad_cuddle "define_safe: Maximum retries reached."
#    return 1
#  fi
#
#  printf "%s" "$safe"
#}
#
#### Assertions ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_
#
## Ensures that the argument is not empty.
#assert_exists() {
#  local res
#  readonly res="$1"
#
#  if [[ -z "$res" ]]; then
#    printf "\n"
#    printf "FAIL :(\n"
#    printf "\n"
#    exit 1
#  else
#    printf "\n"
#    printf "PASS \o/\n"
#    printf "\n"
#  fi
#}
#
## Ensure that the argument equals 1.
#assert_only_single_pod() {
#  local pod_count
#  readonly pod_count=$1
#
#  if [[ "$pod_count" -eq 1 ]]; then
#    printf "\n"
#    printf "PASS \o/\n"
#    printf "\n"
#  else
#    printf "\n"
#    printf "FAIL :(\n"
#    printf "\n"
#    exit 1
#  fi
#}
#
## Ensures that the value and the workload’s secret value are equal.
#assert_workload_secret_value() {
#  local workload
#  local value
#  local res
#
#  readonly workload=$(define_example_workload)
#  readonly value=$1
#
#  printf "assert_workload_secret_value()\n"
#  printf "workload: '%s'\n" "$workload"
#  printf "value: '%s'\n" "$value"
#
#  if [[ -z "$workload" || -z "$value" ]]; then
#    sad_cuddle "assert_workload_secret_value: Failed to define workload or value."
#  fi
#
#  readonly res=$(kubectl exec "$workload" -n default -c main -- ./env; exit $?)
#  if [[ $? -ne 0 ]]; then
#    sad_cuddle "assert_workload_secret_value: Failed to exec kubectl."
#  fi
#
#  if [[ "$res" == "$value" ]]; then
#    printf "\n"
#    printf "PASS \o/\n"
#    printf "\n"
#  else
#    printf "\n"
#    printf "FAIL :(\n"
#    printf "\n"
#    exit 1
#  fi
#}
#
## Ensures that the current workload’s secret is empty.
#assert_workload_secret_no_value() {
#  local workload
#  local res
#
#  readonly workload=$(define_example_workload)
#  if [[ -z "$workload" ]]; then
#    sad_cuddle "assert_workload_secret_no_value: Failed to define workload."
#  fi
#
#  readonly res=$(kubectl exec "$workload" -n default -c main -- ./env; exit $?)
#  if [ $? -ne 0 ]; then
#    sad_cuddle "assert_workload_secret_no_value: Failed to exec kubectl."
#  fi
#
#  printf "assert_workload_secret_no_value()\n"
#  printf "workload: '%s'\n" "$workload"
#  printf "res: '%s'\n" "$res"
#
#  if [[ -z "$res" || "$res" == "NO_SECRET" ]]; then
#    printf "\n"
#    printf "PASS \o/\n"
#    printf "\n"
#  else
#    printf "\n"
#    printf "FAIL :(\n"
#    printf "\n"
#    exit 1
#  fi
#}
#
## Ensures if VSecM Sentinel can encrypt a secret.
#assert_encrypted_secret() {
#  local sentinel
#  local value
#
#  readonly sentinel=$(define_sentinel)
#  readonly value=$1
#  if [[ -z "$sentinel" || -z "$value" ]]; then
#    sad_cuddle "assert_encrypted_secret: Failed to define sentinel or value."
#  fi
#
#  res=$(kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -s "$value" \
#    -e; exit $?)
#  if [[ $? -ne 0 ]]; then
#    sad_cuddle "assert_encrypted_secret: Failed to exec kubectl."
#  fi
#
#  assert_exists "$res"
#}
#
## Ensures that the workload is running.
#assert_workload_is_running() {
#  local workload
#  local pod_count
#
#  workload=$(define_example_workload)
#  if [[ -z "$workload" ]]; then
#    sad_cuddle "assert_workload_is_running: Failed to define workload."
#  fi
#
#  pod_count=$(kubectl get po -n default | grep "$workload" | grep -c Running; exit $?)
#  if [[ $? -ne 0 ]]; then
#    sad_cuddle "assert_workload_is_running: Failed to exec kubectl."
#  fi
#  if [[ -z "$pod_count" ]]; then
#    sad_cuddle "assert_workload_is_running: Empty pod_count"
#  fi
#
#  assert_only_single_pod "$pod_count"
#}
#
## Ensures that the init container is running.
#assert_init_container_running() {
#  local workload
#  local pod_status
#  readonly workload=$(define_example_workload)
#
#  if [[ -z $workload ]]; then
#    sad_cuddle "assert_init_container_running: Failed to define workload."
#  fi
#
#  readonly pod_status=$(kubectl get pod -n default "$workload" \
#    -o jsonpath='{.status.initContainerStatuses[0].state.running}'; exit $?)
#  if [[ $? -ne 0 ]]; then
#    sad_cuddle "assert_init_container_running: Failed to exec kubectl."
#  fi
#
#  if [[ -n "$pod_status" ]]; then
#    printf "Init container of pod '%s' is still running.\n" "$workload"
#  else
#    printf "Init container of pod '%s' is not running.\n" "$workload"
#    exit 1
#  fi
#}
#
#### Conditions ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_
#
## Wait for the workload to be ready.
#wait_for_example_workload() {
#  printf "Waiting for example workload…\n"
#
#  max_retries=5
#  retries=0
#
#  while ((retries < max_retries)); do
#    if kubectl wait --for=condition=Ready pod -n default \
#      --selector=app.kubernetes.io/name=example >/dev/null 2>&1; then
#      return 0
#    else
#      retries=$((retries + 1))
#      echo "Retry $retries/$max_retries: Failed to wait for condition. Retrying in 5 seconds..."
#      sleep 5
#    fi
#  done
#
#  sad_cuddle "wait_for_example_workload: Failed to wait for condition after $max_retries retries."
#}
#
#
## Wait until the workload’s deployment is deleted.
#wait_for_example_workload_deletion() {
#  printf "Waiting for example workload deletion…\n"
#
#  kubectl wait --for=delete deployment -n default \
#    --selector=app.kubernetes.io/name=example || \
#    sad_cuddle "wait_for_example_workload_deletion: Failed to wait for deletion."
#}
#
## Pauses the test for 15 seconds to let the sidecar poll the secret.
#pause() {
#  printf "Waiting for 15 seconds to let the sidecar poll the secret…\n"
#  sleep 15
#}
#
## Pauses the test for 15 seconds to let the init container pull the image.
#pause_for_deploy() {
#  printf "Waiting for 15 seconds to pull the image…\n"
#  sleep 15
#}
#
## Pauses the test for 30 seconds to let the init container run, or for other
## operations to complete.
#pause_just_in_case() {
#  printf "Waiting for 30 seconds, just in case…\n"
#  sleep 30
#}
#
#### Mutations ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
#
## Encrypts a secret and stores it in VSecM Safe.
#set_encrypted_secret() {
#  printf "set_encrypted_secret()\n"
#
#  local value
#  local sentinel
#
#  readonly value=$1
#  readonly sentinel=$(define_sentinel)
#  if [[ -z "$value" || -z "$sentinel" ]]; then
#    sad_cuddle "set_encrypted_secret: Failed to define value or sentinel."
#  fi
#
#  printf "value: '%s'\n" "$value"
#  printf "sentinel: '%s'\n" "$sentinel"
#
#  res=$(kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -s "$value" \
#    -e; exit $?)
#
#  printf "res: '%s'\n" "$res"
#
#  if [[ $? -ne 0 ]]; then
#    sad_cuddle "set_encrypted_secret: Failed to exec kubectl."
#  fi
#  if [[ -z "$res" ]]; then
#    sad_cuddle "set_encrypted_secret: Empty res."
#  fi
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -s "$res" \
#    -e || sad_cuddle "set_encrypted_secret: Failed to exec kubectl."
#
#  printf "done: set_encrypted_secret()\n"
#}
#
## Registers a secret in VSecM Safe.
#set_secret() {
#  printf "set_secret()\n"
#
#  local sentinel
#  local value
#
#  readonly sentinel=$(define_sentinel)
#  readonly value=$1
#  if [[ -z "$sentinel" || -z "$value" ]]; then
#    sad_cuddle "set_secret: Failed to define sentinel or value."
#  fi
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -s "$value" || sad_cuddle "set_secret: Failed to exec kubectl."
#
#  printf "done: set_secret()\n"
#}
#
## Registers a secret in VSecM Safe and transforms it as JSON.
#set_json_secret() {
#  printf "set_json_secret()\n"
#
#  local sentinel
#  local value
#  local transform
#
#  readonly sentinel=$(define_sentinel)
#  readonly value=$1
#  readonly transform=$2
#  if [[ -z "$sentinel" || -z "$value" || -z "$transform" ]]; then
#    sad_cuddle "set_json_secret: Failed to define sentinel, value or transform."
#  fi
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -s "$value" \
#    -t "$transform" \
#    -f "json" || sad_cuddle "set_json_secret: Failed to exec kubectl."
#
#  printf "done: set_json_secret()\n"
#}
#
## Registers a secret in VSecM Safe and transforms it as YAML.
#set_yaml_secret() {
#  printf "set_yaml_secret()\n"
#
#  local sentinel
#  local value
#  local transform
#
#  readonly sentinel=$(define_sentinel)
#  readonly value=$1
#  readonly transform=$2
#  if [[ -z "$sentinel" || -z "$value" || -z "$transform" ]]; then
#    sad_cuddle "set_yaml_secret: Failed to define sentinel, value or transform."
#  fi
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -s "$value" \
#    -t "$transform" \
#    -f "yaml" || sad_cuddle "set_yaml_secret: Failed to exec kubectl."
#
#  printf "done: set_yaml_secret()\n"
#}
#
## Registers a secret in VSecM Safe and transforms it as a Kubernetes secret.
#set_kubernetes_secret() {
#  printf "set_kubernetes_secret()\n"
#
#  local sentinel
#  readonly sentinel=$(define_sentinel)
#  if [[ -z "$sentinel" ]]; then
#    sad_cuddle "set_kubernetes_secret: Failed to define sentinel."
#  fi
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -s '{"username": "root", "password": "SuperSecret", "value": "VSecMRocks"}' \
#    -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "value": "{{.value}}"}' \
#    -k || sad_cuddle "set_kubernetes_secret: Failed to exec kubectl."
#
#  # Wait for the workload to be ready.
#  wait_for_example_workload &
#  wait $!
#
#  printf  "done: set_kubernetes_secret()\n"
#}
#
## Append a secret to the workload.
#append_secret() {
#  printf  "append_secret()\n"
#
#  local sentinel
#  local value
#
#  readonly sentinel=$(define_sentinel)
#  readonly value=$1
#  if [[ -z "$sentinel" || -z "$value" ]]; then
#    sad_cuddle "append_secret: Failed to define sentinel or value."
#  fi
#
#  kubectl exec "$sentinel" -n vsecm-system -- safe \
#    -w "example" \
#    -n "default" \
#    -a \
#    -s "$value" || sad_cuddle "append_secret: Failed to exec kubectl."
#
#  printf "done: append_secret()\n"
#}
#
#### Deployments ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
#
## Deploys a workload that uses VSecM SDK.
#deploy_workload_using_sdk() {
#  printf "Deploying workload that uses the SDK…\n"
#
#  if [ "$ORIGIN" == "remote" ]; then
#    make example-sdk-deploy
#  elif [ "$ORIGIN" == "eks" ]; then
#    make example-sdk-deploy-eks
#  else
#    make example-sdk-deploy-local
#  fi
#
#  # Wait for the workload to be ready.
#  wait_for_example_workload &
#  wait $!
#
#  printf "Deployed workload that uses the SDK.\n"
#}
#
## Deploys a workload that uses VSecM Sidecar.
#deploy_workload_using_sidecar() {
#  printf "Deploying workload that uses the sidecar…\n"
#
#  if [ "$ORIGIN" == "remote" ]; then
#    make example-sidecar-deploy
#  elif [ "$ORIGIN" == "eks" ]; then
#    make example-sidecar-deploy-eks
#  else
#    make example-sidecar-deploy-local
#  fi
#
#  # Wait for the workload to be ready.
#  wait_for_example_workload &
#  wait $!
#
#  printf "Deployed workload that uses the sidecar.\n"
#}
#
## Deploys a workload that uses VSecM Init Container.
#deploy_workload_using_init_container() {
#  printf "Deploying workload that uses the init container…\n"
#
#  if [ "$ORIGIN" == "remote" ]; then
#    make example-init-container-deploy
#  elif [ "$ORIGIN" == "eks" ]; then
#    make example-init-container-deploy-eks
#  else
#    make example-init-container-deploy-local
#  fi
#
#  pause_for_deploy
#
#  printf "I should have something there.\n"
#}
#
## ------------------------------------------------------------------------------
#
## Run Go unit tests before running more expensive tests.
#
## Run Go unit tests
#echo "Running Go unit tests..."
#if ! go test ./... -cover; then
#    echo "Go unit tests failed, exiting."
#    exit 1
#fi
#
## ------------------------------------------------------------------------------
#
## Tests the encryption of secrets.
#test_encrypting_secrets() {
#  printf "Testing: Encrypting secrets…\n"
#
#  local value
#  readonly value="!VSecMRocks!"
#
#  assert_encrypted_secret $value
#
#  deploy_workload_using_sdk
#
#  set_encrypted_secret $value
#  assert_workload_secret_value $value
#
#  cleanup
#
#  printf "Tested: Encrypting secrets.\n"
#}
#
#test_encrypting_secrets
#
## ------------------------------------------------------------------------------
#
#cleanup
#printf "\n"
#printf "________________________________________\n"
#printf "Case: Workload using VSecM SDK…\n"
#printf "\n"
#deploy_workload_using_sdk
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets.
#test_secret_registration() {
#  printf "Testing: Secret registration…\n"
#
#  local value
#  readonly value="!VSecMRocks!"
#
#  set_secret $value
#  assert_workload_secret_value $value
#
#  printf "Tested: Secret registration.\n"
#}
#
#test_secret_registration
#
## ------------------------------------------------------------------------------
#
## Tests the deletion of secrets.
#test_secret_deletion() {
#  printf "Testing: Secret deletion…\n"
#
#  delete_secret
#  assert_workload_secret_no_value
#
#  printf "Tested: Secret deletion.\n"
#}
#
#test_secret_deletion
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets in append mode.
#test_secret_registration_append() {
#  printf "Testing: Secret registration (append mode)…\n"
#
#  local secret1
#  local secret2
#  local value
#  readonly secret1="!VSecM"
#  readonly secret2="Rocks!"
#  readonly value='["'"$secret2"'","'"$secret1"'"]'
#
#  append_secret "$secret1"
#  append_secret "$secret2"
#
#  assert_workload_secret_value "$value"
#  delete_secret
#
#  printf "Tested: Secret registration (append mode).\n"
#}
#
#test_secret_registration_append
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets in JSON format.
#test_secret_registration_json_format() {
#  printf "Testing: Secret registration (JSON transformation)…\n"
#
#  local value
#  local transform
#  readonly value='{"username": "*root*", "password": "*CasHC0w*"}'
#  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'
#
#  set_json_secret "$value" "$transform"
#
#  local transformed
#  readonly transformed='{"USERNAME":"*root*", "PASSWORD":"*CasHC0w*"}'
#
#  assert_workload_secret_value "$transformed"
#  delete_secret
#
#  printf "Tested: Secret registration (JSON transformation).\n"
#}
#
#test_secret_registration_json_format
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets in YAML format.
#test_secret_registration_yaml_format() {
#  printf "Testing: Secret registration (YAML transformation)…\n"
#
#  local value
#  local transform
#  readonly value='{"username": "*root*", "password": "*CasHC0w*"}'
#  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'
#
#  set_yaml_secret "$value" "$transform"
#
#  local transformed
#  readonly transformed=$(cat << EOF
#PASSWORD: '*CasHC0w*'
#USERNAME: '*root*'
#EOF
#  )
#
#  assert_workload_secret_value "$transformed"
#  delete_secret
#
#  printf "Tested: Secret registration (YAML transformation).\n"
#}
#
#test_secret_registration_yaml_format
#
## ------------------------------------------------------------------------------
#
#cleanup
#printf "\n"
#printf "________________________________________\n"
#printf "Case: Workload using VSecM Sidecar…\n"
#printf "\n"
#deploy_workload_using_sidecar
#
## Note: for sidecar case, keep in mind that the poll interval is 5 seconds,
## based on the overridden value in the example’s deployment.
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets using VSecM Sidecar.
#test_secret_registration_sidecar() {
#  printf "Testing: Secret registration…\n"
#
#  local value
#  readonly value="!VSecMRocks!"
#
#  set_secret "$value"
#  pause
#  assert_workload_secret_value "$value"
#
#  printf "Tested: Secret registration.\n"
#}
#
#test_secret_registration_sidecar
#
## ------------------------------------------------------------------------------
#
## Tests the deletion of secrets using VSecM Sidecar.
#test_secret_deletion_sidecar() {
#  printf "Testing: Secret deletion (sidecar)…\n"
#
#  delete_secret
#  pause
#  assert_workload_secret_no_value
#
#  printf "Tested: Secret deletion (sidecar).\n"
#}
#
#test_secret_deletion_sidecar
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets in append mode using VSecM Sidecar.
#test_secret_registration_append_sidecar() {
#  printf "Testing Secret registration (append mode)…\n"
#
#  local secret1
#  local secret2
#  local value
#  readonly secret1="!VSecM"
#  readonly secret2="Rocks!"
#  readonly value='["'"$secret2"'","'"$secret1"'"]'
#
#  append_secret "$secret1"
#  append_secret "$secret2"
#
#  pause
#  assert_workload_secret_value "$value"
#  delete_secret
#
#  printf "Tested: Secret registration (append mode).\n"
#}
#
#test_secret_registration_append_sidecar
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets in JSON format using VSecM Sidecar.
#test_secret_registration_json_format_sidecar() {
#  printf "Testing Secret registration (JSON transformation)…\n"
#
#  local value
#  local transform
#  readonly value='{"username": "*root*", "password": "*CasHC0w*"}'
#  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'
#
#  set_json_secret "$value" "$transform"
#
#  local transformed
#  readonly transformed='{"USERNAME":"*root*", "PASSWORD":"*CasHC0w*"}'
#
#  pause
#  assert_workload_secret_value "$transformed"
#  delete_secret
#
#  printf "Tested: Secret registration (JSON transformation).\n"
#}
#
#test_secret_registration_json_format_sidecar
#
## ------------------------------------------------------------------------------
#
## Tests the registration of secrets in YAML format using VSecM Sidecar.
#test_secret_registration_yaml_format_sidecar() {
#  printf "Testing Secret registration (YAML transformation)…\n"
#
#  local value
#  local transform
#  readonly value='{"username": "*root*", "password": "*CasHC0w*"}'
#  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'
#
#  set_yaml_secret "$value" "$transform"
#
#  transformed=$(cat << EOF
#PASSWORD: '*CasHC0w*'
#USERNAME: '*root*'
#EOF
#  )
#
#  pause
#  assert_workload_secret_value "$transformed"
#  delete_secret
#
#  printf "Tested: Secret registration (YAML transformation).\n"
#}
#
#test_secret_registration_yaml_format_sidecar
#
## ------------------------------------------------------------------------------
#
#cleanup
#printf "\n"
#printf "________________________________________\n"
#printf "Case: Workload using VSecM Init Container…\n"
#printf "\n"
#
## Tests the registration of secrets using VSecM Init Container.
#test_init_container() {
#  printf "Testing: Init Container…\n"
#
#  deploy_workload_using_init_container
#  pause_for_deploy
#  define_example_workload
#
#  assert_init_container_running
#  pause_just_in_case
#
#  set_kubernetes_secret
#
#  assert_workload_is_running
#
#  printf "Tested: Init Container.\n"
#}
#
## ------------------------------------------------------------------------------
#
#printf "All done. Cleaning up…\n"
#
#cleanup
#happy_exit
