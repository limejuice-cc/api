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

package helper

import (
	"fmt"
	"strings"

	"github.com/limejuice-cc/api/pkg/limejuiceerrors"
)

// EnumeratorValues is an enumerator abstraction
type EnumeratorValues map[string]interface{}

// Parse is a helper function that parses an enumerator
func (e EnumeratorValues) Parse(in string) (interface{}, error) {
	for k, v := range e {
		if strings.EqualFold(k, in) {
			return v, nil
		}
	}

	return nil, newEnumParseError(in)
}

// AsString returns an enumerator a a string
func (e EnumeratorValues) AsString(out interface{}) string {
	for k, v := range e {
		if v == out {
			return k
		}
	}
	return ""
}

// EnumParseError is a error that occurs when enum parsing fails
type EnumParseError struct {
	limejuiceerrors.LimeJuiceError
}

func newEnumParseError(in string) error {
	err := &EnumParseError{}
	err.Message = fmt.Sprintf("cannot parse %s", in)
	return err
}
