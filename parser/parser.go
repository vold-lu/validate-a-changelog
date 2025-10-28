package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/vold-lu/validate-a-changelog"
)

var (
	versionRegex           = regexp.MustCompile(`^## \[([0-9.]+)\] - ([0-9]{4}-[0-9]{2}-[0-9]{2})$`)
	unreleasedVersionRegex = regexp.MustCompile(`^## \[Unreleased\]$`)
)

func Parse(r io.Reader) (*validateachangelog.Changelog, error) {
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

			// Parse the new version and register it
			version, releaseDate, err := parseVersionLine(line)
			if version == "" {
				return nil, fmt.Errorf("invalid version line: %s", line)
			}

			if err != nil {
				return nil, err
			}

			currentVersion.Version = version
			currentVersion.ReleaseDate = releaseDate
		}

		// Parse section (Added, Changed, Removed, Fixed)
		if strings.HasPrefix(line, "### ") {
			currentSection = strings.TrimPrefix(line, "### ")

			if _, exists := currentVersion.Entries[currentSection]; exists {
				currentVersion.Entries[currentSection] = []validateachangelog.Entry{}
			}
		}

		// Parse entry
		if strings.HasPrefix(line, "- ") {
			entry := strings.TrimPrefix(line, "- ")

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

func ParseFile(filename string) (*validateachangelog.Changelog, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	return Parse(f)
}

func parseVersionLine(line string) (string, *time.Time, error) {
	// Handle unreleased
	if unreleasedVersionRegex.MatchString(line) {
		return "Unreleased", nil, nil
	}

	parts := versionRegex.FindStringSubmatch(line)

	version := ""
	var releaseDate *time.Time

	// Parse the version
	if len(parts) > 1 {
		version = parts[1]
	}

	// Parse the release date
	if len(parts) > 2 {
		t, err := time.Parse("2006-01-02", parts[2])
		if err == nil {
			releaseDate = &t
		}
	}

	return version, releaseDate, nil
}
