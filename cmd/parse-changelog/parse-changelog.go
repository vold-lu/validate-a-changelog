package main

import (
	"encoding/json"
	"fmt"
	"os"

	validateachangelog "github.com/vold-lu/validate-a-changelog"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: parse-changelog <file>")
		os.Exit(1)
	}

	changelogFile := os.Args[1]

	c, err := validateachangelog.ParseFile(changelogFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.NewEncoder(os.Stdout).Encode(c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
