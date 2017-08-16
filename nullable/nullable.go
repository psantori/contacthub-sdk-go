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

package nullable

import "github.com/guregu/null"

// String is a simple alias to null.String pointer
type String *null.String

// StringFrom is a simple wrapper to get a pointer from guregu/null.StringForm
func StringFrom(s string) String {
	nullString := null.StringFrom(s)
	return &nullString
}

// StringFromPtr is a simple wrapper to get a pointer from guregu/null.StringFromPtr
func StringFromPtr(s *string) String {
	nullString := null.StringFromPtr(s)
	return &nullString
}

// NullString returns an invalid null.String, which will Marshal to a null JSON field
func NullString() String {
	return StringFromPtr(nil)
}

// Int is a simple alias to null.Int pointer
type Int *null.Int

// IntFrom is a simple wrapper to get a pointer from guregu/null.IntForm
func IntFrom(s int64) Int {
	nullInt := null.IntFrom(s)
	return &nullInt
}

// IntFromPtr is a simple wrapper to get a pointer from guregu/null.IntFromPtr
func IntFromPtr(s *int64) Int {
	nullInt := null.IntFromPtr(s)
	return &nullInt
}

// NullInt returns an invalid null.Int, which will Marshal to a null JSON field
func NullInt() Int {
	return IntFromPtr(nil)
}

// Float is a simple alias to null.Float pointer
type Float *null.Float

// FloatFrom is a simple wrapper to get a pointer from guregu/null.FloatForm
func FloatFrom(s float64) Float {
	nullFloat := null.FloatFrom(s)
	return &nullFloat
}

// FloatFromPtr is a simple wrapper to get a pointer from guregu/null.FloatFromPtr
func FloatFromPtr(s *float64) Float {
	nullFloat := null.FloatFromPtr(s)
	return &nullFloat
}

// NullFloat returns an invalid null.Float, which will Marshal to a null JSON field
func NullFloat() Float {
	return FloatFromPtr(nil)
}

// Bool is a simple alias to null.Bool pointer
type Bool *null.Bool

// BoolFrom is a simple wrapper to get a pointer from guregu/null.BoolForm
func BoolFrom(s bool) Bool {
	nullBool := null.BoolFrom(s)
	return &nullBool
}

// BoolFromPtr is a simple wrapper to get a pointer from guregu/null.BoolFromPtr
func BoolFromPtr(s *bool) Bool {
	nullBool := null.BoolFromPtr(s)
	return &nullBool
}

// NullBool returns an invalid null.Bool, which will Marshal to a null JSON field
func NullBool() Bool {
	return BoolFromPtr(nil)
}
