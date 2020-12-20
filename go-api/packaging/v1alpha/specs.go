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
	"fmt"
	"strings"
	"time"

	common "github.com/limejuice-cc/api/go-api/common/v1alpha"
)

// PackageName represents the name of a package
type PackageName string

const (
	validPackageRunes = "abcdefghijklmnopqrstuvwxyz0123456789-_"
)

// Valid checks if a package name is valid
func (n PackageName) Valid() error {
	notValid := func(c rune) bool { return !strings.ContainsAny(validPackageRunes, string(c)) }
	words := strings.FieldsFunc(string(n), notValid)
	if string(n) != strings.Join(words, "") {
		return fmt.Errorf("invalid package name %s can only contain %s", n, validPackageRunes)
	}
	return nil
}

// UnmarshalYAML implements custom unmarshal for version
func (n *PackageName) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var name string
	if err = unmarshal(&name); err != nil {
		return
	}

	if err = PackageName(name).Valid(); err != nil {
		return
	}
	*n = PackageName(name)
	return
}

// MarshalYAML implements custom marshalling for PackageName
func (n PackageName) MarshalYAML() (interface{}, error) {
	return string(n), nil
}

// MetadataItem represents an non-categorized metadata item
type MetadataItem struct {
	Key   string `yaml:"key"`   // Key is the metadata item's key
	Value string `yaml:"value"` // Value is the metadata item's value
}

// Metadata represents package metadata
type Metadata struct {
	Description   string               `yaml:"description,omitempty"` // Description is an optional description of the package
	Architectures common.Architectures `yaml:"arch,omitempty"`        // Architectures is an optional list of architectures
	Items         []*MetadataItem      `yaml:"items,omitempty"`       // Items are additional metadata items
}

// Relationship describes the type of relationship between a  package and a dependency
type Relationship int

const (
	noRelationshipSet Relationship = iota
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

// ParseRelationship parses a Relationship
func ParseRelationship(v string) (Relationship, error) {
	switch v {
	case "suggests":
		return Suggests, nil
	case "recommends":
		return Recommends, nil
	case "depends":
		return Depends, nil
	case "predepends":
		return Predepends, nil
	case "breaks":
		return Breaks, nil
	case "conflicts":
		return Conflicts, nil
	case "provides":
		return Provides, nil
	case "replaces":
		return Replaces, nil
	default:
		return noRelationshipSet, fmt.Errorf("unrecognized relationship %s", v)
	}
}

func (r Relationship) String() string {
	switch r {
	case Suggests:
		return "suggests"
	case Recommends:
		return "recommends"
	case Depends:
		return "depends"
	case Predepends:
		return "predepends"
	case Breaks:
		return "breaks"
	case Conflicts:
		return "conflicts"
	case Provides:
		return "provides"
	case Replaces:
		return "replaces"
	default:
		return ""
	}
}

// UnmarshalYAML implements custom unmarshal for Relationship
func (r *Relationship) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var relation string
	if err = unmarshal(&relation); err != nil {
		return
	}

	*r, err = ParseRelationship(relation)
	return
}

// MarshalYAML implements custom marshalling for Relationship
func (r Relationship) MarshalYAML() (interface{}, error) {
	return r.String(), nil
}

// Required indicates the relationship the required version of the dependency
type Required int

const (
	requiredNotSet Required = 0
	//RequiresEqual indicates that the package requires the dependency to have a version equal to
	RequiresEqual = 1 << 0
	//RequiresGreaterThan indicates that the package requires the dependency to have a version greater than
	RequiresGreaterThan = 1 << 1
	//RequiresLessThan indicates that the package requires the dependency to have a version less than
	RequiresLessThan = 1 << 2
	//RequiresGreaterThanEqual indicates that the package requires the dependency to have a version greater than or equal to
	RequiresGreaterThanEqual = RequiresEqual | RequiresGreaterThan
	//RequiresLessThanEqual indicates that the package requires the dependency to have a version less than or equal to
	RequiresLessThanEqual = RequiresEqual | RequiresLessThan
)

func (r Required) String() string {
	switch r {
	case RequiresEqual:
		return "=="
	case RequiresGreaterThan:
		return ">>"
	case RequiresGreaterThanEqual:
		return ">="
	case RequiresLessThan:
		return "<<"
	case RequiresLessThanEqual:
		return "<="
	default:
		return ""
	}
}

// ParseRequired parses a dependency relationship
func ParseRequired(requires string) (Required, error) {
	switch requires {
	case "==":
		return RequiresEqual, nil
	case ">>":
		return RequiresGreaterThan, nil
	case ">=":
		return RequiresGreaterThanEqual, nil
	case "<<":
		return RequiresLessThan, nil
	case "<=":
		return RequiresLessThanEqual, nil
	default:
		return requiredNotSet, fmt.Errorf("unknown required relationship %s", requires)
	}
}

// UnmarshalYAML implements custom unmarshal for Required
func (r *Required) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var required string
	if err = unmarshal(&required); err != nil {
		return
	}

	*r, err = ParseRequired(required)
	return
}

// MarshalYAML implements custom marshalling for Required
func (r Required) MarshalYAML() (interface{}, error) {
	return r.String(), nil
}

// Dependency is a dependant package
type Dependency struct {
	Name         PackageName    `yaml:"name"`         // Name is the name of the dependant package
	Version      common.Version `yaml:"version,flow"` // Version is the dependant package version
	Requires     Required       `yaml:"requires"`     // Requires specifies the required version of the dependant package
	Relationship Relationship   `yaml:"relation"`     // Relationship is the relationship of the package to the dependant package
}

// Dependencies is a list of dependant packages
type Dependencies []*Dependency

// FileType specifies the type of a package file
type FileType int

const (
	noFileTypeSet FileType = iota
	// ConfigurationFile indicates that the file is a configuration file
	ConfigurationFile
	// ExecutableFile indicates that the file is an executable
	ExecutableFile
	// DataFile indicates that the file is a data file
	DataFile
	// OtherFile indicates that the file is not categorized
	OtherFile
)

func (t FileType) String() string {
	switch t {
	case ConfigurationFile:
		return "config"
	case ExecutableFile:
		return "exec"
	case DataFile:
		return "data"
	case OtherFile:
		return "other"
	default:
		return ""
	}
}

// ParseFileType parses a package file type
func ParseFileType(v string) (FileType, error) {
	switch v {
	case "config":
		return ConfigurationFile, nil
	case "exec":
		return ExecutableFile, nil
	case "data":
		return DataFile, nil
	case "other":
		return OtherFile, nil
	case "":
		return OtherFile, nil
	default:
		return noFileTypeSet, fmt.Errorf("unknown file type %s", v)
	}
}

// UnmarshalYAML implements custom unmarshal for FileType
func (t *FileType) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var fileType string
	if err = unmarshal(&fileType); err != nil {
		return
	}

	*t, err = ParseFileType(fileType)
	return
}

// MarshalYAML implements custom marshalling for FileType
func (t FileType) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

// File is a package file
type File struct {
	Path     string   `yaml:"path"`             // Path is full path of the file
	Type     FileType `yaml:"type"`             // Type is the package type of the file
	IsCommon bool     `yaml:"common,omitempty"` // IsCommon indicates the file is a common file
	SHA256   string   `yaml:"hash"`             // SHA256 hash is the SHA256 hash of the file
	User     string   `yaml:"user,omitempty"`   // User is the user who owns the file
	Group    string   `yaml:"group,omitempty"`  // Group is the group that owns the file
	Mode     int      `yaml:"mode,omitempty"`   // Mode is the mode of the file
}

// Files is a list of package file
type Files []*File

// ActionType specifies the type of action
type ActionType int

const (
	noActionTypeSet ActionType = iota
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

// ParseActionType parses an ActionType
func ParseActionType(v string) (ActionType, error) {
	switch v {
	case "install":
		return Install, nil
	case "reconfigure":
		return Reconfigure, nil
	case "upgrade":
		return Upgrade, nil
	case "remove":
		return Remove, nil
	case "purge":
		return Purge, nil
	default:
		return noActionTypeSet, fmt.Errorf("unknown action type %s", v)
	}
}

func (a ActionType) String() string {
	switch a {
	case Install:
		return "install"
	case Reconfigure:
		return "reconfigure"
	case Upgrade:
		return "upgrade"
	case Remove:
		return "remove"
	case Purge:
		return "purge"
	default:
		return ""
	}
}

// UnmarshalYAML implements custom unmarshal for ActionType
func (a *ActionType) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var actionType string
	if err = unmarshal(&actionType); err != nil {
		return
	}

	*a, err = ParseActionType(actionType)
	return
}

// MarshalYAML implements custom marshalling for ActionType
func (a ActionType) MarshalYAML() (interface{}, error) {
	return a.String(), nil
}

// ActionItem is a step within an action
type ActionItem struct {
	Values interface{} `yaml:"action"` // Values are the action values
}

// ActionItems are a list of ActionItem
type ActionItems []*ActionItem

// Action is a packaging action
type Action struct {
	Type   ActionType  `yaml:"type"`             // Type is the type of action
	Before ActionItems `yaml:"before,omitempty"` // Before are action items before the actiontype occurs
	After  ActionItems `yaml:"after,omitempty"`  // After are action items after the actiontype occurs
}

// Actions is a list of actions
type Actions []*Action

// Plugin desfines a plugin
type Plugin struct {
	Name PackageName `yaml:"name"` // Name is the name of the plugin
}

// Plugins is a lit of plugins
type Plugins []*Plugin

// Manifest describes the contents of a lime package
type Manifest struct {
	Name         PackageName    `yaml:"name"`               // Name is the name of the package
	Version      common.Version `yaml:"version,flow"`       // Version is the package version
	Created      time.Time      `yaml:"created"`            // Created is the datetime that the package was created
	Metadata     Metadata       `yaml:"metadata,omitempty"` // Metadata is package metadata
	Dependencies Dependencies   `yaml:"depends,omitempty"`  // Dependencies are depdenant packages
	Files        Files          `yaml:"files,omitempty"`    // Files are package files
	Actions      Actions        `yaml:"actions,omitempty"`  // Actions are package actions
	Triggers     Actions        `yaml:"triggers,omitempty"` // Triggers are actions triggered by other packages
	Plugins      Plugins        `yaml:"plugins,omitempty"`  // Plugins specifies the plugsins used by this package
}

const (
	// LimePackageMagic is the magic characters for a lime package
	LimePackageMagic string = "LiMedPkg"
)

// LimePackageFileIndexEntry is an entry in the lime package file index
type LimePackageFileIndexEntry struct {
	Path           string `yaml:"path"`       // Path is the file path
	Size           int64  `yaml:"size"`       // Size is the original file size
	CompressedSize int64  `yaml:"compressed"` // CompressedSize is the compressed file size
	FileOffset     int64  `yaml:"offset"`     // FileOffset is the offset of the file in the package
}

// LimePackageFileIndex is the file index for a lime package
type LimePackageFileIndex struct {
	Files []LimePackageFileIndexEntry `yaml:"files"` // Files are the entries in the file index
}

// RawLimePackageFile is a raw lime package file
type RawLimePackageFile []byte

// RawLimePackage is a raw lime package
type RawLimePackage struct {
	Magic          [8]byte
	ManifestLength [8]byte
	Manifest       []byte
	IndexLength    [8]byte
	Index          []byte
	Files          []byte
}
