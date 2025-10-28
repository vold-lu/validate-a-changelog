package linter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	validateachangelog "github.com/vold-lu/validate-a-changelog"
	"github.com/vold-lu/validate-a-changelog/internal"
)

var (
	versionRegex           = regexp.MustCompile(`^## \[?([0-9.]+)\]? ?-? ?([0-9]{4}-[0-9]{2}-[0-9]{2})?$`)
	unreleasedVersionRegex = regexp.MustCompile(`^## \[?Unreleased\]?$`)
)

func Lint(r io.Reader) (*validateachangelog.Changelog, error) {
	c := &validateachangelog.Changelog{}

	currentVersion := validateachangelog.Version{
		Version:     "",
		ReleaseDate: &time.Time{},
		Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
	}
	currentSection := ""

	standardChangeTypes := internal.GetStandardChangeTypes()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse title
		if internal.IsTitleLine(line) {
			c.Title = internal.ParseTitleLine(line)
		}

		// Parse version
		if strings.HasPrefix(line, "## ") {
			// Push current version if there is one
			if currentVersion.Version != "" {
				c.Versions = append(c.Versions, currentVersion)
				currentVersion = validateachangelog.Version{
					Version:     "",
					ReleaseDate: &time.Time{},
					Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
				}
				currentSection = ""
			}

			var version string
			var releaseDate *time.Time

			// Determinate whether it is a valid line
			if internal.IsVersionLine(line) {
				version, releaseDate, _ = internal.ParseVersionLine(line)

				currentVersion.Version = version
				currentVersion.ReleaseDate = releaseDate
			} else {
				// Try to manually recover the line
				if versionRegex.MatchString(line) {
					parts := versionRegex.FindStringSubmatch(line)
					if len(parts) > 1 {
						version = parts[1]
					}
					if len(parts) > 2 && parts[2] != "" {
						t, err := time.Parse("2006-01-02", parts[2])
						if err == nil {
							releaseDate = &t
						}
					}

				} else if unreleasedVersionRegex.MatchString(line) || len(c.Versions) == 0 {
					version = "Unreleased"
				}
			}

			// Validate that we at least have a version
			if version == "" {
				return nil, fmt.Errorf("invalid version line: %s", line)
			}

			currentVersion.Version = version
			currentVersion.ReleaseDate = releaseDate
		} else if internal.IsSectionLine(line) {
			// Parse section (Added, Changed, Removed, Fixed)

			currentSection = internal.ParseSectionLine(line)

			// Hardcoded mapping
			// TODO: How can we improve/extended this?
			if _, exists := standardChangeTypes[currentSection]; !exists {
				switch strings.ToLower(currentSection) {
				case "fix":
					currentSection = "Fixed"
					break
				case "change":
					currentSection = "Changed"
					break
				case "new":
					currentSection = "Added"
					break
				}
			}

			if currentVersion.Version == "" {
				return nil, fmt.Errorf("invalid changelog section: %s (no version found)", line)
			}

			if !currentVersion.Entries.Has(currentSection) {
				_ = currentVersion.Entries.Set(currentSection, []validateachangelog.Entry{})
			}
		} else if internal.IsEntryLine(line) {
			// Parse entry
			entry := internal.ParseEntryLine(line)

			if !strings.HasSuffix(entry, ".") {
				entry = fmt.Sprintf("%s.", entry)
			}

			if currentSection == "" {
				return nil, fmt.Errorf("invalid changelog entry: %s (no section found)", line)
			}

			// Todo: optimise?
			currentVersionEntries, _ := currentVersion.Entries.Get(currentSection)
			currentVersionEntries = append(currentVersionEntries, validateachangelog.Entry{
				Description: entry,
			})
			_ = currentVersion.Entries.Set(currentSection, currentVersionEntries)
		} else if strings.Trim(line, " ") != "" {
			if !strings.HasSuffix(line, ".") {
				line = fmt.Sprintf("%s.", line)
			}

			if currentSection == "" {
				currentSection = "Added"
			}

			if !currentVersion.Entries.Has(currentSection) {
				_ = currentVersion.Entries.Set(currentSection, []validateachangelog.Entry{})
			}

			// Todo: optimise?
			currentVersionEntries, _ := currentVersion.Entries.Get(currentSection)
			currentVersionEntries = append(currentVersionEntries, validateachangelog.Entry{
				Description: line,
			})
			_ = currentVersion.Entries.Set(currentSection, currentVersionEntries)
		}
	}

	// Push the latest version (if any)
	if currentSection != "" {
		c.Versions = append(c.Versions, currentVersion)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(c.Versions) == 0 {
		return nil, fmt.Errorf("no versions found in changelog")
	}

	return c, nil
}

func LintFile(filename string) (*validateachangelog.Changelog, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	return Lint(f)
}
