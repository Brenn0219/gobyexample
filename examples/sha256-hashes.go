package examples

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Hashes() {
	s := "sha256 this string"
	h := sha256.New()

	h.Write([]byte(s))
	println(h)

	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
