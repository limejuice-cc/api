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

// *** Relationship ***

// Relationship describes the type of relationship between a  package and a dependency
type Relationship int

const (
	_ Relationship = iota
	// Suggests indicates the package suggests the dependency
	Suggests
	// Recommends indicates that the package recommends the dependency
	Recommends
	// Depends indicates that the package depends on the dependency
	Depends
	// Predepends indicates that the packate depends on the dependency and that the dependency needs to be installed first
	Predepends
	// Breaks indicates that the package breaks the dependency
	Breaks
	// Conflicts indicates that that package conflicts with the dependency
	Conflicts
	// Provides indicates that the package provides the dependency
	Provides
	// Replaces indicates that he package replaces the dependency
	Replaces
)

var relationshipValues = helper.EnumeratorValues{
	"breaks":     Breaks,
	"conflicts":  Conflicts,
	"depends":    Depends,
	"predepends": Predepends,
	"provides":   Provides,
	"recommends": Recommends,
	"replaces":   Replaces,
	"suggests":   Suggests,
}

// String implements the Stringer interface.
func (r Relationship) String() string {
	return relationshipValues.AsString(r)
}

// ParseRelationship attempts to convert a string to a Relationship
func ParseRelationship(name string) (Relationship, error) {
	x, err := relationshipValues.Parse(name)
	if err != nil {
		return Relationship(0), err
	}
	return x.(Relationship), nil
}

// MarshalText implements the text marshaller method
func (r Relationship) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (r *Relationship) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseRelationship(name)
	if err != nil {
		return err
	}
	*r = tmp
	return nil
}

// *** Required ***

// Required indicates the relationship the required version of the dependency
type Required int

const (
	_ Required = 0
	//RequiresEqual indicates that the package requires the dependency to have a version equal to
	RequiresEqual Required = 1 << 0
	//RequiresGreaterThan indicates that the package requires the dependency to have a version greater than
	RequiresGreaterThan Required = 1 << 1
	//RequiresLessThan indicates that the package requires the dependency to have a version less than
	RequiresLessThan Required = 1 << 2
	//RequiresGreaterThanEqual indicates that the package requires the dependency to have a version greater than or equal to
	RequiresGreaterThanEqual Required = RequiresEqual | RequiresGreaterThan
	//RequiresLessThanEqual indicates that the package requires the dependency to have a version less than or equal to
	RequiresLessThanEqual Required = RequiresEqual | RequiresLessThan
)

var requiredValues = helper.EnumeratorValues{
	"==": RequiresEqual,
	">>": RequiresGreaterThan,
	">=": RequiresGreaterThanEqual,
	"<<": RequiresLessThan,
	"<=": RequiresLessThanEqual,
}

// String implements the Stringer interface.
func (r Required) String() string {
	return requiredValues.AsString(r)
}

// ParseRequired attempts to convert a string to a Required
func ParseRequired(name string) (Required, error) {
	x, err := requiredValues.Parse(name)
	if err != nil {
		return Required(0), err
	}
	return x.(Required), nil
}

// MarshalText implements the text marshaller method
func (r Required) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (r *Required) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseRequired(name)
	if err != nil {
		return err
	}
	*r = tmp
	return nil
}

// *** FileType ***

// FileType specifies the type of a package file
type FileType int

const (
	_ FileType = iota
	// ConfigurationFile indicates that the file is a configuration file
	ConfigurationFile
	// ExecutableFile indicates that the file is an executable
	ExecutableFile
	// DataFile indicates that the file is a data file
	DataFile
	// OtherFile indicates that the file is not categorized
	OtherFile
)

var fileTypeValues = helper.EnumeratorValues{
	"config": ConfigurationFile,
	"exec":   ExecutableFile,
	"data":   DataFile,
	"other":  OtherFile,
}

// String implements the Stringer interface.
func (f FileType) String() string {
	if f == FileType(0) {
		return OtherFile.String()
	}
	return fileTypeValues.AsString(f)
}

// ParseFileType attempts to convert a string to a FileType
func ParseFileType(name string) (FileType, error) {
	if name == "" {
		return OtherFile, nil
	}
	x, err := fileTypeValues.Parse(name)
	if err != nil {
		return FileType(0), err
	}
	return x.(FileType), nil
}

// MarshalText implements the text marshaller method
func (f FileType) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (f *FileType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseFileType(name)
	if err != nil {
		return err
	}
	*f = tmp
	return nil
}

// *** ActionType ***

// ActionType specifies the type of action
type ActionType int

const (
	_ ActionType = iota
	// Install specifies the installation action
	Install
	// Reconfigure specifies the reconfigure action
	Reconfigure
	// Upgrade specifies the upgrade action
	Upgrade
	// Remove specifies the remove action
	Remove
	// Purge specifies the purge action
	Purge
)

var actionTypeValues = helper.EnumeratorValues{
	"install":     Install,
	"reconfigure": Reconfigure,
	"upgrade":     Upgrade,
	"remove":      Remove,
	"purge":       Purge,
}

// String implements the Stringer interface.
func (t ActionType) String() string {
	return actionTypeValues.AsString(t)
}

// ParseActionType attempts to convert a string to a ActionType
func ParseActionType(name string) (ActionType, error) {
	x, err := actionTypeValues.Parse(name)
	if err != nil {
		return ActionType(0), err
	}
	return x.(ActionType), nil
}

// MarshalText implements the text marshaller method
func (t ActionType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (t *ActionType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseActionType(name)
	if err != nil {
		return err
	}
	*t = tmp
	return nil
}
