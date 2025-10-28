package validator

import (
	"regexp"

	"github.com/vold-lu/validate-a-changelog"
)

const unreleasedVersion = "Unreleased"

var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

type Options struct {
	AllowEmptyVersion       bool
	AllowMissingReleaseDate bool
	AllowInvalidChangeType  bool
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

	standardChangeTypes := getStandardChangeTypes()

	for _, version := range c.Versions {
		// Make sure version is valid
		if version.Version != unreleasedVersion && !semverRegex.MatchString(version.Version) {
			err.pushIssue(version.Version, "", "invalid version")
		}

		if version.ReleaseDate == nil && !opts.AllowMissingReleaseDate && version.Version != unreleasedVersion {
			err.pushIssue(version.Version, "", "missing release date in changelog entry")
		}

		if len(version.Entries) == 0 && !opts.AllowEmptyVersion && version.Version != unreleasedVersion {
			err.pushIssue(version.Version, "", "no sections found in changelog entry")
		}

		if !opts.AllowInvalidChangeType {
			for changeType := range version.Entries {
				if _, exists := standardChangeTypes[changeType]; !exists {
					err.pushIssue(version.Version, changeType, "invalid change type in changelog entry")
				}
			}
		}
	}

	if err.hasIssues() {
		return err
	} else {
		return nil
	}
}

func getStandardChangeTypes() map[string]interface{} {
	return map[string]interface{}{
		"Added":      nil,
		"Changed":    nil,
		"Deprecated": nil,
		"Removed":    nil,
		"Fixed":      nil,
		"Security":   nil,
	}
}
