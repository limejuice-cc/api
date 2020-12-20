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
	"os"

	common "github.com/limejuice-cc/api/go-api/common/v1alpha"
	pkg "github.com/limejuice-cc/api/go-api/packaging/v1alpha"
	plug "github.com/limejuice-cc/api/go-api/plugins/v1alpha"
)

// DockerBuildRequest represents a request a build request using docker
type DockerBuildRequest struct {
	Dockerfile     string               `yaml:"dockerfile"`             // Dockerfile is the contents of the Dockerfile
	DockerIgnore   string               `yaml:"dockerignore,omitempty"` // DockerIgnore is the contents of the .dockerignore file
	Tags           []string             `yaml:"tags,omitempty"`         // Tags are tags to apply to the built docker image
	BuildArgs      map[string]string    `yaml:"buildargs,omitempty"`    // BuildArgs are arguments to pass while building the docker image
	ExtraFiles     common.EmbeddedFiles `yaml:"files, omitempty"`       // ExtraFiles are files to include in the docker build process
	BuildDirectory string               `yaml:"buildDirectory"`         // BuildDirectory is the output directory where built files are generated
}

// BuiltFile represents a built file
type BuiltFile interface {
	Name() string
	User() string
	Group() string
	Body() []byte
	Size() int
	Mode() os.FileMode
	Type() pkg.FileType
}

// BuildRequest is a interface to request a build
type BuildRequest interface {
	Run(options ...BuildRequestOption) ([]BuiltFile, error)
}

// BuildRequestOption is a option for a BuildRequest
type BuildRequestOption interface {
	Apply(BuildRequest) error
}

// BuildContext is an interface that acts as glue between BuildRequest and BuildRequestContext
type BuildContext interface {
	Architecture() common.Architecture
	OperatingSystem() common.OperatingSystem
	Request() BuildRequest
}

// BuildRequestProvider is a provider that processes build requests
type BuildRequestProvider interface {
	plug.LimePlugin
	Initialize(options ...BuildRequestProviderOption) error
	Run(ctx BuildContext) ([]BuiltFile, error)
}

// BuildRequestProviderOption is a option for a BuildRequestProvider
type BuildRequestProviderOption interface {
	Apply(BuildRequestProvider) error
}
