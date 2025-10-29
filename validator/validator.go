package validator

import (
	"fmt"
	"regexp"

	"github.com/vold-lu/validate-a-changelog"
	"github.com/vold-lu/validate-a-changelog/internal"
	"golang.org/x/mod/semver"
)

const unreleasedVersion = "Unreleased"

var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

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

			if semver.Compare("v"+previousVersion, "v"+currentVersion) < 1 {
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
