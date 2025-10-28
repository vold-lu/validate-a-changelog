package validator

import (
	"github.com/vold-lu/validate-a-changelog"
)

const unreleasedVersion = "Unreleased"

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
		if version.ReleaseDate == nil && !opts.AllowMissingReleaseDate && version.Version != unreleasedVersion {
			err.pushIssue(version.Version, "", "missing release date in changelog entry")
		}

		if len(version.Entries) == 0 && !opts.AllowEmptyVersion {
			err.pushIssue(version.Version, "", "no sections found in changelog entry")
		}

		if !opts.AllowInvalidChangeType {
			for changeType := range version.Entries {
				if _, exists := standardChangeTypes[changeType]; !exists {
					err.pushIssue(version.Version, changeType, "no change type in changelog entry")
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
