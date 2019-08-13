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
	"bytes"
)

// MarshalJSON implements json.Marshaler.
func (e Enum) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Enum) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, `" `)
	if len(data) == 0 || bytes.Equal(data, []byte(`null`)) {
		return nil
	}

	enum, err := ParseIgnoreCase(string(data))
	if err != nil {
		return err
	}
	*e = enum

	return nil
}
