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

// FileType is the type of virtual file
type FileType int

const (
	_ FileType = iota
	// RegularFileType is a regular file type
	RegularFileType
	// DirectoryFileType is a directory file type
	DirectoryFileType
)

var fileTypeValues = helper.EnumeratorValues{
	"file": RegularFileType,
	"dir":  DirectoryFileType,
}

// String implements the Stringer interface.
func (t FileType) String() string {
	if t == FileType(0) {
		return RegularFileType.String()
	}
	return fileTypeValues.AsString(t)
}

// ParseFileType attempts to convert a string to a FileType
func ParseFileType(name string) (FileType, error) {
	if name == "" {
		return RegularFileType, nil
	}
	x, err := fileTypeValues.Parse(name)
	if err != nil {
		return FileType(0), err
	}
	return x.(FileType), nil
}

// MarshalText implements the text marshaller method
func (t FileType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (t *FileType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseFileType(name)
	if err != nil {
		return err
	}
	*t = tmp
	return nil
}
