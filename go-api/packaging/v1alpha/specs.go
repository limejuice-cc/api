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

// Dependency is a dependant package
type Dependency struct {
	Name         PackageName    `yaml:"name"`         // Name is the name of the dependant package
	Version      common.Version `yaml:"version,flow"` // Version is the dependant package version
	Requires     Required       `yaml:"requires"`     // Requires specifies the required version of the dependant package
	Relationship Relationship   `yaml:"relation"`     // Relationship is the relationship of the package to the dependant package
}

// Dependencies is a list of dependant packages
type Dependencies []*Dependency

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
