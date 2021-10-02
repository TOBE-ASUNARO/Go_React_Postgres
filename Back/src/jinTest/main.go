package main

import (
  "jinTest/server"
  "jinTest/db"
)

func main() {
  db.Init()
  server.Init()
}
