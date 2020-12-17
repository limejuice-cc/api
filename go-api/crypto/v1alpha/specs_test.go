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
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestKeyAlgorithm(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome KeyAlgorithm
	}{
		{"rsa", RSAKey},
		{"ecdsa", ECDSAKey},
	}

	for _, v := range testValues {
		a, err := ParseKeyAlgorithm(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, a)
			assert.Equal(t, v.value, a.String())
		}
	}

	_, err := ParseKeyAlgorithm("")
	assert.Error(t, err)
	assert.Equal(t, "", keyAlgorithmNotSet.String())
	assert.Error(t, keyAlgorithmNotSet.ValidKeySize(123))
	assert.Error(t, RSAKey.ValidKeySize(minRSAKeySize-1))
	assert.Error(t, RSAKey.ValidKeySize(maxRSAKeySize+1))
	assert.Error(t, ECDSAKey.ValidKeySize(123))
	assert.Equal(t, 0, keyAlgorithmNotSet.DefaultSize())
	assert.NoError(t, RSAKey.ValidKeySize(RSAKey.DefaultSize()))
	assert.NoError(t, ECDSAKey.ValidKeySize(ECDSAKey.DefaultSize()))

	var ka KeyAlgorithm
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &ka))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &ka))
}

func TestParseCertificateKeyRequest(t *testing.T) {
	key := CertificateKeyRequest{ECDSAKey, ECDSAKey.DefaultSize()}
	out, err := yaml.Marshal(&key)
	if assert.NoError(t, err) {
		var k CertificateKeyRequest
		assert.NoError(t, yaml.Unmarshal(out, &k))
	}
}
