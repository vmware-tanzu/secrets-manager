/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
	"reflect"
	"testing"
	"time"
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
					t.Errorf("SecretBufferSizeForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SecretBufferSizeForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SecretBufferSizeForSafe(); got != tt.want {
				t.Errorf("SecretBufferSizeForSafe() = %v, want %v", got, tt.want)
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
					t.Errorf("K8sSecretBufferSizeForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("K8sSecretBufferSizeForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := K8sSecretBufferSizeForSafe(); got != tt.want {
				t.Errorf("K8sSecretBufferSizeForSafe() = %v, want %v", got, tt.want)
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
					t.Errorf("SecretDeleteBufferSizeForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SecretDeleteBufferSizeForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SecretDeleteBufferSizeForSafe(); got != tt.want {
				t.Errorf("SecretDeleteBufferSizeForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestSafeRemoveLinkedK8sSecrets(t *testing.T) {
//	tests := []struct {
//		name    string
//		setup   func() error
//		cleanup func() error
//		want    bool
//	}{
//		{
//			name: "default_safe_remove_linked_k8s_secrets",
//			want: false,
//		},
//		{
//			name: "safe_remove_linked_k8s_secrets_from_env_true",
//			setup: func() error {
//				return os.Setenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS", "true")
//			},
//			cleanup: func() error {
//				return os.Unsetenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS")
//			},
//			want: true,
//		},
//		{
//			name: "safe_remove_linked_k8s_secrets_from_env_false",
//			setup: func() error {
//				return os.Setenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS", "false")
//			},
//			cleanup: func() error {
//				return os.Unsetenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS")
//			},
//			want: false,
//		},
//		{
//			name: "invalid_safe_remove_linked_k8s_secrets_from_env",
//			setup: func() error {
//				return os.Setenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS", "test")
//			},
//			cleanup: func() error {
//				return os.Unsetenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS")
//			},
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if tt.setup != nil {
//				if err := tt.setup(); err != nil {
//					t.Errorf("RemoveLinkedK8sSecretsModeForSafe() = failed to setup, with error: %+v", err)
//				}
//			}
//			defer func() {
//				if tt.cleanup != nil {
//					if err := tt.cleanup(); err != nil {
//						t.Errorf("RemoveLinkedK8sSecretsModeForSafe() = failed to cleanup, with error: %+v", err)
//					}
//				}
//			}()
//			if got := RemoveLinkedK8sSecretsModeForSafe(); got != tt.want {
//				t.Errorf("RemoveLinkedK8sSecretsModeForSafe() = %v, want %v", got, tt.want)
//			}
//		},
//		)
//	}
//}

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
					t.Errorf("FipsCompliantModeForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("FipsCompliantModeForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := FipsCompliantModeForSafe(); got != tt.want {
				t.Errorf("FipsCompliantModeForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestSafeBackingStore(t *testing.T) {
//	tests := []struct {
//		name    string
//		setup   func() error
//		cleanup func() error
//		want    data.BackingStore
//	}{
//		{
//			name: "default_safe_backing_store",
//			want: data.File,
//		},
//		{
//			name: "safe_backing_store_from_env_file",
//			setup: func() error {
//				return os.Setenv("VSECM_SAFE_BACKING_STORE", "file")
//			},
//			cleanup: func() error {
//				return os.Unsetenv("VSECM_SAFE_BACKING_STORE")
//			},
//			want: data.File,
//		},
//		{
//			name: "safe_backing_store_from_env_memory",
//			setup: func() error {
//				return os.Setenv("VSECM_SAFE_BACKING_STORE", "memory")
//			},
//			cleanup: func() error {
//				return os.Unsetenv("VSECM_SAFE_BACKING_STORE")
//			},
//			want: data.Memory,
//		},
//		{
//			name: "invalid_safe_backing_store_from_env",
//			setup: func() error {
//				return os.Setenv("VSECM_SAFE_BACKING_STORE", "test")
//			},
//			cleanup: func() error {
//				return os.Unsetenv("VSECM_SAFE_BACKING_STORE")
//			},
//			want: data.Memory,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if tt.setup != nil {
//				if err := tt.setup(); err != nil {
//					t.Errorf("BackingStoreForSafe() = failed to setup, with error: %+v", err)
//				}
//			}
//			defer func() {
//				if tt.cleanup != nil {
//					if err := tt.cleanup(); err != nil {
//						t.Errorf("BackingStoreForSafe() = failed to cleanup, with error: %+v", err)
//					}
//				}
//			}()
//			if got := BackingStoreForSafe(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("BackingStoreForSafe() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
					t.Errorf("SecretBackupCountForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SecretBackupCountForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SecretBackupCountForSafe(); got != tt.want {
				t.Errorf("SecretBackupCountForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootKeyInputMode(t *testing.T) {
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
				return os.Setenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL", "true")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL")
			},
			want: true,
		},
		{
			name: "safe_manual_key_input_from_env_false",
			setup: func() error {
				return os.Setenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL", "false")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL")
			},
			want: false,
		},
		{
			name: "invalid_safe_manual_key_input_from_env",
			setup: func() error {
				return os.Setenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL")
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("RootKeyInputModeManual() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("RootKeyInputModeManual() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := RootKeyInputModeManual(); got != tt.want {
				t.Errorf("RootKeyInputModeManual() = %v, want %v", got, tt.want)
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
			want: "/var/local/vsecm/data",
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
					t.Errorf("DataPathForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("DataPathForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := DataPathForSafe(); got != tt.want {
				t.Errorf("DataPathForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeRootKeyPath(t *testing.T) {
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
				return os.Setenv("VSECM_ROOT_KEY_PATH", "/opt/test_key.txt")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_ROOT_KEY_PATH")
			},
			want: "/opt/test_key.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("RootKeyPathForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("RootKeyPathForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := RootKeyPathForSafe(); got != tt.want {
				t.Errorf("RootKeyPathForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSourceAcquisitionTimeout(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    time.Duration
	}{
		{
			name: "default_safe_source_acquisition_timeout",
			want: 10000 * time.Millisecond,
		},
		{
			name: "safe_source_acquisition_timeout_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT", "500")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT")
			},
			want: 500 * time.Millisecond,
		},
		{
			name: "invalid_safe_source_acquisition_timeout_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT")
			},
			want: 10000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SourceAcquisitionTimeoutForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SourceAcquisitionTimeoutForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SourceAcquisitionTimeoutForSafe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SourceAcquisitionTimeoutForSafe() = %v, want %v", got, tt.want)
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
			want: 300000 * time.Millisecond,
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
			want: 300000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("BootstrapTimeoutForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("BootstrapTimeoutForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := BootstrapTimeoutForSafe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BootstrapTimeoutForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeRootKeySecretName(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_crypto_key_name",
			want: "vsecm-root-key",
		},
		{
			name: "crypto_key_name_from_env",
			setup: func() error {
				return os.Setenv("VSECM_ROOT_KEY_NAME", "vsecm-root-key-test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_ROOT_KEY_NAME")
			},
			want: "vsecm-root-key-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("RootKeySecretNameForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("RootKeySecretNameForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := RootKeySecretNameForSafe(); got != tt.want {
				t.Errorf("RootKeySecretNameForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeIvInitializationInterval(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		cleanup func()
		want    int
	}{
		{
			name: "get default IV initialization interval",
			want: 50,
		},
		{
			name: "get custom IV initialization interval",
			setup: func() {
				if err := os.Setenv("VSECM_SAFE_IV_INITIALIZATION_INTERVAL", "20"); err != nil {
					t.Errorf("IvInitializationIntervalForSafe = failed to setup, with error: %+v", err)
				}
			},
			cleanup: func() {
				if err := os.Unsetenv("VSECM_SAFE_IV_INITIALIZATION_INTERVAL"); err != nil {
					t.Errorf("IvInitializationIntervalForSafe = failed to cleanup, with error: %+v", err)
				}
			},
			want: 20,
		},
		{
			name: "invalid IV initialization interval",
			setup: func() {
				if err := os.Setenv("VSECM_SAFE_IV_INITIALIZATION_INTERVAL", "abc"); err != nil {
					t.Errorf("IvInitializationIntervalForSafe = failed to setup, with error: %+v", err)
				}
			},
			cleanup: func() {
				if err := os.Unsetenv("VSECM_SAFE_IV_INITIALIZATION_INTERVAL"); err != nil {
					t.Errorf("IvInitializationIntervalForSafe = failed to cleanup, with error: %+v", err)
				}
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if got := IvInitializationIntervalForSafe(); got != tt.want {
				t.Errorf("IvInitializationIntervalForSafe() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}
