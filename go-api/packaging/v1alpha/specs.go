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
	"time"
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
	Description   string          `yaml:"description,omitempty"` // Description is an optional description of the package
	Architectures Architectures   `yaml:"arch,omitempty"`        // Architectures is an optional list of architectures
	Items         []*MetadataItem `yaml:"items,omitempty"`       // Items are additional metadata items
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
	Name         PackageName  `yaml:"name"`         // Name is the name of the dependant package
	Version      Version      `yaml:"version,flow"` // Version is the dependant package version
	Requires     Required     `yaml:"requires"`     // Requires specifies the required version of the dependant package
	Relationship Relationship `yaml:"relation"`     // Relationship is the relationship of the package to the dependant package
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
	Name         PackageName  `yaml:"name"`               // Name is the name of the package
	Version      Version      `yaml:"version,flow"`       // Version is the package version
	Created      time.Time    `yaml:"created"`            // Created is the datetime that the package was created
	Metadata     Metadata     `yaml:"metadata,omitempty"` // Metadata is package metadata
	Dependencies Dependencies `yaml:"depends,omitempty"`  // Dependencies are depdenant packages
	Files        Files        `yaml:"files,omitempty"`    // Files are package files
	Actions      Actions      `yaml:"actions,omitempty"`  // Actions are package actions
	Triggers     Actions      `yaml:"triggers,omitempty"` // Triggers are actions triggered by other packages
	Plugins      Plugins      `yaml:"plugins,omitempty"`  // Plugins specifies the plugsins used by this package
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

// DockerBuildRequest represents a request a build request using docker
type DockerBuildRequest struct {
	Dockerfile     string            `yaml:"dockerfile"`          // Dockerfile is the contents of the Dockerfile
	Tags           []string          `yaml:"tags,omitempty"`      // Tags are tags to apply to the built docker image
	BuildArgs      map[string]string `yaml:"buildargs,omitempty"` // BuildArgs are arguments to pass while building the docker image
	ExtraFiles     EmbeddedFiles     `yaml:"files, omitempty"`    // ExtraFiles are files to include in the docker build process
	BuildDirectory string            `yaml:"buildDirectory"`      // BuildDirectory is the directory within the docker image to extract built files
}
