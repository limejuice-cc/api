// Copyright 2020 Limejuice-cc Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"time"
)

// KeyAlgorithm specifies the type of key algorithm to use
type KeyAlgorithm int

const (
	keyAlgorithmNotSet KeyAlgorithm = iota
	// ECDSAKey specifies the ecdsa algorithm
	ECDSAKey
	// RSAKey specifies an RSA key
	RSAKey
)

// ParseKeyAlgorithm parses a key algorithm
func ParseKeyAlgorithm(in string) (KeyAlgorithm, error) {
	switch in {
	case "ecdsa":
		return ECDSAKey, nil
	case "rsa":
		return RSAKey, nil
	default:
		return keyAlgorithmNotSet, fmt.Errorf("unknown key type: %s", in)
	}
}

func (a KeyAlgorithm) String() string {
	switch a {
	case ECDSAKey:
		return "ecdsa"
	case RSAKey:
		return "rsa"
	}
	return ""
}

// DefaultSize returns the default key size for the specified algorithm
func (a KeyAlgorithm) DefaultSize() int {
	switch a {
	case ECDSAKey:
		return 256
	case RSAKey:
		return 4096
	}
	return 0
}

const (
	minRSAKeySize = 2048
	maxRSAKeySize = 8192
)

// ValidKeySize checks if the supplied key size is valid for the KeyAlgorithm
func (a KeyAlgorithm) ValidKeySize(size int) error {
	switch a {
	case ECDSAKey:
		if !(size == 0 || size == 256 || size == 384 || size == 521) {
			return fmt.Errorf("invalid ecdsa key size %d - key size must be either 256, 384 or 521", size)
		}
		return nil
	case RSAKey:
		if !(size == 0 || (size >= minRSAKeySize && size <= maxRSAKeySize)) {
			return fmt.Errorf("invalid rsa key size %d - key size must be between %d and %d", size, minRSAKeySize, maxRSAKeySize)
		}
		return nil
	}

	return fmt.Errorf("invalid key algorithm")
}

// UnmarshalYAML implements custom unmarshal for KeyAlgorithm
func (a *KeyAlgorithm) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var actionType string
	if err = unmarshal(&actionType); err != nil {
		return
	}

	*a, err = ParseKeyAlgorithm(actionType)
	return
}

// MarshalYAML implements custom marshalling for KeyAlgorithm
func (a KeyAlgorithm) MarshalYAML() (interface{}, error) {
	return a.String(), nil
}

// CertificateName contains subject fields
type CertificateName struct {
	C            string `yaml:"C"`                      // Country
	ST           string `yaml:"ST"`                     // Province
	L            string `yaml:"L"`                      // Locality
	O            string `yaml:"O"`                      // OrganizationName
	OU           string `yaml:"OU,omitempty"`           // OrganizationalUnitName
	SerialNumber string `yaml:"serialNumber,omitempty"` // SerialNumber
}

// CertificateKeyRequest represents a certificate key
type CertificateKeyRequest struct {
	Algorithm KeyAlgorithm `yaml:"algorithm"` // Algorithm
	Size      int          `yaml:"size"`      // Size
}

// CertificatePath represents the the full paths for the requested certificate
type CertificatePath struct {
	Certificate string `yaml:"cert"` // Certificate is full path of the certificate
	Key         string `yaml:"key"`  // Key is full path of the private key
}

// CertificateRequest represents a certificate request
type CertificateRequest struct {
	Key          CertificateKeyRequest `yaml:"key"`                    // Key
	CommonName   string                `yaml:"commonName,omitempty"`   // CommonName
	Names        []CertificateName     `yaml:"names,omitempty"`        // Names
	Hosts        []string              `yaml:"hosts,omitempty"`        // Hosts
	SerialNumber string                `yaml:"serialNumber,omitempty"` // SerialNumber
	Usage        []string              `yaml:"usage,omitempty"`        // Usage
	Expires      time.Duration         `yaml:"expires,omitempty"`      // Expires
	IsCA         bool                  `yaml:"ca,omitempty"`           // Certificate Authority
	Path         CertificatePath       `yaml:"path"`                   // Path
}

// CertificateRequests is a list of certificate requests
type CertificateRequests []*CertificateKeyRequest

// CertificatePackage represents a package of certificates
type CertificatePackage struct {
	CertificateAuthorityRequest *CertificateRequest `yaml:"caRequest,omitempty"` // CertificateAuthorityRequest is the certificate authority request
	CertificateAuthority        string              `yaml:"ca,omitempty"`        // CertificateAuthority is the pem encoded certificate authority
	CertificateAuthorityKey     string              `yaml:"caKey,omitempty"`     // CertificateAuthorityKey is the pem encoded private key of the certificate authority
	Requests                    CertificateRequests `yaml:"requests,omitempty"`  // Requests is a list of certificates to use the CA
}

// CertificateHosts is a generic interface for hosts
type CertificateHosts interface {
	DNSNames() []string
	EmailAddresses() []string
	IPAddresses() []net.IP
	URIs() []*url.URL
}

// DistinguishedName represents certificate subject information
type DistinguishedName interface {
	CommonName() string
	Countries() []string
	Provinces() []string
	States() []string
	Localities() []string
	Organizations() []string
	OrganizationalUnits() []string
	SerialNumber() *big.Int
}

// Key represents a key
type Key interface {
	Algorithm() KeyAlgorithm
	Size() int
	Encoded() []byte
	PrivateKey() crypto.PrivateKey
	PublicKeyAlgorithm() x509.PublicKeyAlgorithm
	PublicKey() crypto.PublicKey
	SignatureAlgorithm() x509.SignatureAlgorithm
}

// Certificate represents a generic certificate interface
type Certificate interface {
	Encoded() []byte
	Certificate() *x509.Certificate
	PrivateKey() Key

	CA() bool
	Subject() DistinguishedName
	Hosts() CertificateHosts
	Expires() time.Time
	Usage() CertificateKeyUsages
	SerialNumber() *big.Int

	SelfSign() (Certificate, error)
	Sign(Certificate) (Certificate, error)
}

// CertificateKeyUsages is an interface for a set of certificate key usages
type CertificateKeyUsages interface {
	Standard() x509.KeyUsage
	Extended() []x509.ExtKeyUsage
}

// CertificateProvider is a generic interface to a provider capable of generate certificates
type CertificateProvider interface {
	Initialize(options ...CertificateProviderOption) error
	ParseEncoded(certificate, privateKey []byte) (Certificate, error)
	Generate(request *CertificateRequest) (Certificate, error)
}

// CertificateProviderOption is an option when initalizing a CertificateProvider
type CertificateProviderOption interface {
	Apply(CertificateProvider) error
}
