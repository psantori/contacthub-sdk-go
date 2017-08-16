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
	_NotificationServiceTypeNameToValue = map[string]NotificationServiceType{
		"APN": APN,
		"GCM": GCM,
		"WNS": WNS,
		"ADM": ADM,
		"SNS": SNS,
	}

	_NotificationServiceTypeValueToName = map[NotificationServiceType]string{
		APN: "APN",
		GCM: "GCM",
		WNS: "WNS",
		ADM: "ADM",
		SNS: "SNS",
	}
)

func init() {
	var v NotificationServiceType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_NotificationServiceTypeNameToValue = map[string]NotificationServiceType{
			interface{}(APN).(fmt.Stringer).String(): APN,
			interface{}(GCM).(fmt.Stringer).String(): GCM,
			interface{}(WNS).(fmt.Stringer).String(): WNS,
			interface{}(ADM).(fmt.Stringer).String(): ADM,
			interface{}(SNS).(fmt.Stringer).String(): SNS,
		}
	}
}

// MarshalJSON is generated so NotificationServiceType satisfies json.Marshaler.
func (r NotificationServiceType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _NotificationServiceTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid NotificationServiceType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so NotificationServiceType satisfies json.Unmarshaler.
func (r *NotificationServiceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NotificationServiceType should be a string, got %s", data)
	}
	v, ok := _NotificationServiceTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid NotificationServiceType %q", s)
	}
	*r = v
	return nil
}
