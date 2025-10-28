package main

import (
	"encoding/json"
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
	allowInvalidChangeTypeOrder := flag.Bool("allow-invalid-change-type-order", false, "allow section with invalid change type ordering")
	jsonOutput := flag.Bool("json", false, "output validation issues as json")

	flag.Parse()

	args := flag.Args()

	// Args
	if len(args) < 1 {
		fmt.Println("Usage: validate-changelog [-allow-empty-version] [-allow-missing-release-date] [-allow-invalid-change-type] [-allow-invalid-change-type-order] [-json] <file>")
		os.Exit(1)
	}

	c, err := parser.ParseFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	opts := &validator.Options{
		AllowEmptyVersion:           *allowEmptyVersion,
		AllowMissingReleaseDate:     *allowMissingReleaseDate,
		AllowInvalidChangeType:      *allowInvalidChangeType,
		AllowInvalidChangeTypeOrder: *allowInvalidChangeTypeOrder,
	}

	if err := validator.Validate(c, opts); err != nil {
		if *jsonOutput {
			if err := json.NewEncoder(os.Stdout).Encode(err.(*validator.ValidationError)); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
