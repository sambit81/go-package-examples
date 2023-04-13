package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(str string) string {
	var hashcode = md5.Sum([]byte(str))    // It returns the checksum from the string
	return hex.EncodeToString(hashcode[:]) // and then we use the hex package to find the hexadecimal
	// encoding of this input.
}

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("Password encrypted to:", encodeWithMD5(password))
	// We can use this hexadecimal string as our encrpted password
}
