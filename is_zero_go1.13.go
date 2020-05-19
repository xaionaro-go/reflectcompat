// +build go1.13

package reflectcompat

import (
	"reflect"
)

// IsZero reports whether v is the zero value for its type.
// It panics if the argument is invalid.
func IsZero(v reflect.Value) bool {
	return v.IsZero()
}
