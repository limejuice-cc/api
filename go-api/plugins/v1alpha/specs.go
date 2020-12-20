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

func (t LimePluginType) String() string {
	switch t {
	case GenericFileGenerator:
		return "GenericFileGenerator"
	case CertificateGenerator:
		return "CertificateGenerator"
	case CommandProxy:
		return "CommandProxy"
	default:
		return ""
	}
}

// ParseLimePluginType parses a LimePluginType
func ParseLimePluginType(v string) (LimePluginType, error) {
	switch strings.ToLower(v) {
	case "genericfilegenerator":
		return GenericFileGenerator, nil
	case "certificategenerator":
		return CertificateGenerator, nil
	case "commandproxy":
		return CommandProxy, nil
	default:
		return limePluginNotSet, fmt.Errorf("unknown lime pluging type %s", v)
	}
}

// LimePlugin is a generic interface for lime plugins
type LimePlugin interface {
	Name() string
	Description() string
	Version() common.Version
	BuildDate() time.Time
	Type() LimePluginType
}
