package berkeleydb

/*
#cgo LDFLAGS: -ldb
#include <db.h>
#include "bdb.h"
*/
import "C"

type Environment struct {
	environ *C.DB_ENV
}

func NewEnvironment() (*Environment, error) {
	var env *C.DB_ENV
	err := C.db_env_create(&env, 0)
	if err > 0 {
		return nil, createError(err)
	}

	return &Environment{env}, nil
}

func (env *Environment) Open(path string, flags C.u_int32_t, fileMode int) error {
	mode := C.u_int32_t(fileMode)

	err := C.go_env_open(env.environ, C.CString(path), flags, mode)
	return createError(err)
}

func (env *Environment) Close() error {
	err := C.go_env_close(env.environ, 0)
	return createError(err)
}
