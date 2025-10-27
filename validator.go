package validateachangelog

import "fmt"

type Options struct {
	AllowEmptyVersion       bool
	AllowMissingReleaseDate bool
	AllowInvalidChangeType  bool
}

func Validate(c *Changelog, opts *Options) error {
	if opts == nil {
		opts = &Options{}
	}

	if c == nil {
		return fmt.Errorf("nil Changelog")
	}

	if len(c.Versions) == 0 {
		return fmt.Errorf("no versions found in Changelog")
	}

	standardChangeTypes := getStandardChangeTypes()

	for _, version := range c.Versions {
		if version.ReleaseDate.IsZero() && !opts.AllowMissingReleaseDate {
			return fmt.Errorf("missing release date in Changelog entry %s", version.Version)
		}

		if len(version.Entries) == 0 && !opts.AllowEmptyVersion {
			return fmt.Errorf("no entries found in Changelog entry %s", version.Version)
		}

		if !opts.AllowInvalidChangeType {
			for changeType := range version.Entries {
				if _, exists := standardChangeTypes[changeType]; !exists {
					return fmt.Errorf("invalid change type (%s) in Changelog entry %s", changeType, version.Version)
				}
			}
		}
	}

	return nil
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
