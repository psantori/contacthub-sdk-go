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
