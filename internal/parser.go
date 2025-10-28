package internal

import (
	"regexp"
	"time"
)

var (
	titleRegex             = regexp.MustCompile(`^# (.*)$`)
	versionRegex           = regexp.MustCompile(`^## \[([0-9.]+)\] ?-? ?([0-9]{4}-[0-9]{2}-[0-9]{2})?$`)
	unreleasedVersionRegex = regexp.MustCompile(`^## \[Unreleased\]$`)
	sectionRegex           = regexp.MustCompile(`^### (.*)$`)
	entryRegex             = regexp.MustCompile(`^- (.*)$`)
)

func IsTitleLine(line string) bool {
	return titleRegex.MatchString(line)
}

func ParseTitleLine(line string) string {
	matches := titleRegex.FindStringSubmatch(line)
	if len(matches) == 0 {
		return ""
	}

	return matches[1]
}

func IsVersionLine(line string) bool {
	return versionRegex.MatchString(line) || unreleasedVersionRegex.MatchString(line)
}

func ParseVersionLine(line string) (string, *time.Time, error) {
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

	// Parse the release date (if any)
	if len(parts) > 2 && parts[2] != "" {
		t, err := time.Parse("2006-01-02", parts[2])
		if err != nil {
			return "", nil, err
		}

		releaseDate = &t
	}

	return version, releaseDate, nil
}

func IsSectionLine(line string) bool {
	return sectionRegex.MatchString(line)
}

func ParseSectionLine(line string) string {
	matches := sectionRegex.FindStringSubmatch(line)
	if len(matches) == 0 {
		return ""
	}

	return matches[1]
}

func IsEntryLine(line string) bool {
	return entryRegex.MatchString(line)
}

func ParseEntryLine(line string) string {
	matches := entryRegex.FindStringSubmatch(line)
	if len(matches) == 0 {
		return ""
	}

	return matches[1]
}

func GetStandardChangeTypes() map[string]int {
	return map[string]int{
		"Added":      0,
		"Changed":    1,
		"Deprecated": 2,
		"Removed":    3,
		"Fixed":      4,
		"Security":   5,
	}
}
