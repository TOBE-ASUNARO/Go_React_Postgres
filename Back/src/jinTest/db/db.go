package db

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var (
  db *gorm.DB
  err error
)

// init
func Init() {


//
  dsn := "host=db user=root password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//


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
