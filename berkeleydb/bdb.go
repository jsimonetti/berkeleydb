package berkeleydb

/*
#cgo LDFLAGS: -ldb
#include <db.h> 
#include "bdb.h"
*/
import "C"

type BDB struct {
  db *C.DB
}

type Errno int

func CreateDB() (*BDB, int) {
  var db *C.DB
  err := C.db_create(&db, nil, 0)

  if err > 0 {
    return nil, int(err)
  }

  return &BDB{db}, 0
}

func (handle *BDB) Open(filename string) int {
  db := handle.db

  ret := C.go_db_open(db, nil, C.CString(filename), nil, C.DB_BTREE, C.DB_CREATE, 0)

  return int(ret)
}
