package berkeleydb

import "testing"

const TEST_FILENAME = "test_db.db"

func TestCreate(t *testing.T) {

	_, err := CreateDB()

	if err > 0 {
		t.Errorf("Expected error code 0, got %d", err)
	}
}
func TestOpen(t *testing.T) {
	db, err := CreateDB()

	if err > 0 {
		t.Errorf("Unexpected failure of CreateDB")
	}

	err = db.Open(TEST_FILENAME, DB_BTREE, DB_CREATE)

	if err > 0 {
		t.Errorf("Could not open test_db.db. Error code %d", err)
	}

	err = db.Close()
	if err > 0 {
		t.Errorf("Could not close file %s", TEST_FILENAME)
		return
	}

	/*
	  err = db.Remove(TEST_FILENAME)
	  if err > 0 {
	    t.Errorf("Could not remove file %s", TEST_FILENAME)
	  }
	*/

}
