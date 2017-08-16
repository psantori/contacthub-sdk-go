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

package client

import (
	"fmt"
	"strings"
	"time"
)

const (
	simpleDateFormat  = "2006-01-02"
	defaultDateFormat = "2006-01-02T15:04:05.999-0700"
)

// SimpleDate is a date formatted as simpleDateFormat
type SimpleDate struct {
	time.Time
}

// UnmarshalJSON implements the Unmarshaler interface
func (sd *SimpleDate) UnmarshalJSON(b []byte) (err error) {
	sd.Time, err = unmarshalTime(b, simpleDateFormat)
	return
}

// MarshalJSON implements the Marshaler interface
func (sd *SimpleDate) MarshalJSON() ([]byte, error) {
	return marshalTime(sd.Time, simpleDateFormat)
}

// CustomDate is a date formatted as defaultDateFormat
type CustomDate struct {
	time.Time
}

// UnmarshalJSON implements the Unmarshaler interface
func (sd *CustomDate) UnmarshalJSON(b []byte) (err error) {
	sd.Time, err = unmarshalTime(b, defaultDateFormat)
	return
}

// MarshalJSON implements the Marshaler interface
func (sd *CustomDate) MarshalJSON() ([]byte, error) {
	return marshalTime(sd.Time, defaultDateFormat)
}

func marshalTime(datetime time.Time, format string) ([]byte, error) {
	if datetime.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", datetime.Format(format))), nil
}

func unmarshalTime(b []byte, format string) (time.Time, error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return time.Time{}, nil
	}
	return time.Parse(format, s)
}
