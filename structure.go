package validateachangelog

import "time"

type Changelog struct {
	Versions []Version `json:"versions"`
}

type Version struct {
	Version     string    `json:"version"`
	ReleaseDate time.Time `json:"release_date"`

	Entries map[string][]Entry `json:"entries"`
}

type Entry struct {
	Description string `json:"description"`
}
