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
	_SchoolTypeNameToValue = map[string]SchoolType{
		"PRIMARY_SCHOOL":   PrimarySchool,
		"SECONDARY_SCHOOL": SecondarySchool,
		"HIGH_SCHOOL":      HighSchool,
		"COLLEGE":          College,
		"OTHER":            OtherSchool,
	}

	_SchoolTypeValueToName = map[SchoolType]string{
		PrimarySchool:   "PRIMARY_SCHOOL",
		SecondarySchool: "SECONDARY_SCHOOL",
		HighSchool:      "HIGH_SCHOOL",
		College:         "COLLEGE",
		OtherSchool:     "OTHER",
	}
)

func init() {
	var v SchoolType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SchoolTypeNameToValue = map[string]SchoolType{
			interface{}(PrimarySchool).(fmt.Stringer).String():   PrimarySchool,
			interface{}(SecondarySchool).(fmt.Stringer).String(): SecondarySchool,
			interface{}(HighSchool).(fmt.Stringer).String():      HighSchool,
			interface{}(College).(fmt.Stringer).String():         College,
			interface{}(OtherSchool).(fmt.Stringer).String():     OtherSchool,
		}
	}
}

// MarshalJSON is generated so SchoolType satisfies json.Marshaler.
func (r SchoolType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SchoolTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid SchoolType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so SchoolType satisfies json.Unmarshaler.
func (r *SchoolType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("SchoolType should be a string, got %s", data)
	}
	v, ok := _SchoolTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid SchoolType %q", s)
	}
	*r = v
	return nil
}
