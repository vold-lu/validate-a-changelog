package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vold-lu/validate-a-changelog/parser"
	"github.com/vold-lu/validate-a-changelog/validator"
)

func main() {
	// Flags
	allowEmptyVersion := flag.Bool("allow-empty-version", false, "allow version without entries")
	allowMissingReleaseDate := flag.Bool("allow-missing-release-date", false, "allow version without release date")
	allowInvalidChangeType := flag.Bool("allow-invalid-change-type", false, "allow section with invalid change type")

	flag.Parse()

	args := flag.Args()

	// Args
	if len(args) < 1 {
		fmt.Println("Usage: validate-changelog [-allow-empty-version] [-allow-invalid-change-type] [-allow-missing-release-date] <file>")
		os.Exit(1)
	}

	c, err := parser.ParseFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	opts := &validator.Options{
		AllowEmptyVersion:       *allowEmptyVersion,
		AllowMissingReleaseDate: *allowMissingReleaseDate,
		AllowInvalidChangeType:  *allowInvalidChangeType,
	}

	if err := validator.Validate(c, opts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
