package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestIsTitleLine(t *testing.T) {
	cases := []struct {
		Line    string
		IsValid bool
	}{
		{
			Line:    "## [Unreleased]",
			IsValid: false,
		},
		{
			Line:    "## [0.1.0]",
			IsValid: false,
		},
		{
			Line:    "test",
			IsValid: false,
		},
		{
			Line:    "# 0.1.0",
			IsValid: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("IsTitleLine(%s)", c.Line), func(t *testing.T) {
			if ok := IsTitleLine(c.Line); ok != c.IsValid {
				t.Logf("IsTitleLine(%s). Got %v, wanted %v", c.Line, ok, c.IsValid)
				t.Fail()
			}
		})
	}
}

func TestParseTitleLine(t *testing.T) {
	cases := []struct {
		Line  string
		Title string
	}{
		{
			Line: "## [Unreleased]",
		},
		{
			Line: "## [0.1.0]",
		},
		{
			Line:  "# 0.1.0",
			Title: "0.1.0",
		},
		{
			Line: "Test",
		},
		{
			Line: "## 0.1.0 - 2025-10-28",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ParseTitleLine(%s)", c.Line), func(t *testing.T) {
			if titleLine := ParseTitleLine(c.Line); titleLine != c.Title {
				t.Logf("ParseTitleLine(%s). Expected %s got: %s", c.Line, c.Title, titleLine)
				t.Fail()
			}
		})
	}
}

func TestIsVersionLine(t *testing.T) {
	cases := []struct {
		Line    string
		IsValid bool
	}{
		{
			Line:    "## [Unreleased]",
			IsValid: true,
		},
		{
			Line:    "## [0.1.0]",
			IsValid: true,
		},
		{
			Line:    "## 0.1.0",
			IsValid: false,
		},
		{
			Line:    "# 0.1.0",
			IsValid: false,
		},
		{
			Line:    "Test",
			IsValid: false,
		},
		{
			Line:    "## 0.1.0 - 2025-10-28",
			IsValid: false,
		},
		{
			Line:    "## 0.1.0 - 100-10-10",
			IsValid: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("IsVersionLine(%s)", c.Line), func(t *testing.T) {
			if ok := IsVersionLine(c.Line); ok != c.IsValid {
				t.Logf("IsVersionLine(%s). Got %v, wanted %v", c.Line, ok, c.IsValid)
				t.Fail()
			}
		})
	}
}

func TestParseVersionLine(t *testing.T) {
	date := time.Date(2025, 10, 28, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		Line        string
		IsValid     bool
		Version     string
		ReleaseDate *time.Time
	}{
		{
			Line:    "## [Unreleased]",
			IsValid: true,
			Version: "Unreleased",
		},
		{
			Line:    "## [0.1.0]",
			IsValid: true,
			Version: "0.1.0",
		},
		{
			Line:    "## 0.1.0",
			IsValid: false,
		},
		{
			Line:    "# 0.1.0",
			IsValid: false,
		},
		{
			Line:    "Test",
			IsValid: false,
		},
		{
			Line:    "## 0.1.0 - 2025-10-28",
			IsValid: false,
		},
		{
			Line:        "## [0.1.0] - 2025-10-28",
			IsValid:     true,
			Version:     "0.1.0",
			ReleaseDate: &date,
		},
		{
			Line:    "## 0.1.0 - 100-10-10",
			IsValid: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ParseVersionLine(%s)", c.Line), func(t *testing.T) {
			version, releaseDate, err := ParseVersionLine(c.Line)
			if err != nil && c.IsValid {
				t.Logf("ParseVersionLine(%s). Expected no errors but: %v", c.Line, err)
				t.Fail()

				return
			}

			if version != c.Version {
				t.Logf("ParseVersionLine(%s). Got %v, wanted %v", c.Line, version, c.Version)
				t.Fail()
			}

			if (releaseDate == nil && c.ReleaseDate != nil) || (releaseDate != nil && c.ReleaseDate == nil) {
				t.Logf("ParseVersionLine(%s). Got %v, wanted %v", c.Line, releaseDate, c.ReleaseDate)
				t.Fail()

				return
			}

			if releaseDate != nil && c.ReleaseDate != nil {
				if *releaseDate != *c.ReleaseDate {
					t.Logf("ParseVersionLine(%s). Got %v, wanted %v", c.Line, releaseDate, c.ReleaseDate)
					t.Fail()
				}
			}
		})
	}
}

func TestIsSectionLine(t *testing.T) {
	cases := []struct {
		Line    string
		IsValid bool
	}{
		{
			Line:    "### Added",
			IsValid: true,
		},
		{
			Line:    "### Changed",
			IsValid: true,
		},
		{
			Line:    "## 0.1.0",
			IsValid: false,
		},
		{
			Line:    "# Added",
			IsValid: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("IsSectionLine(%s)", c.Line), func(t *testing.T) {
			if ok := IsSectionLine(c.Line); ok != c.IsValid {
				t.Logf("IsSectionLine(%s). Got %v, wanted %v", c.Line, ok, c.IsValid)
				t.Fail()
			}
		})
	}
}

func TestParseSectionLine(t *testing.T) {
	cases := []struct {
		Line    string
		Section string
	}{
		{
			Line: "## [Unreleased]",
		},
		{
			Line: "## [0.1.0]",
		},
		{
			Line:    "### Added",
			Section: "Added",
		},
		{
			Line:    "### Other",
			Section: "Other",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ParseSectionLine(%s)", c.Line), func(t *testing.T) {
			section := ParseSectionLine(c.Line)
			if section != c.Section {
				t.Logf("ParseSectionLine(%s). Got %s, wanted %s", c.Line, section, c.Section)
				t.Fail()
			}
		})
	}
}

func TestIsEntryLine(t *testing.T) {
	cases := []struct {
		Line    string
		IsValid bool
	}{
		{
			Line:    "### Changed",
			IsValid: false,
		},
		{
			Line:    "- Something",
			IsValid: true,
		},
		{
			Line:    "Added",
			IsValid: false,
		},
		{
			Line:    " - Test",
			IsValid: true,
		},
		{
			Line:    "	- Test",
			IsValid: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("IsEntryLine(%s)", c.Line), func(t *testing.T) {
			if ok := IsEntryLine(c.Line); ok != c.IsValid {
				t.Logf("IsEntryLine(%s). Got %v, wanted %v", c.Line, ok, c.IsValid)
				t.Fail()
			}
		})
	}
}

func TestParseEntryLine(t *testing.T) {
	cases := []struct {
		Line  string
		Entry string
	}{
		{
			Line: "## [0.1.0]",
		},
		{
			Line: "### Added",
		},
		{
			Line:  "- Other",
			Entry: "Other",
		},
		{
			Line: "Other",
		},
		{
			Line:  " - Test",
			Entry: "Test",
		},
		{
			Line:  "	- Test",
			Entry: "Test",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ParseEntryLine(%s)", c.Line), func(t *testing.T) {
			entry := ParseEntryLine(c.Line)
			if entry != c.Entry {
				t.Logf("ParseEntryLine(%s). Got %s, wanted %s", c.Line, entry, c.Entry)
				t.Fail()
			}
		})
	}
}
