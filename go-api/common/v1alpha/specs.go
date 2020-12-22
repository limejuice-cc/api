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

// Architectures is a list of architectures
type Architectures []Architecture

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

// MarshalText implements the text marshaller method
func (v *Version) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (v *Version) UnmarshalText(text []byte) error {
	value := string(text)
	tmp, err := ParseVersion(value)
	if err != nil {
		return err
	}
	v = tmp
	return nil
}

// EmbeddedFileContents represents the contents of an embedded file
type EmbeddedFileContents []byte

// UnmarshalText implements the text unmarshaller method
func (c *EmbeddedFileContents) UnmarshalText(text []byte) (err error) {
	out := make([]byte, len(text))
	written, _, err := ascii85.Decode(out, text, true)
	if err == nil {
		*c = out[:written]
	}
	return
}

// MarshalText implements the text marshaller method
func (c EmbeddedFileContents) MarshalText() ([]byte, error) {
	out := make([]byte, ascii85.MaxEncodedLen(len(c)))
	written := ascii85.Encode(out, []byte(c))
	return out[:written], nil
}

// EmbeddedFiles is a lit of embedded files
type EmbeddedFiles map[string]EmbeddedFileContents
