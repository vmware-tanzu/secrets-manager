source ./env.sh

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "k8s:keycloak-secret" \
  -n "default" \
  -s 'gen:{"username": "admin-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}' \
  -t '{"KEYCLOAK_ADMIN_USER":"{{.username}}", "KEYCLOAK_ADMIN_PASSWORD":"{{.password}}"}'

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "k8s:keycloak.smo-postgres.credentials" \
  -n "default" \
  -s 'gen:{"username": "dbroot-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}' \
  -t '{"KEYCLOAK_DATABASE_USER":"{{.username}}", "KEYCLOAK_DATABASE_PASSWORD":"{{.password}}"}'
