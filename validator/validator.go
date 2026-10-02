package validator

import (
	"cmp"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/vold-lu/validate-a-changelog"
	"github.com/vold-lu/validate-a-changelog/internal"
	"golang.org/x/mod/semver"
)

const unreleasedVersion = "Unreleased"

var (
	buildChunkRegex = regexp.MustCompile(`\d+|\D+`)
	semverRegex     = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
)

type Options struct {
	AllowEmptyVersion           bool
	AllowMissingReleaseDate     bool
	AllowInvalidChangeType      bool
	AllowInvalidChangeTypeOrder bool
}

func Validate(c *validateachangelog.Changelog, opts *Options) error {
	if opts == nil {
		opts = &Options{}
	}

	err := &ValidationError{}

	if c == nil {
		err.pushIssue("", "", "nil changelog")

		return err
	}

	if len(c.Versions) == 0 {
		err.pushIssue("", "", "no versions found in the changelog")

		return err
	}

	standardChangeTypes := internal.GetStandardChangeTypes()
	standardChangeTypeNames := make([]string, len(standardChangeTypes))

	i := 0
	for changeType := range standardChangeTypes {
		standardChangeTypeNames[i] = changeType
		i++
	}

	previousVersion := ""

	for _, version := range c.Versions {
		// Make sure version is valid
		if version.Version != unreleasedVersion && !semverRegex.MatchString(version.Version) {
			err.pushIssue(version.Version, "", "invalid version")
		}

		// Make sure release have a date
		if version.ReleaseDate == nil && !opts.AllowMissingReleaseDate && version.Version != unreleasedVersion {
			err.pushIssue(version.Version, "", "missing release date in changelog entry")
		}

		// Make sure release contains entries
		if version.Entries.Len() == 0 && !opts.AllowEmptyVersion && version.Version != unreleasedVersion {
			err.pushIssue(version.Version, "", "no sections found in changelog entry")
		}

		// Make sure entries have valid change type
		if !opts.AllowInvalidChangeType {
			for _, changeType := range version.Entries.Keys() {
				if _, exists := standardChangeTypes[changeType]; !exists {
					err.pushIssue(version.Version, changeType, fmt.Sprintf("invalid section `%s` in changelog entry (available values: %v)", changeType, standardChangeTypeNames))
				}
			}
		}

		// Make sure version are in good order
		if previousVersion != "" {
			currentVersion := version.Version

			if previousVersion == unreleasedVersion {
				previousVersion = "99.99.99"
			}
			if version.Version == unreleasedVersion {
				currentVersion = "99.99.99"
			}

			if compareVersions(previousVersion, currentVersion) < 1 {
				err.pushIssue(version.Version, "", "version is not in the right order")
			}
		}

		// Validate that the change type is in the good order
		if !opts.AllowInvalidChangeTypeOrder {
			previousChangeType := ""

			for _, changeType := range version.Entries.Keys() {
				if previousChangeType != "" {

					var previousChangeTypeWeight int
					if val, ok := standardChangeTypes[previousChangeType]; ok {
						previousChangeTypeWeight = val
					} else {
						previousChangeTypeWeight = 999
					}

					var currentChangeTypeWeight int
					if val, ok := standardChangeTypes[changeType]; ok {
						currentChangeTypeWeight = val
					} else {
						currentChangeTypeWeight = 999
					}

					if previousChangeTypeWeight > currentChangeTypeWeight {
						err.pushIssue(version.Version, changeType, fmt.Sprintf("unsorted change type in changelog entry (%s > %s)", changeType, previousChangeType))
					}
				}

				previousChangeType = changeType
			}
		}

		previousVersion = version.Version
	}

	if err.hasIssues() {
		return err
	} else {
		return nil
	}
}

// compareVersions compares two versions like semver.Compare does, except that
// it falls back to the build metadata when both versions have the same semver
// precedence: semver ignores the build metadata, but we rely on it to carry our
// own release counter (1.2.0+5d42a37-vold2 comes after 1.2.0+5d42a37-vold1).
func compareVersions(a, b string) int {
	if c := semver.Compare("v"+a, "v"+b); c != 0 {
		return c
	}

	return compareBuild(semver.Build("v"+a), semver.Build("v"+b))
}

// compareBuild compares two build metadata, digit runs being compared as
// numbers so that `vold10` comes after `vold2`.
func compareBuild(a, b string) int {
	aChunks := buildChunkRegex.FindAllString(a, -1)
	bChunks := buildChunkRegex.FindAllString(b, -1)

	for i := 0; i < len(aChunks) && i < len(bChunks); i++ {
		aNumber, aErr := strconv.Atoi(aChunks[i])
		bNumber, bErr := strconv.Atoi(bChunks[i])

		if aErr == nil && bErr == nil {
			if aNumber != bNumber {
				return cmp.Compare(aNumber, bNumber)
			}

			continue
		}

		if aChunks[i] != bChunks[i] {
			return strings.Compare(aChunks[i], bChunks[i])
		}
	}

	return cmp.Compare(len(aChunks), len(bChunks))
}
