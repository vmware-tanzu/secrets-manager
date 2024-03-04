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

echo ""
echo ""
echo "KEYCLOAK_ADMIN_PASSWORD:"
kubectl get secret keycloak-secret -n smo-app -o jsonpath='{.data.KEYCLOAK_ADMIN_PASSWORD}' | base64 --decode
echo ""
echo ""

echo "KEYCLOAK_ADMIN_USER:"
kubectl get secret keycloak-secret -n smo-app -o jsonpath='{.data.KEYCLOAK_ADMIN_USER}' | base64 --decode
echo ""
echo ""

echo "KEYCLOAK_DB_PASSWORD:"
kubectl get secret keycloak.ric-postgres.credentials -n smo-app -o jsonpath='{.data.KEYCLOAK_DATABASE_PASSWORD}' | base64 --decode
echo ""
echo ""

echo "KEYCLOAK_DB_USER:"
kubectl get secret keycloak.ric-postgres.credentials -n smo-app -o jsonpath='{.data.KEYCLOAK_DATABASE_USER}' | base64 --decode
echo ""
echo ""
