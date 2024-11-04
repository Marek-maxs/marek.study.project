package vault

import "time"

// Config may setup the vault client configuration. If any field is a zero
// value, it will be ignored and the default used.
type Config struct {
	Address string
	AgentAddress string
	ClientTimeout time.Duration
	MaxRetries int
	Token string
	TLSConfig
}

// TLSConfig is the configuration for TLS.
type TLSConfig struct {
	CACert             string
	CAPath             string
	ClientCert         string
	ClientKey          string
	InsecureSkipVerify bool
	TLSServerName      string
}