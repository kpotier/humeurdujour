package gorm

import (
	"time"

	"github.com/kpotier/humeurdujour/pkg/store"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) store.Store {
	db.AutoMigrate(&store.Mood{})
	return &Store{db}
}

func (s *Store) GetMood(from time.Time, to time.Time) ([]*store.Mood, error) {
	var m []*store.Mood
	err := s.DB.Where("date>=? AND date <= ?", from, to).Find(&m).Error
	return m, err
}

func (s *Store) SetMood(m *store.Mood) error {
	var mood store.Mood
	err := s.DB.Where("date=?", m.Date).Take(&mood).Error
	if err != nil {
		return s.DB.Create(m).Error
	}
	return s.DB.Model(&store.Mood{}).Where("date=?", m.Date).Updates(m).Error
}
