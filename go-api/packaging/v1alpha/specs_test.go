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
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestParseVersion(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome Version
	}{
		{"v1.2.0", Version{1, 2, 0, ""}},
		{"v1.2", Version{1, 2, 0, ""}},
		{"v1", Version{1, 0, 0, ""}},
		{"1.2.0", Version{1, 2, 0, ""}},
		{"v1.2.0-test", Version{1, 2, 0, "test"}},
	}

	for _, v := range testValues {
		ver, err := ParseVersion(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome.Major, ver.Major)
			assert.Equal(t, v.outcome.Minor, ver.Minor)
			assert.Equal(t, v.outcome.Patch, ver.Patch)
			assert.Equal(t, v.outcome.Tag, ver.Tag)
		}
	}

	ver, err := ParseVersion("1")
	if assert.NoError(t, err) {
		assert.Equal(t, "v1.0.0", ver.String())
	}

	ver, err = ParseVersion("1.0.0-test")
	if assert.NoError(t, err) {
		assert.Equal(t, "v1.0.0-test", ver.String())
	}

	errVals := []string{"", "x2e", "1.x", "1.1.x"}
	for _, v := range errVals {
		_, err := ParseVersion(v)
		assert.Error(t, err)
	}
}

func TestParseRelationship(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome Relationship
	}{
		{"suggests", Suggests},
		{"recommends", Recommends},
		{"depends", Depends},
		{"predepends", Predepends},
		{"breaks", Breaks},
		{"conflicts", Conflicts},
		{"provides", Provides},
		{"replaces", Replaces},
	}

	for _, v := range testValues {
		r, err := ParseRelationship(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, r)
			assert.Equal(t, v.value, r.String())
		}
	}

	_, err := ParseRelationship("")
	assert.Error(t, err)
	assert.Equal(t, "", noRelationshipSet.String())
}

func TestParseRequired(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome Required
	}{
		{"==", RequiresEqual},
		{">>", RequiresGreaterThan},
		{">=", RequiresGreaterThanEqual},
		{"<<", RequiresLessThan},
		{"<=", RequiresLessThanEqual},
	}

	for _, v := range testValues {
		r, err := ParseRequired(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, r)
			assert.Equal(t, v.value, r.String())
		}
	}

	_, err := ParseRequired("")
	assert.Error(t, err)
	assert.Equal(t, "", requiredNotSet.String())
}

func TestParseFileType(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome FileType
	}{
		{"config", ConfigurationFile},
		{"exec", ExecutableFile},
		{"data", DataFile},
		{"other", OtherFile},
	}

	for _, v := range testValues {
		f, err := ParseFileType(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, f)
			assert.Equal(t, v.value, f.String())
		}
	}

	_, err := ParseFileType("")
	assert.NoError(t, err)
	_, err = ParseFileType("nothing")
	assert.Error(t, err)
}

func TestParseActionType(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome ActionType
	}{
		{"install", Install},
		{"reconfigure", Reconfigure},
		{"upgrade", Upgrade},
		{"remove", Remove},
		{"purge", Purge},
	}

	for _, v := range testValues {
		a, err := ParseActionType(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, a)
			assert.Equal(t, v.value, a.String())
		}
	}

	_, err := ParseActionType("")
	assert.Error(t, err)
	assert.Equal(t, "", noActionTypeSet.String())
}

func TestMarshalManifest(t *testing.T) {
	manifest := Manifest{}
	manifest.Name = "test"
	manifest.Version.Major = 1
	manifest.Created = time.Now()
	manifest.Metadata.Architectures = Architectures{AMD64}
	manifest.Dependencies = Dependencies{&Dependency{"test", Version{1, 0, 0, ""}, RequiresEqual, Depends}}
	manifest.Files = Files{&File{Path: "test file"}}
	manifest.Actions = Actions{&Action{Type: Install}}
	out, err := yaml.Marshal(&manifest)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, manifest.Version.Major)
		assert.Equal(t, 0, manifest.Version.Minor)
		assert.Equal(t, 0, manifest.Version.Patch)
		assert.Equal(t, "", manifest.Version.Tag)
		var m Manifest
		assert.NoError(t, yaml.Unmarshal(out, &m))
	}

	var pn PackageName
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &pn))
	assert.Error(t, yaml.Unmarshal([]byte("NONEn"), &pn))

	assert.Error(t, yaml.Unmarshal([]byte("name: test\nversion: [1,2,3]"), &manifest))
	assert.Error(t, yaml.Unmarshal([]byte("name: test\nversion: dead beef"), &manifest))

	var requires Required
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &requires))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &requires))

	var relation Relationship
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &relation))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &relation))

	var ft FileType
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &ft))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &ft))

	var at ActionType
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &at))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &at))
}

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

func TestParseArchitecture(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome Architecture
	}{
		{"amd64", AMD64},
	}

	for _, v := range testValues {
		a, err := ParseArchitecture(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, a)
			assert.Equal(t, v.value, a.String())
		}
	}

	_, err := ParseArchitecture("")
	assert.NoError(t, err)
	assert.Equal(t, "amd64", architectureNotSet.String())

	var arch Architecture
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &arch))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &arch))

	main, variant := AMD64.Split()
	assert.Equal(t, "amd64", main)
	assert.Equal(t, "", variant)
	main, variant = architectureNotSet.Split()
	assert.Equal(t, "amd64", main)
	assert.Equal(t, "", variant)

}

func TestParseOperatingSystem(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome OperatingSystem
	}{
		{"linux", Linux},
	}

	for _, v := range testValues {
		a, err := ParseOperatingSystem(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, a)
			assert.Equal(t, v.value, a.String())
		}
	}

	_, err := ParseOperatingSystem("")
	assert.NoError(t, err)
	assert.Equal(t, "linux", noOperatingSystemSet.String())

	var os OperatingSystem
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &os))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &os))

	_, err = yaml.Marshal([]OperatingSystem{Linux})
	assert.NoError(t, err)
}

func TestParseCertificateKeyRequest(t *testing.T) {
	key := CertificateKeyRequest{ECDSAKey, ECDSAKey.DefaultSize()}
	out, err := yaml.Marshal(&key)
	if assert.NoError(t, err) {
		var k CertificateKeyRequest
		assert.NoError(t, yaml.Unmarshal(out, &k))
	}
}

func TestEmbeddedFileContents(t *testing.T) {
	raw := make([]byte, 256)
	_, err := rand.Read(raw)
	if assert.NoError(t, err) {
		f := EmbeddedFiles{}
		f["test"] = raw
		out, err := yaml.Marshal(&f)
		if assert.NoError(t, err) {
			var v EmbeddedFiles
			assert.NoError(t, yaml.Unmarshal(out, &v))
		}
	}

	var c EmbeddedFileContents
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &c))
	assert.Error(t, yaml.Unmarshal([]byte("!!!!!>>>>"), &c))

}
