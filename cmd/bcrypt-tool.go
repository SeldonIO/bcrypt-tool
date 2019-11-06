package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var clearPassword = []byte("testpassword")
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword(clearPassword, cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%v", string(hashedPassword))
	}
}
