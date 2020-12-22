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
	"fmt"

	"github.com/limejuice-cc/api/helper"
)

// KeyAlgorithm specifies the type of key algorithm to use
type KeyAlgorithm int

const (
	_ KeyAlgorithm = iota
	// ECDSAKey specifies the ecdsa algorithm
	ECDSAKey
	// RSAKey specifies an RSA key
	RSAKey
)

var keyAlgorithmValues = helper.EnumeratorValues{
	"ecdsa": ECDSAKey,
	"rsa":   RSAKey,
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

// String implements the Stringer interface.
func (a KeyAlgorithm) String() string {
	return keyAlgorithmValues.AsString(a)
}

// ParseKeyAlgorithm attempts to convert a string to a KeyAlgorithm
func ParseKeyAlgorithm(name string) (KeyAlgorithm, error) {
	x, err := keyAlgorithmValues.Parse(name)
	if err != nil {
		return KeyAlgorithm(0), err
	}
	return x.(KeyAlgorithm), nil
}

// MarshalText implements the text marshaller method
func (a KeyAlgorithm) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (a *KeyAlgorithm) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseKeyAlgorithm(name)
	if err != nil {
		return err
	}
	*a = tmp
	return nil
}
