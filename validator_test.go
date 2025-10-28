package validateachangelog

import (
	"testing"
	"time"
)

func TestValidateEmptyChangelog(t *testing.T) {
	c := &Changelog{}

	if err := Validate(c, nil); err == nil {
		t.Fail()
	}
}

func TestValidateChangelogMissingReleaseDateAndAllowed(t *testing.T) {
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]Entry{},
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

func TestValidateChangelogMissingReleaseDateAndNotAllowed(t *testing.T) {
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]Entry{},
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

	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: &releaseDate,
				Entries:     map[string][]Entry{},
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
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]Entry{},
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
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries:     map[string][]Entry{},
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

func TestValidateChangelogNonEmptyVersionAndNotAllowed(t *testing.T) {
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]Entry{
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
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]Entry{
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
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]Entry{
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
	c := &Changelog{
		Versions: []Version{
			{
				Version:     "1.0.0",
				ReleaseDate: nil,
				Entries: map[string][]Entry{
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
