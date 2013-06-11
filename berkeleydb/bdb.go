package berkeleydb

/*
#cgo LDFLAGS: -ldb
#include <db.h> 
#include "bdb.h"
*/
import "C"

import (
  "fmt"
)

const BDB_VERSION string = "1.0.0"

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

func (handle *BDB) Close() int {
  ret := C.go_db_close(handle.db, 0)

  return int(ret)
}


// UTILITY FUNCTIONS

func Version() string {
  lib_version := C.GoString(C.db_full_version(nil, nil, nil, nil, nil))

  tpl := "%s (Go bindings v%s)"
  return fmt.Sprintf(tpl, lib_version, BDB_VERSION)
}
