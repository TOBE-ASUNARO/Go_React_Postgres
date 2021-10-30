package entity

import (
  "time"
)

type User struct {
  Id uint
  NameSei string `gorm:"size:1"`
  NameMei string `gorm:"size:1"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

