package berkeleydb

import "testing"

const TEST_FILENAME = "test_db.db"

func TestNewDB(t *testing.T) {

	_, err := NewDB()

	if err != nil {
		t.Errorf("Expected error code 0, got %d", err)
	}
}
func TestOpen(t *testing.T) {
	db, err := NewDB()

	if err != nil {
		t.Errorf("Unexpected failure of CreateDB")
	}

	err = db.Open(TEST_FILENAME, DB_BTREE, DB_CREATE)

	if err != nil  {
		t.Errorf("Could not open test_db.db. Error code %s", err)
	}

	flags, err := db.OpenFlags()
	if err != nil {
		t.Errorf("Could not get OpenFlags: %s", err)
	}
	if flags != DB_CREATE {
		t.Errorf("Expected flag to match DB_CREATE, got %d", flags)
	}

	err = db.Close()
	if err != nil {
	  t.Errorf("Could not close file %s: %s", TEST_FILENAME, err)
		return
	}


}

func openDB() (*BDB, error) {
	db, err := NewDB()

	if err != nil {
		return nil, err
	}

	err = db.Open(TEST_FILENAME, DB_BTREE, DB_CREATE)

	if err != nil  {
		return nil, err
	}

	return db, nil;
}

func closeDB(db *BDB) error {
	return db.Close()
}

func TestPutString(t *testing.T) {
	db, err := openDB()
	defer closeDB(db)


	err = db.PutString("key", "value")
	if err != nil {
		t.Error("Expected clean PutString.", err)
	}
}

func TestGetString(t *testing.T) {
	db, err := openDB()
	defer closeDB(db)


	err = db.PutString("key", "value")
	if err != nil {
		t.Error("Expected clean PutString: ", err)
	}

	val, err := db.GetString("key")
	if err != nil {
		t.Error("Unexpected error in GetString: ", err)
		return
	}

	if val != "value" {
		t.Error("Expected 'value', got ", val)
	}
}

func TestDeleteString(t *testing.T) {
	db, err := openDB()
	defer closeDB(db)


	err = db.PutString("key", "value")
	if err != nil {
		t.Error("Expected clean PutString: ", err)
	}

	err = db.DeleteString("key")
	if err != nil {
		t.Error("Expected a clean delete, got ", err)
	}

	err = db.DeleteString("nosuchkey")
	if err != nil {
		t.Error("Expected empty delete, got ", err)
	}
}

func TestRemove(t *testing.T) {
	db, _ := NewDB()

	err := db.Remove(TEST_FILENAME)
	if err != nil {
		t.Errorf("Could not delete %s. Expected 0, got %s", TEST_FILENAME, err)
	}
}

func TestRename(t *testing.T) {
	db, _ := NewDB()
	db.Open(TEST_FILENAME, DB_HASH, DB_CREATE)
	db.Close()

	db, _ = NewDB()

	newname := "foo_" + TEST_FILENAME
	err := db.Rename(TEST_FILENAME, newname)
	if err != nil {
		t.Errorf("Could not rename %s to %s", TEST_FILENAME, newname)
	}

	db, _ = NewDB()
	err = db.Remove(newname)
	if err != nil {
		t.Errorf("Could not remove %s", newname)
	}
}


