# PerkDB

A persistent key-value database with support for arbitrary indexes. It
provides a REST API and stores documents as JSON data.

**NOTICE: Due to the change in license for BerkeleyDB, I have discontinued development on this package.**

## Sub-Packages

* `perkdb/berkeleydb`: Go bindings for the BerkeleyDB C library.

### BerkeleyDB Bindings

This package provides BerkeleyDB wrappers for the C library using `cgo`.

To build, you will need a relatively recent version of BerkeleyDB.
package main



### Example
```go

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

	print("Opening database.\n")
	err = db.Open("test.db", bdb.DB_BTREE, bdb.DB_CREATE)
	if err > 0 {
		print("Failed to open database.")
	}

	print("Closing database.\n")
	err = db.Close()
	if err > 0 {
		print("failed to close database.")
	}
}

```
