package main

import (
	"encoding/json"
	"fmt"
	"os"

	validateachangelog "github.com/vold-lu/validate-a-changelog/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: parse-changelog <file> [version]")
		os.Exit(1)
	}

	changelogFile := os.Args[1]
	version := ""

	if len(os.Args) > 2 {
		version = os.Args[2]
	}

	c, err := validateachangelog.ParseFile(changelogFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output the whole changelog
	if version == "" {
		if err := json.NewEncoder(os.Stdout).Encode(c); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Output specific version
		for _, entry := range c.Versions {
			if entry.Version == version {
				if err := json.NewEncoder(os.Stdout).Encode(entry); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				return
			}
		}
	}
}
