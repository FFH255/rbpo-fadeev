package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"ssd-lab-pswd-go/src/pkg/io"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Enter alg (md5, sha256): ")
	alg := io.ReadLine(r)

	fmt.Print("Enter password: ")
	password := io.ReadLine(r)

	var h hash.Hash

	if alg == "sha256" {
		h = sha256.New()
	} else {
		h = md5.New()
	}

	passwordHash := h.Sum([]byte(password))

	fmt.Printf("password hash: %s\n", hex.EncodeToString(passwordHash))
}
