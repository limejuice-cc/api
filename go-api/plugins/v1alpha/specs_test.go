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

func TestLimePluginType(t *testing.T) {
	var testValues = []struct {
		in       string
		out      LimePluginType
		hasError bool
	}{
		{in: "Builder", out: Builder, hasError: false},
		{in: "CertificateGenerator", out: CertificateGenerator, hasError: false},
		{in: "CommandProxy", out: CommandProxy, hasError: false},
		{in: "ConfigStore", out: ConfigStore, hasError: false},
		{in: "GenericFileGenerator", out: GenericFileGenerator, hasError: false},
		{in: "", out: LimePluginType(0), hasError: true},
	}

	for _, tst := range testValues {
		assert.Equal(t, tst.in, tst.out.String())
		v, err := ParseLimePluginType(tst.in)
		if tst.hasError {
			assert.Error(t, err)
			continue
		}
		assert.Equal(t, tst.out, v)
	}

	var lpt LimePluginType
	assert.Error(t, yaml.Unmarshal([]byte("[1,2,3]"), &lpt))
	assert.Error(t, yaml.Unmarshal([]byte("unknown"), &lpt))
	assert.NoError(t, yaml.Unmarshal([]byte("Builder"), &lpt))
	_, err := yaml.Marshal(&lpt)
	assert.NoError(t, err)
}
