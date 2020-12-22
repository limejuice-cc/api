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

// *** ArchitectureVariant ***

// ArchitectureVariant represents a variant in the target system's architecture
type ArchitectureVariant int

const (
	_ ArchitectureVariant = iota
	// NoVariant indicates that there is no variant in the architecture
	NoVariant
)

var architectureVariantValues = helper.EnumeratorValues{
	"none": NoVariant,
}

// String implements the Stringer interface.
func (v ArchitectureVariant) String() string {
	if v == ArchitectureVariant(0) {
		return NoVariant.String()
	}
	return architectureVariantValues.AsString(v)
}

// ParseArchitectureVariant attempts to convert a string to a ArchitectureVariant
func ParseArchitectureVariant(name string) (ArchitectureVariant, error) {
	if name == "" {
		return NoVariant, nil
	}
	x, err := architectureVariantValues.Parse(name)
	if err != nil {
		return ArchitectureVariant(0), err
	}
	return x.(ArchitectureVariant), nil
}

// MarshalText implements the text marshaller method
func (v ArchitectureVariant) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (v *ArchitectureVariant) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseArchitectureVariant(name)
	if err != nil {
		return err
	}
	*v = tmp
	return nil
}

// *** Architecture **

// Architecture represents a target system's architecture
type Architecture int

const (
	_ Architecture = iota
	// AMD64 represents the x80_64 architecture
	AMD64
)

var architectureValues = helper.EnumeratorValues{
	"amd64": AMD64,
}

// String implements the Stringer interface.
func (a Architecture) String() string {
	if a == Architecture(0) {
		return AMD64.String()
	}
	return architectureValues.AsString(a)
}

// ParseArchitecture attempts to convert a string to a Architecture
func ParseArchitecture(name string) (Architecture, error) {
	if name == "" {
		return AMD64, nil
	}
	x, err := architectureValues.Parse(name)
	if err != nil {
		return Architecture(0), err
	}
	return x.(Architecture), nil
}

// MarshalText implements the text marshaller method
func (a Architecture) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (a *Architecture) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseArchitecture(name)
	if err != nil {
		return err
	}
	*a = tmp
	return nil
}

// Variant returns the architecture's Variant
func (a Architecture) Variant() ArchitectureVariant {
	switch a {
	case AMD64:
		return NoVariant
	default:
		return NoVariant
	}
}

// *** OperatingSystem

// OperatingSystem specifies the operating system
type OperatingSystem int

const (
	_ OperatingSystem = iota
	// Linux represents the linux operating system
	Linux
)

var operatingSystemValues = helper.EnumeratorValues{
	"linux": Linux,
}

// String implements the Stringer interface.
func (o OperatingSystem) String() string {
	if o == OperatingSystem(0) {
		return Linux.String()
	}
	return operatingSystemValues.AsString(o)
}

// ParseOperatingSystem attempts to convert a string to a Architecture
func ParseOperatingSystem(name string) (OperatingSystem, error) {
	if name == "" {
		return Linux, nil
	}
	x, err := operatingSystemValues.Parse(name)
	if err != nil {
		return OperatingSystem(0), err
	}
	return x.(OperatingSystem), nil
}

// MarshalText implements the text marshaller method
func (o OperatingSystem) MarshalText() ([]byte, error) {
	return []byte(o.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (o *OperatingSystem) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseOperatingSystem(name)
	if err != nil {
		return err
	}
	*o = tmp
	return nil
}
