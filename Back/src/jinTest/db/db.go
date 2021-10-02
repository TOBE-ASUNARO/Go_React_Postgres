package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var (
  db *gorm.DB
  err error
)

// init
func Init() {
  db, err = gorm.Open(postgres.Open("db"), &gorm.Config{})

// DB接続
//   db, err = gorm.Open("postgres", "db")

  if err != nil {
    panic(err)
  }

}


// DB取得
func GetDB() *gorm.DB {
  return db
}

// DB接続終了
// func Close() {
//   if err := db.Close(); err != nil {
//     panic(err)
//   }
// }
//
