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
	chunkRegex  = regexp.MustCompile(`\d+|\D+`)
	semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
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
// it compares the prerelease and the build metadata as natural strings (see
// compareChunks) instead of following the semver rules: semver compares the
// alphanumeric prerelease identifiers lexically and ignores the build metadata
// altogether, but we rely on both to carry our own release counter
// (1.2.0-vold10 comes after 1.2.0-vold9, same for 1.2.0+5d42a37-vold10).
func compareVersions(a, b string) int {
	va, vb := "v"+a, "v"+b

	// Compare the release part (major.minor.patch) using the semver rules.
	if c := semver.Compare(releasePart(va), releasePart(vb)); c != 0 {
		return c
	}

	aPrerelease, bPrerelease := semver.Prerelease(va), semver.Prerelease(vb)

	// A version with a prerelease has a lower precedence than the release one.
	if (aPrerelease == "") != (bPrerelease == "") {
		if aPrerelease == "" {
			return 1
		}

		return -1
	}

	if c := compareChunks(aPrerelease, bPrerelease); c != 0 {
		return c
	}

	return compareChunks(semver.Build(va), semver.Build(vb))
}

// releasePart returns the version stripped from its prerelease & build
// metadata, i.e. only its major.minor.patch part.
func releasePart(v string) string {
	return strings.TrimSuffix(semver.Canonical(v), semver.Prerelease(v))
}

// compareChunks compares two version chunks (prerelease or build metadata),
// digit runs being compared as numbers so that `vold10` comes after `vold2`.
func compareChunks(a, b string) int {
	aChunks := chunkRegex.FindAllString(a, -1)
	bChunks := chunkRegex.FindAllString(b, -1)

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
