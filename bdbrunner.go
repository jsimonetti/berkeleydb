package main

import (
  bdb "./berkeleydb"
)

func main() {
  print(bdb.Version())
  print("\n")

  db, err := bdb.CreateDB()
  if err > 0 {
    print("Found error.\n")
    return
  }

  print("Opening database.\n");
  err = db.Open("test.db")
  if err > 0 {
    print("Failed to open database.")
  }

  print("Closing database.\n");
  err = db.Close()
  if err > 0 {
    print("failed to close database.")
  }
}
