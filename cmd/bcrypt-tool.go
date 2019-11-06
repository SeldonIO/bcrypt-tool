package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawInputText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	rawInputText = strings.TrimSpace(rawInputText)

	var clearPassword = []byte(rawInputText)
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword(clearPassword, cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	fmt.Printf("%v", string(hashedPassword))
}
