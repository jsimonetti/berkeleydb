package berkeleydb

import "testing"

const TEST_FILENAME = "test_db.db"

func TestCreate(t *testing.T) {

	_, err := CreateDB()

	if err != nil {
		t.Errorf("Expected error code 0, got %d", err)
	}
}
func TestOpen(t *testing.T) {
	db, err := CreateDB()

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

func TestRemove(t *testing.T) {
	db, _ := CreateDB()

	err := db.Remove(TEST_FILENAME)
	if err != nil {
		t.Errorf("Could not delete %s. Expected 0, got %s", TEST_FILENAME, err)
	}
}

func TestRename(t *testing.T) {
	db, _ := CreateDB()
	db.Open(TEST_FILENAME, DB_HASH, DB_CREATE)
	db.Close()

	db, _ = CreateDB()

	newname := "foo_" + TEST_FILENAME
	err := db.Rename(TEST_FILENAME, newname)
	if err != nil {
		t.Errorf("Could not rename %s to %s", TEST_FILENAME, newname)
	}

	db, _ = CreateDB()
	err = db.Remove(newname)
	if err != nil {
		t.Errorf("Could not remove %s", newname)
	}
}


