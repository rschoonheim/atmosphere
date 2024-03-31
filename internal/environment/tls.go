package environment

import (
	"errors"
	"os"
)

type TlsCertificate struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

// Exists - returns true if the TLS certificate and key files exist.
func (tlsCert *TlsCertificate) Exists() (bool, error) {
	if _, err := os.Stat(tlsCert.Cert); os.IsNotExist(err) {
		return false, errors.New("certificate_file_not_found")
	}
	if _, err := os.Stat(tlsCert.Key); os.IsNotExist(err) {
		return false, errors.New("key_file_not_found")
	}
	return true, nil
}
