#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
export SENTINEL=$SENTINEL

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "k8s:keycloak-secret" \
  -n "smo-app" \
  -s 'gen:{"username": "admin-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}' \
  -t '{"KEYCLOAK_ADMIN_USER":"{{.username}}", "KEYCLOAK_ADMIN_PASSWORD":"{{.password}}"}'

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "k8s:keycloak.smo-postgres.credentials" \
  -n "smo-app" \
  -s 'gen:{"username": "dbroot-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}' \
  -t '{"KEYCLOAK_DATABASE_USER":"{{.username}}", "KEYCLOAK_DATABASE_PASSWORD":"{{.password}}"}'
