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

// ConfigStoreFormat specifies a configuration store format
type ConfigStoreFormat int

const (
	_ ConfigStoreFormat = iota
	// YAMLFormat indicates a yaml configuration file
	YAMLFormat
)

var configStoreFormatValues = helper.EnumeratorValues{
	"yaml": YAMLFormat,
}

// String implements the Stringer interface.
func (f ConfigStoreFormat) String() string {
	return configStoreFormatValues.AsString(f)
}

// ParseConfigStoreFormat attempts to convert a string to a ConfigStoreFormat
func ParseConfigStoreFormat(name string) (ConfigStoreFormat, error) {
	x, err := configStoreFormatValues.Parse(name)
	if err != nil {
		return ConfigStoreFormat(0), err
	}
	return x.(ConfigStoreFormat), nil
}

// MarshalText implements the text marshaller method
func (f ConfigStoreFormat) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (f *ConfigStoreFormat) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseConfigStoreFormat(name)
	if err != nil {
		return err
	}
	*f = tmp
	return nil
}
