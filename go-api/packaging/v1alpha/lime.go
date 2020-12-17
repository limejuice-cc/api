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

const (
	// LimePackageMagic is the magic characters for a limee package
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
