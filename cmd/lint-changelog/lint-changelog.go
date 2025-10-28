package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vold-lu/validate-a-changelog/internal"
	"github.com/vold-lu/validate-a-changelog/linter"
)

func main() {
	// Flags
	jsonOutput := flag.Bool("json", false, "output validation issues as json")

	flag.Parse()

	args := flag.Args()

	// Args
	if len(args) < 1 {
		fmt.Println("Usage: lint-changelog [-json] <file>")
		os.Exit(1)
	}

	c, err := linter.LintFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	standardChangeTypes := internal.GetStandardChangeTypes()

	// Reverse map the standard change types by their weight
	sortedStandardChangeTypes := make([]string, len(standardChangeTypes))
	for k, v := range standardChangeTypes {
		sortedStandardChangeTypes[v] = k
	}

	if *jsonOutput {
		if err := json.NewEncoder(os.Stdout).Encode(c); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Handle title (if any)
		if c.Title != "" {
			fmt.Printf("# %s\n\n", c.Title)
		}

		for _, v := range c.Versions {
			var sb strings.Builder

			// Handle version line
			sb.WriteString("## [")
			sb.WriteString(v.Version)
			sb.WriteString("]")

			if v.ReleaseDate != nil {
				sb.WriteString(" ")
				sb.WriteString(v.ReleaseDate.Format("2006-01-02"))
			}

			sb.WriteString("\n\n")

			// Handle section (sorted)
			for _, changeType := range sortedStandardChangeTypes {
				if entries, exists := v.Entries.Get(changeType); exists {
					// Handle section line
					sb.WriteString("### ")
					sb.WriteString(changeType)
					sb.WriteString("\n\n")

					// Handle entries
					for _, entry := range entries {
						sb.WriteString("- ")
						sb.WriteString(entry.Description)
						sb.WriteString("\n")
					}
					sb.WriteString("\n")
				}
			}

			fmt.Print(sb.String())
		}
	}
}
