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

const BDB_VERSION string = "0.0.1"

// Flags for opening a database.
const (
  DB_CREATE = C.DB_CREATE
  DB_EXCL = C.DB_EXCL
  DB_RDONLY = C.DB_RDONLY
  DB_TRUNCATE= C.DB_TRUNCATE
)

// Database types.
const (
  DB_BTREE = C.DB_BTREE
  DB_HASH = C.DB_HASH
  DB_HEAP = C.DB_HEAP
  DB_RECNO = C.DB_RECNO
  DB_QUEUE = C.DB_QUEUE
  DB_UNKNOWN = C.DB_UNKNOWN
)

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

func (handle *BDB) Open(filename string, dbtype C.DBTYPE, flags C.u_int32_t) int {
  db := handle.db

  ret := C.go_db_open(db, nil, C.CString(filename), nil, dbtype, flags, 0)

  return int(ret)
}

func (handle *BDB) Close() int {
  ret := C.go_db_close(handle.db, 0)

  return int(ret)
}

func (handle *BDB) OpenFlags() (C.u_int32_t, int) {
  var flags C.u_int32_t

  ret := C.go_db_get_open_flags(handle.db, &flags)

  return flags, int(ret)
}

func (handle *BDB) Remove(filename string) int {
  ret := C.go_db_remove(handle.db, C.CString(filename))

  return int(ret)
}


// UTILITY FUNCTIONS

func Version() string {
  lib_version := C.GoString(C.db_full_version(nil, nil, nil, nil, nil))

  tpl := "%s (Go bindings v%s)"
  return fmt.Sprintf(tpl, lib_version, BDB_VERSION)
}
