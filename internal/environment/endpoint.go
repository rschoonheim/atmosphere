package environment

import (
	"errors"
)

type Endpoint struct {
	// Name = the name of the endpoint.
	Name string `json:"name"`

	// Port - the port that the service will listen on.
	Port int `json:"port"`

	// TlsCert - the paths of the TLS certificate and key files for the service.
	TlsCert TlsCertificate `json:"tls_cert"`
}

// HasCertificates - returns true if endpoint has TLS certificates.
func (endpoint *Endpoint) HasCertificates() (bool, error) {
	if endpoint.TlsCert.Cert == "" {
		return false, errors.New("certificate_file_not_configured")
	}
	if endpoint.TlsCert.Key == "" {
		return false, errors.New("key_file_not_configured")
	}
	return true, nil
}
