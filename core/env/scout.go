package env

func ScoutEnableTls() bool {
	return false
}

type AuthenticationMode string

var AuthenticationModeNone AuthenticationMode = "none"
var AuthenticationModeJwt AuthenticationMode = "jwt"
var AuthenticationModeBasic AuthenticationMode = "basic"

func ScoutAuthenticationMode() AuthenticationMode {
	return AuthenticationModeNone
}

func ScoutTlsEnabled() bool {
	return false
}

func ScoutHttpPort() string {
	return ":8080"
}

func ScoutHttpsPort() string {
	return ":8443"
}
