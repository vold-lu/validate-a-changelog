package validateachangelog

import (
	"time"

	"github.com/vold-lu/validate-a-changelog/internal"
)

type Changelog struct {
	Title    string     `json:"title"`
	Versions []*Version `json:"versions"`
}

type Version struct {
	Version     string     `json:"version"`
	ReleaseDate *time.Time `json:"release_date"`

	Entries internal.SortedMap[string, []Entry] `json:"entries"`
}

type Entry struct {
	Description string `json:"description"`
}
