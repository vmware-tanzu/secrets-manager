/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
	"reflect"
	"testing"
	"time"

	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
)

func TestSafeSecretBufferSize(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_safe_secret_buffer_side",
			want: 10,
		},
		{
			name: "safe_secret_buffer_side_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_BUFFER_SIZE", "50")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_BUFFER_SIZE")
			},
			want: 50,
		},
		{
			name: "invalid_safe_secret_buffer_side_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_BUFFER_SIZE", "2a")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_BUFFER_SIZE")
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeSecretBufferSize() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeSecretBufferSize() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeSecretBufferSize(); got != tt.want {
				t.Errorf("SafeSecretBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeK8sSecretBufferSize(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_safe_k8s_secret_buffer_size",
			want: 10,
		},
		{
			name: "safe_k8s_secret_buffer_size_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_K8S_SECRET_BUFFER_SIZE", "90")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_K8S_SECRET_BUFFER_SIZE")
			},
			want: 90,
		},
		{
			name: "invalid_safe_k8s_secret_buffer_size_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_K8S_SECRET_BUFFER_SIZE", "abc")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_K8S_SECRET_BUFFER_SIZE")
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeK8sSecretBufferSize() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeK8sSecretBufferSize() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeK8sSecretBufferSize(); got != tt.want {
				t.Errorf("SafeK8sSecretBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSecretDeleteBufferSize(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_safe_secret_delete_buffer_size",
			want: 10,
		},
		{
			name: "safe_secret_delete_buffer_size_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE", "5")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE")
			},
			want: 5,
		},
		{
			name: "invalid_safe_secret_delete_buffer_size_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE", "xyz")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE")
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeSecretDeleteBufferSize() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeSecretDeleteBufferSize() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeSecretDeleteBufferSize(); got != tt.want {
				t.Errorf("SafeSecretDeleteBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeK8sSecretDeleteBufferSize(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_safe_k8s_secret_delete_buffer_size",
			want: 10,
		},
		{
			name: "safe_k8s_secret_delete_buffer_size_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE", "1")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE")
			},
			want: 1,
		},
		{
			name: "invalid_safe_k8s_secret_delete_buffer_size_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE")
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeK8sSecretDeleteBufferSize() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeK8sSecretDeleteBufferSize() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeK8sSecretDeleteBufferSize(); got != tt.want {
				t.Errorf("SafeK8sSecretDeleteBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeFipsCompliant(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    bool
	}{
		{
			name: "default_safe_fips_compliant",
			want: false,
		},
		{
			name: "safe_fips_compliant_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_FIPS_COMPLIANT", "true")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_FIPS_COMPLIANT")
			},
			want: true,
		},
		{
			name: "invalid_safe_fips_compliant_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_FIPS_COMPLIANT", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_FIPS_COMPLIANT")
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeFipsCompliant() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeFipsCompliant() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeFipsCompliant(); got != tt.want {
				t.Errorf("SafeFipsCompliant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeBackingStore(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    data.BackingStore
	}{
		{
			name: "default_safe_backing_store",
			want: data.File,
		},
		{
			name: "safe_backing_store_from_env_file",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_BACKING_STORE", "file")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_BACKING_STORE")
			},
			want: data.File,
		},
		{
			name: "safe_backing_store_from_env_meomry",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_BACKING_STORE", "memory")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_BACKING_STORE")
			},
			want: data.Memory,
		},
		{
			name: "invalid_safe_backing_store_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_BACKING_STORE", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_BACKING_STORE")
			},
			want: data.Memory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeBackingStore() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeBackingStore() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeBackingStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeBackingStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUseKubernetesSecrets(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    bool
	}{
		{
			name: "default_safe_use_kubernetes_secrets",
			want: false,
		},
		{
			name: "safe_use_kubernetes_secrets_from_env_true",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_USE_KUBERNETES_SECRETS", "true")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_USE_KUBERNETES_SECRETS")
			},
			want: true,
		},
		{
			name: "safe_use_kubernetes_secrets_from_env_true",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_USE_KUBERNETES_SECRETS", "false")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_USE_KUBERNETES_SECRETS")
			},
			want: false,
		},
		{
			name: "invalid_safe_use_kubernetes_secrets_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_USE_KUBERNETES_SECRETS", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_USE_KUBERNETES_SECRETS")
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeUseKubernetesSecrets() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeUseKubernetesSecrets() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeUseKubernetesSecrets(); got != tt.want {
				t.Errorf("SafeUseKubernetesSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSecretBackupCount(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_safe_secret_backup_count",
			want: 3,
		},
		{
			name: "safe_secret_backup_count_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_BACKUP_COUNT", "10")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_BACKUP_COUNT")
			},
			want: 10,
		},
		{
			name: "invalid_safe_secret_backup_count_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_BACKUP_COUNT", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_BACKUP_COUNT")
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeSecretBackupCount() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeSecretBackupCount() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeSecretBackupCount(); got != tt.want {
				t.Errorf("SafeSecretBackupCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeManualKeyInput(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    bool
	}{
		{
			name: "default_safe_manual_key_input",
			want: false,
		},
		{
			name: "safe_manual_key_input_from_env_true",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_MANUAL_KEY_INPUT", "true")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_MANUAL_KEY_INPUT")
			},
			want: true,
		},
		{
			name: "safe_manual_key_input_from_env_false",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_MANUAL_KEY_INPUT", "false")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_MANUAL_KEY_INPUT")
			},
			want: false,
		},
		{
			name: "invalid_safe_manual_key_input_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_MANUAL_KEY_INPUT", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_MANUAL_KEY_INPUT")
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeManualKeyInput() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeManualKeyInput() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeManualKeyInput(); got != tt.want {
				t.Errorf("SafeManualKeyInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeDataPath(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_data_path",
			want: "/data",
		},
		{
			name: "safe_data_path_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_DATA_PATH", "/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_DATA_PATH")
			},
			want: "/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeDataPath() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeDataPath() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeDataPath(); got != tt.want {
				t.Errorf("SafeDataPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeAgeKeyPath(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_crypto_key_path",
			want: "/key/key.txt",
		},
		{
			name: "crypto_key_path_from_env",
			setup: func() error {
				return os.Setenv("VSECM_CRYPTO_KEY_PATH", "/opt/test_key.txt")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_CRYPTO_KEY_PATH")
			},
			want: "/opt/test_key.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeAgeKeyPath() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeAgeKeyPath() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeAgeKeyPath(); got != tt.want {
				t.Errorf("SafeAgeKeyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeBootstrapTimeout(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    time.Duration
	}{
		{
			name: "default_safe_bootstrap_timeout",
			want: 30000 * time.Millisecond,
		},
		{
			name: "safe_bootstrap_timeout_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_BOOTSTRAP_TIMEOUT", "500")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_BOOTSTRAP_TIMEOUT")
			},
			want: 500 * time.Millisecond,
		},
		{
			name: "invalid_safe_bootstrap_timeout_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_BOOTSTRAP_TIMEOUT", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_BOOTSTRAP_TIMEOUT")
			},
			want: 30000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeBootstrapTimeout() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeBootstrapTimeout() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeBootstrapTimeout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeBootstrapTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeAgeKeySecretName(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_crypto_key_name",
			want: "vsecm-safe-age-key",
		},
		{
			name: "crypto_key_name_from_env",
			setup: func() error {
				return os.Setenv("VSECM_CRYPTO_KEY_NAME", "vsecm-safe-age-key-test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_CRYPTO_KEY_NAME")
			},
			want: "vsecm-safe-age-key-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeAgeKeySecretName() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeAgeKeySecretName() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeAgeKeySecretName(); got != tt.want {
				t.Errorf("SafeAgeKeySecretName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSecretNamePrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_secret_name_prefix",
			want: "vsecm-secret-",
		},
		{
			name: "safe_secret_name_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SECRET_NAME_PREFIX", "vsecm-secret-test-")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SECRET_NAME_PREFIX")
			},
			want: "vsecm-secret-test-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeSecretNamePrefix() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeSecretNamePrefix() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeSecretNamePrefix(); got != tt.want {
				t.Errorf("SafeSecretNamePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
