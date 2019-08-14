// Copyright 2019 simia.tech UG (haftungsbeschränkt)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enum

import (
	"strconv"
	"strings"
)

// Enum defines an extendable uint enum type.
type Enum uint

var (
	names          = make(map[Enum]string)
	enums          = make(map[string]Enum)
	lowerCaseEnums = make(map[string]Enum)
	next           = Enum(0)
)

// New returns a new enum within the scope.
func New(name string) Enum {
	if enum, ok := enums[name]; ok {
		return enum
	}

	enum := next
	names[enum] = name
	enums[name] = enum
	lowerCaseEnums[strings.ToLower(name)] = enum
	next++
	return enum
}

// Parse parses an enum from the provided string.
func Parse(raw string) (Enum, error) {
	enum, ok := enums[strings.TrimSpace(raw)]
	if !ok {
		return 0, ErrNoSuchEnum
	}
	return enum, nil
}

// ParseIgnoreCase parses an enum from the provided string ignoring the case.
func ParseIgnoreCase(raw string) (Enum, error) {
	enum, ok := lowerCaseEnums[strings.ToLower(strings.TrimSpace(raw))]
	if !ok {
		return 0, ErrNoSuchEnum
	}
	return enum, nil
}

func (e Enum) String() string {
	name, ok := names[e]
	if !ok {
		return "unknown name for enum " + strconv.Itoa(int(e))
	}
	return name
}
