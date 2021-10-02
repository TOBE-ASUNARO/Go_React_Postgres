package entity

import (
  "time"
)

type User struct {
  Id uint
  NameMyouzi string `gorm:"size:1"`
  NameMei string `gorm:"size:1"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

