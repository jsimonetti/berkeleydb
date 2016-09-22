### BerkeleyDB Bindings

This package provides BerkeleyDB wrappers for the C library using `cgo`.

To build, you will need a relatively recent version of BerkeleyDB.
package main



### Example
```go

package main

import (
        "fmt"

        "github.com/jsimonetti/berkeleydb"
)

func main() {
        var err error

        db, err := berkeleydb.NewDB()
        if err != nil {
                fmt.Printf("Unexpected failure of CreateDB %s\n", err)
        }

        err = db.Open("./test.db", berkeleydb.DB_HASH, berkeleydb.DB_CREATE)
        if err != nil {
                fmt.Printf("Could not open test_db.db. Error code %s", err)
                return
        }
        defer db.Close()

        err = db.PutString("key", "value")
        if err != nil {
                fmt.Printf("Expected clean PutString: %s\n", err)
        }

        value, err := db.GetString("key")
        if err != nil {
                fmt.Printf("Unexpected error in GetString: %s\n", err)
                return
        }
        fmt.Printf("value: %s\n", value)

}

```
