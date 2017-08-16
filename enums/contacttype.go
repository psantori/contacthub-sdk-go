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
	_ContactTypeNameToValue = map[string]ContactType{
		"MOBILE": Mobile,
		"PHONE":  Phone,
		"EMAIL":  Email,
		"FAX":    Fax,
		"OTHER":  OtherContact,
	}

	_ContactTypeValueToName = map[ContactType]string{
		Mobile:       "MOBILE",
		Phone:        "PHONE",
		Email:        "EMAIL",
		Fax:          "FAX",
		OtherContact: "OTHER",
	}
)

func init() {
	var v ContactType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ContactTypeNameToValue = map[string]ContactType{
			interface{}(Mobile).(fmt.Stringer).String():       Mobile,
			interface{}(Phone).(fmt.Stringer).String():        Phone,
			interface{}(Email).(fmt.Stringer).String():        Email,
			interface{}(Fax).(fmt.Stringer).String():          Fax,
			interface{}(OtherContact).(fmt.Stringer).String(): OtherContact,
		}
	}
}

// MarshalJSON is generated so ContactType satisfies json.Marshaler.
func (r ContactType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ContactTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ContactType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ContactType satisfies json.Unmarshaler.
func (r *ContactType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ContactType should be a string, got %s", data)
	}
	v, ok := _ContactTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ContactType %q", s)
	}
	*r = v
	return nil
}
