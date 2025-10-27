package validate_a_changelog

import "time"

type Changelog struct {
	Versions []Version
}

type Version struct {
	Version     string
	ReleaseDate time.Time

	Entries map[string][]Entry
}

type Entry struct {
	Description string
}
