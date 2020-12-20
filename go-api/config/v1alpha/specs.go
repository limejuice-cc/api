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

import "io"

// ConfigStoreFormat specifies a configuration store format
type ConfigStoreFormat int

const (
	configStoreFormatNotSet ConfigStoreFormat = iota
	// YAMLFormat indicates a yaml configuration file
	YAMLFormat
)

func (f ConfigStoreFormat) String() string {
	switch f {
	case YAMLFormat:
		return "yaml"
	default:
		return ""
	}
}

// ConfigStore is a generic interface for a configuration store
type ConfigStore interface {
	HasItem(namespace, key string) bool
	SetItem(namespace, key string, value interface{}) error
	GetItem(namespace, key string) (interface{}, error)

	GetString(namespace, key string) string
	GetStringSlice(namespace, key string) []string
	GetStringMap(namespace, key string) map[string]string

	GetBool(namespace, key string) bool
	GetInt(namespace, key string) int
	GetFloat(namespace, key string) float64

	Load(r io.Reader, format ConfigStoreFormat) error
	Save(w io.Writer, format ConfigStoreFormat) error
}

// ConfigStoreProvider is a generic interface that provides a ConfigStore
type ConfigStoreProvider interface {
	Initialize(options ...ConfigStoreProviderOption) error
}

// ConfigStoreProviderOption is a option for a BuildRequestProvider
type ConfigStoreProviderOption interface {
	Apply(ConfigStoreProviderOption) error
}
