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
	"github.com/limejuice-cc/api/helper"
)

// LimePluginType is the type of lime plugin
type LimePluginType int

const (
	_ LimePluginType = iota
	// Builder is a plugin that builds executables
	Builder
	// CertificateGenerator is a plugin that generates certificates
	CertificateGenerator
	// CommandProxy is a plugin that acts as a proxy for executing commands
	CommandProxy
	// ConfigStore is a plugin that as as a configuration storage
	ConfigStore
	// GenericFileGenerator is a generic plugin that generates files
	GenericFileGenerator
)

var limePluginTypeValues = helper.EnumeratorValues{
	"Builder":              Builder,
	"CertificateGenerator": CertificateGenerator,
	"CommandProxy":         CommandProxy,
	"ConfigStore":          ConfigStore,
	"GenericFileGenerator": GenericFileGenerator,
}

// String implements the Stringer interface.
func (t LimePluginType) String() string {
	return limePluginTypeValues.AsString(t)
}

// ParseLimePluginType attempts to convert a string to a LimePluginType
func ParseLimePluginType(name string) (LimePluginType, error) {
	x, err := limePluginTypeValues.Parse(name)
	if err != nil {
		return LimePluginType(0), err
	}
	return x.(LimePluginType), nil
}

// MarshalText implements the text marshaller method
func (t LimePluginType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (t *LimePluginType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseLimePluginType(name)
	if err != nil {
		return err
	}
	*t = tmp
	return nil
}
