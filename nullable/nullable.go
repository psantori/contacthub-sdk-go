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
