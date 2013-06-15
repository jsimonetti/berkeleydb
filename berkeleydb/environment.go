package berkeleydb

/*
#cgo LDFLAGS: -ldb
#include <db.h>
#include "bdb.h"
*/
import "C"

type Environment struct {
	env *C.DB_ENV
}

func NewEnvironment() (*Environment, error) {
	var env *C.DB_ENV
	err := C.db_env_create(&env, 0)
	if err > 0 {
		return nil, createError(err)
	}

	return &Environment{env}, nil
}
