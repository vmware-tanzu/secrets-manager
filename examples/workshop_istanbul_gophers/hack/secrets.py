import base64
import subprocess
import json

def get_secret(secret_name):
  try:
    result = subprocess.run(["kubectl", "get", "secret", secret_name, "-o", "json"], capture_output=True, check=True, text=True)
    secret_data = json.loads(result.stdout)

    decoded_data = {key: base64.b64decode(value).decode('utf-8') for key, value in secret_data['data'].items()}
    return decoded_data
  except subprocess.CalledProcessError as e:
    print(f"An error occurred: {e}")
    return {}

keycloak_secret_data = get_secret("keycloak-secret")
postgres_credentials_data = get_secret("keycloak.smo-postgres.credentials")

print("Keycloak Secret Data:", keycloak_secret_data)
print("Postgres Credentials Data:", postgres_credentials_data)
