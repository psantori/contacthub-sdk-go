/**
 * This file is part of contacthub-sdk-go.
 *
 * contacthub-sdk-go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * contacthub-sdk-go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with contacthub-sdk-go. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2017 Arduino AG
 *
 * @author Luca Osti
 *
 */

package enums

import (
	"encoding/json"
	"fmt"
)

var (
	_BringBackPropertyTypeNameToValue = map[string]BringBackPropertyType{
		"SESSION_ID":  SessionId,
		"EXTERNAL_ID": ExternalId,
	}

	_BringBackPropertyTypeValueToName = map[BringBackPropertyType]string{
		SessionId:  "SESSION_ID",
		ExternalId: "EXTERNAL_ID",
	}
)

func init() {
	var v BringBackPropertyType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BringBackPropertyTypeNameToValue = map[string]BringBackPropertyType{
			interface{}(SessionId).(fmt.Stringer).String():  SessionId,
			interface{}(ExternalId).(fmt.Stringer).String(): ExternalId,
		}
	}
}

// MarshalJSON is generated so BringBackPropertyType satisfies json.Marshaler.
func (r BringBackPropertyType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BringBackPropertyTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BringBackPropertyType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BringBackPropertyType satisfies json.Unmarshaler.
func (r *BringBackPropertyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BringBackPropertyType should be a string, got %s", data)
	}
	v, ok := _BringBackPropertyTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BringBackPropertyType %q", s)
	}
	*r = v
	return nil
}
