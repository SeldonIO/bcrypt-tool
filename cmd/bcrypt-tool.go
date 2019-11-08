package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Version stores the program version info
var Version string = "unkown"

type commandLineInfo struct {
	versionFlag      *bool   // ptr to value of -version
	clearPasswordArg *string // ptr to value of arg 1 or nil if no arg
}

func main() {
	clInfo := getCommandLineInfo()

	if *clInfo.versionFlag { // check -version flag
		printVersionAndExit()
	}

	var clearPassword []byte
	if clInfo.clearPasswordArg != nil {
		clearPassword = []byte(*clInfo.clearPasswordArg)
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

func getCommandLineInfo() commandLineInfo {
	clInfo := commandLineInfo{}

	{ // setup cli flags
		clInfo.versionFlag = flag.Bool("version", false, "Show version.")
		flag.Parse()
	}

	{ // setup cli args
		if len(os.Args) > 1 { // is there a password arg
			clInfo.clearPasswordArg = &os.Args[1]
		} else {
			clInfo.clearPasswordArg = nil
		}
	}

	return clInfo
}
