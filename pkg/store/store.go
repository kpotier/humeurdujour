package store

import "time"

type Mood struct {
	Date time.Time `gorm:"unique"`

	Global       string
	Sleep        string
	Anxiety      string
	DarkThoughts string
	Ruminations  string
	Notes        string
}

type Store interface {
	GetMood(from time.Time, to time.Time) ([]*Mood, error)
	SetMood(m *Mood) error
}
