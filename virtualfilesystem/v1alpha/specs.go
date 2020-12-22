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
	"io"
	"os"
	"time"
)

// Owner is a virtual owner of a file
type Owner interface {
	User() string
	UID() int
	RealUID() int

	Group() string
	GID() int
	RealGID() int
}

// File is a generic interface representing a virtual file
type File interface {
	Name() string
	Dir() string
	Path() string
	Type() FileType
	Size() int64
	FileMode() os.FileMode
	Owner() Owner
	ATime() time.Time
	MTime() time.Time
	FileInfo() os.FileInfo
}

// ReadableFile is a generic interface for a readable file
type ReadableFile interface {
	File
	io.Reader
	io.Seeker
	io.Closer
}

// WritableFile is a generic interface for a writable file
type WritableFile interface {
	ReadableFile
	io.Writer
}

// Directory is a generic interface representing a virtual directory
type Directory interface {
	File
	Files() []File
}

// VirtualFileSystemProvider is a generic interface to a provider hosting a virtual filesystem
type VirtualFileSystemProvider interface {
	Initialize(options ...VirtualFileSystemProviderOption) error

	AddOwner(owner Owner) error

	Open(name string) (ReadableFile, error)
	Create(name string) (WritableFile, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
	OpenFile(name string, flag int, perm os.FileMode) (WritableFile, error)
	Truncate(name string, size int64) error

	Lstat(name string) (os.FileInfo, error)
	Stat(name string) (os.FileInfo, error)

	ReadDir(dirname string) ([]os.FileInfo, error)
	ReadFile(filename string) ([]byte, error)

	Mkdir(name string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error

	Chmod(name string, mode os.FileMode) error
	Chown(name string, uid, gid int) error
	Chtimes(name string, atime, mtime time.Time) error

	Close() error
}

// VirtualFileSystemProviderOption is an option when initalizing a VirtualFileSystemProvider
type VirtualFileSystemProviderOption interface {
	Apply(VirtualFileSystemProvider) error
}

// VirtualFileSystemProviderOnDiskOption creates a virtual filesystem that is persisted on disk
type VirtualFileSystemProviderOnDiskOption interface {
	VirtualFileSystemProviderOption
}

// VirtualFileSystemProviderInMemory creates a virtual filesystem that is created in memory
type VirtualFileSystemProviderInMemory interface {
	VirtualFileSystemProviderOption
}
