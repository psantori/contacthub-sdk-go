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
	_MobileDeviceTypeNameToValue = map[string]MobileDeviceType{
		"IOS":           IOS,
		"ANDROID":       Android,
		"WINDOWS_PHONE": WindowsPhone,
		"FIREOS":        FireOS,
	}

	_MobileDeviceTypeValueToName = map[MobileDeviceType]string{
		IOS:          "IOS",
		Android:      "ANDROID",
		WindowsPhone: "WINDOWS_PHONE",
		FireOS:       "FIREOS",
	}
)

func init() {
	var v MobileDeviceType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MobileDeviceTypeNameToValue = map[string]MobileDeviceType{
			interface{}(IOS).(fmt.Stringer).String():          IOS,
			interface{}(Android).(fmt.Stringer).String():      Android,
			interface{}(WindowsPhone).(fmt.Stringer).String(): WindowsPhone,
			interface{}(FireOS).(fmt.Stringer).String():       FireOS,
		}
	}
}

// MarshalJSON is generated so MobileDeviceType satisfies json.Marshaler.
func (r MobileDeviceType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MobileDeviceTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid MobileDeviceType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so MobileDeviceType satisfies json.Unmarshaler.
func (r *MobileDeviceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MobileDeviceType should be a string, got %s", data)
	}
	v, ok := _MobileDeviceTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid MobileDeviceType %q", s)
	}
	*r = v
	return nil
}
