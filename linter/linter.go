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
		Entries:     map[string][]validateachangelog.Entry{},
	}
	currentSection := ""

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse version
		if strings.HasPrefix(line, "## ") {
			// Push current version if there is one
			if currentVersion.Version != "" {
				c.Versions = append(c.Versions, currentVersion)
				currentVersion = validateachangelog.Version{
					Version:     "",
					ReleaseDate: &time.Time{},
					Entries:     map[string][]validateachangelog.Entry{},
				}
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
					version = "[Unreleased]"
				}
			}

			// Validate that we at least have a version
			if version == "" {
				return nil, fmt.Errorf("invalid version line: %s", line)
			}

			currentVersion.Version = version
			currentVersion.ReleaseDate = releaseDate
		}

		// Parse section (Added, Changed, Removed, Fixed)
		if internal.IsSectionLine(line) {
			currentSection = internal.ParseSectionLine(line)

			if currentVersion.Version == "" {
				return nil, fmt.Errorf("invalid changelog section: %s (no version found)", line)
			}

			if _, exists := currentVersion.Entries[currentSection]; exists {
				currentVersion.Entries[currentSection] = []validateachangelog.Entry{}
			}
		}

		// Parse entry
		if internal.IsEntryLine(line) {
			entry := internal.ParseEntryLine(line)

			if currentSection == "" {
				return nil, fmt.Errorf("invalid changelog entry: %s (no section found)", line)
			}

			currentVersion.Entries[currentSection] = append(currentVersion.Entries[currentSection], validateachangelog.Entry{
				Description: entry,
			})
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
