package validator

import (
	"testing"
	"time"

	"github.com/vold-lu/validate-a-changelog"
)

func TestValidateEmptyChangelog(t *testing.T) {
	c := &validateachangelog.Changelog{}

	if err := Validate(c, nil); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: false,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateUnreleasedVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: false,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogNonMissingReleaseDateAndNotAllowed(t *testing.T) {
	releaseDate := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC)

	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: &releaseDate,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: false,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogEmptyVersionAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogEmptyVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       false,
		AllowInvalidChangeType:  true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogEmptyUnreleasedVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       false,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogNonEmptyVersionAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]validateachangelog.Entry{
					"Added": {
						{Description: "Test description"},
					},
				},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       false,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeAndAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]validateachangelog.Entry{
					"InvalidType": {
						{Description: "Test description"},
					},
				},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidChangeTypeAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]validateachangelog.Entry{
					"InvalidType": {
						{Description: "Test description"},
					},
				},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  false,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogValidChangeTypeAndNotAllowed(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]validateachangelog.Entry{
					"Added": {
						{Description: "Test description"},
					},
				},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  false,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogInvalidVersion(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "Test",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogGoodVersionOrder(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
			{
				Version:     "0.15.10",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogGoodVersionOrderWithUnreleased(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err != nil {
		t.Fail()
	}
}

func TestValidateChangelogBadVersionOrder(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "0.15.10",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogBadVersionOrderWithUnreleased(t *testing.T) {
	c := &validateachangelog.Changelog{
		Versions: []validateachangelog.Version{
			{
				Version:     "0.15.10",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
			{
				Version:     "Unreleased",
				ReleaseDate: nil,
				Entries:     map[string][]validateachangelog.Entry{},
			},
		},
	}

	if err := Validate(c, &Options{
		AllowMissingReleaseDate: true,
		AllowEmptyVersion:       true,
		AllowInvalidChangeType:  true,
	}); err == nil {
		t.Fail()
	}
}
