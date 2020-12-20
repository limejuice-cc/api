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
	"time"

	common "github.com/limejuice-cc/api/go-api/common/v1alpha"
)

// LimePluginType is the type of lime plugin
type LimePluginType int

const (
	limePluginNotSet LimePluginType = iota
	// GenericFileGenerator is a generic plugin that generates files
	GenericFileGenerator
	// CertificateGenerator is a plugin that generates certificates
	CertificateGenerator
	// CommandProxy is a plugin that acts as a proxy for executing commands
	CommandProxy
)

// LimePlugin is a generic interface for lime plugins
type LimePlugin interface {
	Name() string
	Description() string
	Version() common.Version
	BuildDate() time.Time
	Type() LimePluginType
}
