package validator

import (
	"testing"
	"time"

	"github.com/vold-lu/validate-a-changelog"
	"github.com/vold-lu/validate-a-changelog/internal"
)

func TestValidateEmptyChangelog(t *testing.T) {
	c := &validateachangelog.Changelog{}

	if err := Validate(c, nil); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     false,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateUnreleasedVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     false,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogNonMissingReleaseDateAndNotAllowed(t *testing.T) {
	releaseDate := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC)

	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: &releaseDate,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     false,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogEmptyVersionAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogEmptyVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           false,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogEmptyUnreleasedVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           false,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogNonEmptyVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Added"}, map[string][]validateachangelog.Entry{
					"Added": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           false,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"InvalidType"}, map[string][]validateachangelog.Entry{
					"InvalidType": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"InvalidType"}, map[string][]validateachangelog.Entry{
					"InvalidType": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      false,
		AllowInvalidChangeTypeOrder: true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogValidChangeTypeAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Added"}, map[string][]validateachangelog.Entry{
					"Added": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      false,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidVersion(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "Test",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogGoodVersionOrder(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
			{
				Version:     "0.15.10",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogGoodVersionOrderWithUnreleased(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogBadVersionOrder(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "0.15.10",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogBadVersionOrderWithUnreleased(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "0.15.10",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     *internal.NewEmptyMap[string, []validateachangelog.Entry](),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeOrderAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Removed", "Added"}, map[string][]validateachangelog.Entry{
					"Removed": {
						{Description: "Test description"},
					},
					"Added": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeOrderAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Removed", "Added"}, map[string][]validateachangelog.Entry{
					"Removed": {
						{Description: "Test description"},
					},
					"Added": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: false,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogValidChangeTypeOrderAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Added", "Changed", "Removed", "Fixed"}, map[string][]validateachangelog.Entry{
					"Added": {
						{Description: "Test description"},
					},
					"Changed": {
						{Description: "Test description"},
					},
					"Removed": {
						{Description: "Test description"},
					},
					"Fixed": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: false,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeOrderAndNotAllowedWithCustomChangeType(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Waaza", "Removed", "Added"}, map[string][]validateachangelog.Entry{
					"Waaza": {
						{Description: "Test description"},
					},
					"Removed": {
						{Description: "Test description"},
					},
					"Added": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: false,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogValidChangeTypeOrderAndNotAllowedWithCustomChangeType(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []*validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: *internal.NewSortedMap([]string{"Added", "Removed", "Waaza"}, map[string][]validateachangelog.Entry{
					"Added": {
						{Description: "Test description"},
					},
					"Removed": {
						{Description: "Test description"},
					},
					"Waaza": {
						{Description: "Test description"},
					},
				}),
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate:     true,
		AllowEmptyVersion:           true,
		AllowInvalidChangeType:      true,
		AllowInvalidChangeTypeOrder: false,
	}); err != nil {
		t.Fail()
	}
}
