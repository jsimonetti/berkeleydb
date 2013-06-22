package berkeleydb

/*
#cgo LDFLAGS: -ldb
#include <db.h>
#include "bdb.h"
*/
import "C"

import (
	"fmt"
	"errors"
)

const BDB_VERSION string = "0.0.1"

// Flags for opening a database or environment.
const (
	DB_CREATE   = C.DB_CREATE
	DB_EXCL     = C.DB_EXCL
	DB_RDONLY   = C.DB_RDONLY
	DB_TRUNCATE = C.DB_TRUNCATE

  // Env
  DB_INIT_MPOOL = C.DB_INIT_MPOOL
)

// Database types.
const (
	DB_BTREE   = C.DB_BTREE
	DB_HASH    = C.DB_HASH
	DB_HEAP    = C.DB_HEAP
	DB_RECNO   = C.DB_RECNO
	DB_QUEUE   = C.DB_QUEUE
	DB_UNKNOWN = C.DB_UNKNOWN
)

type BDB struct {
	db *C.DB
}

type Errno int

func NewDB() (*BDB, error) {
	var db *C.DB
	err := C.db_create(&db, nil, 0)

	if err > 0 {
		return nil, createError(err)
	}

	return &BDB{db}, nil
}

func NewDBInEnvironment(env *Environment) (*BDB, error) {
	var db *C.DB
	err := C.db_create(&db, env.environ, 0)

	if err > 0 {
		return nil, createError(err)
	}

	return &BDB{db}, nil
}

func (handle *BDB) OpenWithTxn(filename string, txn *C.DB_TXN, dbtype C.DBTYPE, flags C.u_int32_t) error {
	db := handle.db

	ret := C.go_db_open(db, txn, C.CString(filename), nil, dbtype, flags, 0)

	return createError(ret)
}

func (handle *BDB) Open(filename string, dbtype C.DBTYPE, flags C.u_int32_t) error {
	ret := C.go_db_open(handle.db, nil, C.CString(filename), nil, dbtype, flags, 0)

	return createError(ret)
}

func (handle *BDB) Close() error {
	ret := C.go_db_close(handle.db, 0)

	return createError(ret)
}

func (handle *BDB) OpenFlags() (C.u_int32_t, error) {
	var flags C.u_int32_t

	ret := C.go_db_get_open_flags(handle.db, &flags)

	return flags, createError(ret)
}

func (handle *BDB) Remove(filename string) error {
	ret := C.go_db_remove(handle.db, C.CString(filename))

	return createError(ret)
}

func (handle *BDB) Rename(oldname, newname string) error {
	ret := C.go_db_rename(handle.db, C.CString(oldname), C.CString(newname))

	return createError(ret)
}

// Convenience function to store a string.
func (handle *BDB) PutString(name, value string) error {
	ret := C.go_db_put_string(handle.db, C.CString(name), C.CString(value), 0)
	if ret > 0 {
		return createError(ret)
	}
	return nil
}

// Convenience function to get a string.
func (handle *BDB) GetString(name string) (string, error) {
	value := C.CString("")
	ret := C.go_db_get_string(handle.db, C.CString(name), value)
	return C.GoString(value), createError(ret)
}

func (handle *BDB) DeleteString(name string) error {
	ret := C.go_db_del_string(handle.db, C.CString(name));
	return createError(ret);
}

// UTILITY FUNCTIONS

func Version() string {
	lib_version := C.GoString(C.db_full_version(nil, nil, nil, nil, nil))

	tpl := "%s (Go bindings v%s)"
	return fmt.Sprintf(tpl, lib_version, BDB_VERSION)
}

type DBError struct {
	Code int
	Message string
}

func createError(code C.int) error {
	if code == 0 {
		return nil
	}
	msg := C.db_strerror(code)
	e := DBError{int(code), C.GoString(msg)}
	return errors.New(e.Error())
}

func (e *DBError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

