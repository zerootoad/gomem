package helpers

import (
	"fmt"
)

// PtrToHex converts uintptr to hex string.
func PtrToHex(ptr uintptr) string {
	return fmt.Sprintf("%#x", ptr)
}
