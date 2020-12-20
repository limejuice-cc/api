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
	"encoding/ascii85"
	"fmt"
	"strconv"
	"strings"
)

// Architecture represents a target system's architecture
type Architecture int

const (
	architectureNotSet Architecture = iota
	// AMD64 represents the x80_64 architecture
	AMD64
)

func (a Architecture) String() string {
	switch a {
	case AMD64:
		return "amd64"
	default:
		return "amd64"
	}
}

// ParseArchitecture parses an architecture
func ParseArchitecture(v string) (Architecture, error) {
	switch v {
	case "amd64":
		return AMD64, nil
	case "":
		return AMD64, nil
	default:
		return architectureNotSet, fmt.Errorf("unknown architecture %s", v)
	}
}

// UnmarshalYAML implements custom unmarshal for Architecture
func (a *Architecture) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var arch string
	if err = unmarshal(&arch); err != nil {
		return
	}

	*a, err = ParseArchitecture(arch)
	return
}

// MarshalYAML implements custom marshalling for Architecture
func (a Architecture) MarshalYAML() (interface{}, error) {
	return a.String(), nil
}

// Split splits an architecture into a base architecture and a variant. This is used with docker's platform definition
func (a Architecture) Split() (string, string) {
	switch a {
	case AMD64:
		return "amd64", ""
	default:
		return "amd64", ""
	}
}

// Architectures is a list of architectures
type Architectures []Architecture

// OperatingSystem specifies the operating system
type OperatingSystem int

const (
	noOperatingSystemSet OperatingSystem = iota
	// Linux represents the linux operating system
	Linux
)

func (o OperatingSystem) String() string {
	switch o {
	case Linux:
		return "linux"
	default:
		return "linux"
	}
}

// ParseOperatingSystem parses an OperatingSystem
func ParseOperatingSystem(v string) (OperatingSystem, error) {
	switch v {
	case "linux":
		return Linux, nil
	case "":
		return Linux, nil
	default:
		return noOperatingSystemSet, fmt.Errorf("unknown operating system %s", v)
	}
}

// UnmarshalYAML implements custom unmarshal for OperatingSystem
func (o *OperatingSystem) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var os string
	if err = unmarshal(&os); err != nil {
		return
	}

	*o, err = ParseOperatingSystem(os)
	return
}

// MarshalYAML implements custom marshalling for OperatingSystem
func (o OperatingSystem) MarshalYAML() (interface{}, error) {
	return o.String(), nil
}

// Version represents the version of a lime package
type Version struct {
	Major int    // Major is the package's major version
	Minor int    // Minor is the package's minor version
	Patch int    // Patch is the package's patch version
	Tag   string // Tag is the package version's tag
}

func (v *Version) String() string {
	o := fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Tag != "" {
		return fmt.Sprintf("%s-%s", o, v.Tag)
	}
	return o
}

// ParseVersion parses a version
func ParseVersion(v string) (*Version, error) {
	if strings.HasPrefix(v, "v") {
		v = v[1:]
	}
	if v == "" {
		return nil, fmt.Errorf("invalid version %s", v)
	}

	parsed := &Version{}
	parts := strings.SplitN(v, "-", 2)
	if len(parts) > 1 {
		parsed.Tag = parts[1]
	}

	parts = strings.Split(parts[0], ".")

	major, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}
	parsed.Major = int(major)

	if len(parts) < 2 {
		return parsed, nil
	}
	minor, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}
	parsed.Minor = int(minor)

	if len(parts) < 3 {
		return parsed, nil
	}

	patch, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}
	parsed.Patch = int(patch)

	return parsed, nil
}

// UnmarshalYAML implements custom unmarshal for version
func (v *Version) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var version string
	if err = unmarshal(&version); err != nil {
		return
	}

	v, err = ParseVersion(version)
	return
}

// MarshalYAML implements custom marshalling for version
func (v Version) MarshalYAML() (interface{}, error) {
	return v.String(), nil
}

// EmbeddedFileContents represents the contents of an embedded file
type EmbeddedFileContents []byte

// UnmarshalYAML implements custom unmarshal for EmbeddedFileContents
func (c *EmbeddedFileContents) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var contents string
	if err = unmarshal(&contents); err != nil {
		return
	}

	out := make([]byte, len(contents))
	written, _, err := ascii85.Decode(out, []byte(contents), true)
	if err == nil {
		*c = out[:written]
	}

	return
}

// MarshalYAML implements custom marshalling for EmbeddedFileContents
func (c EmbeddedFileContents) MarshalYAML() (interface{}, error) {
	out := make([]byte, ascii85.MaxEncodedLen(len(c)))
	written := ascii85.Encode(out, []byte(c))
	return string(out[:written]), nil
}

// EmbeddedFiles is a lit of embedded files
type EmbeddedFiles map[string]EmbeddedFileContents
