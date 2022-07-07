package helpers

import (
	"crypto/rand"
	"fmt"
)

func GenToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
