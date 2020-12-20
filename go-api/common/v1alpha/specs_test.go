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
	"crypto/rand"
	"testing"

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

	for _, v := range []string{"", "x2e", "1.x", "1.1.x"} {
		_, err := ParseVersion(v)
		assert.Error(t, err)
	}

	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &ver))
	assert.Error(t, yaml.Unmarshal([]byte("NONEn"), &ver))

	_, err = yaml.Marshal(&ver)
	assert.NoError(t, err)
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

	_, err = yaml.Marshal(&arch)
	assert.NoError(t, err)

	variant := AMD64.Variant()
	assert.Equal(t, "", variant.String())
	variant = architectureNotSet.Variant()
	assert.Equal(t, "", variant.String())

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
