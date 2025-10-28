package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/vold-lu/validate-a-changelog"
	"github.com/vold-lu/validate-a-changelog/internal"
)

func Parse(r io.Reader) (*validateachangelog.Changelog, error) {
	c := &validateachangelog.Changelog{}

	currentVersion := &validateachangelog.Version{
		Version:     "",
		ReleaseDate: &time.Time{},
		Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
	}
	currentSection := ""

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse title
		if internal.IsTitleLine(line) {
			c.Title = internal.ParseTitleLine(line)
		}

		// Parse version
		if internal.IsVersionLine(line) {
			// Push current version if there is one
			if currentVersion.Version != "" {
				c.Versions = append(c.Versions, currentVersion)
				currentVersion = &validateachangelog.Version{
					Version:     "",
					ReleaseDate: &time.Time{},
					Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
				}
				currentSection = ""
			}

			// Parse the new version and register it
			version, releaseDate, err := internal.ParseVersionLine(line)
			if err != nil {
				return nil, err
			}

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

			if !currentVersion.Entries.Has(currentSection) {
				_ = currentVersion.Entries.Set(currentSection, []validateachangelog.Entry{})
			}
		}

		// Parse entry
		if internal.IsEntryLine(line) {
			entry := internal.ParseEntryLine(line)

			if currentSection == "" {
				return nil, fmt.Errorf("invalid changelog entry: %s (no section found)", line)
			}

			// Todo: optimise?
			currentVersionEntries, _ := currentVersion.Entries.Get(currentSection)
			currentVersionEntries = append(currentVersionEntries, validateachangelog.Entry{
				Description: entry,
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
