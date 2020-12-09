package util

import "unsafe"

func toString(v []byte) string {
	return *(*string)(unsafe.Pointer(&v))
}

// StringToReadOnlyBytes returns a byte slice over the read-only data of the
// given string.
//
// Attempting to modify to this byte slice will result in a panic.
func StringToReadOnlyBytes(v *string) []byte {
	return *(*[]byte)(unsafe.Pointer(v))
}