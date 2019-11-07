package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var Version string = "unkown"

type cliFlags struct {
	version *bool // value of -version
}

type cliArgs struct {
	clearPassword *string // value of arg 1
}

func main() {
	cli_flags := getCliFlags()
	cli_args := getCliArgs()

	if *cli_flags.version { // check -version flag
		printVersionAndExit()
	}

	var clearPassword []byte
	if cli_args.clearPassword != nil {
		clearPassword = []byte(*cli_args.clearPassword)
	} else {
		reader := bufio.NewReader(os.Stdin)
		rawInputText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		rawInputText = strings.TrimSpace(rawInputText)

		clearPassword = []byte(rawInputText)
	}

	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword(clearPassword, cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	fmt.Printf("%v", string(hashedPassword))
}

func printVersionAndExit() {
	fmt.Println(Version)
	os.Exit(0)
}

func getCliFlags() cliFlags {
	cli_flags := cliFlags{}

	cli_flags.version = flag.Bool("version", false, "Show version.")
	flag.Parse()

	return cli_flags
}

func getCliArgs() cliArgs {
	cli_args := cliArgs{}

	if len(os.Args) > 1 { // is there a password arg
		cli_args.clearPassword = &os.Args[1]
	} else {
		cli_args.clearPassword = nil
	}

	return cli_args
}
