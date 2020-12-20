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
	"time"

	common "github.com/limejuice-cc/api/go-api/common/v1alpha"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

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
	manifest.Metadata.Architectures = common.Architectures{common.AMD64}
	manifest.Dependencies = Dependencies{&Dependency{"test", common.Version{Major: 1, Minor: 0, Patch: 0, Tag: ""}, RequiresEqual, Depends}}
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
