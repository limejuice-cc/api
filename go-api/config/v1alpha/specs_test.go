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

func TestParseConfigStoreFormat(t *testing.T) {
	var testValues = []struct {
		value   string
		outcome ConfigStoreFormat
	}{
		{"yaml", YAMLFormat},
	}

	for _, v := range testValues {
		a, err := ParseConfigStoreFormat(v.value)
		if assert.NoError(t, err) {
			assert.Equal(t, v.outcome, a)
			assert.Equal(t, v.value, a.String())
		}
	}

	var format ConfigStoreFormat
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &format))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &format))
	assert.NoError(t, yaml.Unmarshal([]byte("yaml"), &format))

	_, err := yaml.Marshal(&format)
	assert.NoError(t, err)
}
