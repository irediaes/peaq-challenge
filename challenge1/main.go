package main

import (
	"fmt"
	"unsafe"
)

// SOURCE:: https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga
func main() {
	// this could be some public information, e.g. version information
	harmlessData := [8]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'}
	// however, this could be critical private information such as a private key
	secret := [17]byte{'l', '3', '3', 't', '-', 'h', '4', 'x', 'x', '0', 'r', '-', 'w', '1', 'n', 's', '!'}

	// (accidentally) cast harmless buffer into a new buffer type of wrong size
	var dangerousData = (*[8 + 17]byte)(unsafe.Pointer(&harmlessData[0]))

	// print (misused) buffer
	fmt.Println(string((*dangerousData)[:]))
	fmt.Println(secret)
}
