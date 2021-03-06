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
	_SubscriptionKindNameToValue = map[string]SubscriptionKind{
		"DIGITAL_MESSAGE": DigitalMessage,
		"SERVICE":         Service,
		"OTHER":           OtherSubscription,
	}

	_SubscriptionKindValueToName = map[SubscriptionKind]string{
		DigitalMessage:    "DIGITAL_MESSAGE",
		Service:           "SERVICE",
		OtherSubscription: "OTHER",
	}
)

func init() {
	var v SubscriptionKind
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SubscriptionKindNameToValue = map[string]SubscriptionKind{
			interface{}(DigitalMessage).(fmt.Stringer).String():    DigitalMessage,
			interface{}(Service).(fmt.Stringer).String():           Service,
			interface{}(OtherSubscription).(fmt.Stringer).String(): OtherSubscription,
		}
	}
}

// MarshalJSON is generated so SubscriptionKind satisfies json.Marshaler.
func (r SubscriptionKind) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SubscriptionKindValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid SubscriptionKind: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so SubscriptionKind satisfies json.Unmarshaler.
func (r *SubscriptionKind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("SubscriptionKind should be a string, got %s", data)
	}
	v, ok := _SubscriptionKindNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid SubscriptionKind %q", s)
	}
	*r = v
	return nil
}
