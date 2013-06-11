package main

import (
  bdb "./berkeleydb"
)

func main() {
  print("test\n")

  db, err := bdb.CreateDB()
  if err > 0 {
    print("Found error.\n")
    return
  }

  db.Open("test.db")
}
