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

func TestParseFileType(t *testing.T) {
	var testValues = []struct {
		in       string
		out      FileType
		hasError bool
	}{
		{in: "file", out: RegularFileType, hasError: false},
		{in: "dir", out: DirectoryFileType, hasError: false},
	}

	for _, tst := range testValues {
		assert.Equal(t, tst.in, tst.out.String())
		v, err := ParseFileType(tst.in)
		if tst.hasError {
			assert.Error(t, err)
			continue
		}
		assert.Equal(t, tst.out, v)
	}

	var ft FileType
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &ft))
	assert.NoError(t, yaml.Unmarshal([]byte(""), &ft))
	assert.NoError(t, yaml.Unmarshal([]byte("file"), &ft))
	_, err := yaml.Marshal(&ft)
	assert.NoError(t, err)

	_, err = ParseFileType(FileType(0).String())
	assert.NoError(t, err)
	_, err = ParseFileType("")
	assert.NoError(t, err)
}
