package berkeleydb

import(
	"testing"
	"os"
)

const TEST_DIR = "./TEST_ENV"

func TestNewEnvironment(t *testing.T) {
	_, err := os.Stat(TEST_DIR)
	if err != nil && os.IsNotExist(err) {
		e := os.Mkdir(TEST_DIR, os.ModeDir | os.ModePerm)
		if e != nil {
			t.Fatal("Failed to create directory: %s", e)
		}
	}

	_, err = NewEnvironment()

	if err != nil {
		t.Error("Expected environment, got %s", err)
	}
	
}

func TestOpenEnvironment(t *testing.T) {
	env, _ := NewEnvironment()
	err := env.Open(TEST_DIR, DB_CREATE | DB_INIT_MPOOL, 0)
	if err != nil {
		t.Error("Expected to open DB, got %s", err)
	}

	err = env.Close()
	if err != nil {
		t.Error("Expected to close DB, got %s", err)
	}
}

func TestOpenDBInEnvironment(t *testing.T) {
	env, _ := NewEnvironment()
	err := env.Open(TEST_DIR, DB_CREATE | DB_INIT_MPOOL, 0755)
	if err != nil {
		t.Error("Expected to open DB, got ", err)
		return
	}

	// Now create, open, and close a DB
	db, err := NewDBInEnvironment(env)
	if err != nil {
		t.Error("Expected to create new DB: ", err)
	}

	err = db.Open(TEST_FILENAME, DB_BTREE, DB_CREATE)
	if err != nil {
		t.Error("Expected to open DB, got ", err)
	}

	// Test that the DB file was actually created.
	_, err = os.Stat(TEST_DIR + "/" + TEST_FILENAME)
	if err != nil  {
		t.Error("Expected to stat .db, got ", err)
	}

	err = db.Close()
	if err != nil {
		t.Error("Expected to close the DB, got ", err)
	}

	err = env.Close()
	if err != nil {
		t.Error("Expected to close DB, got %s", err)
	}
}
func TestTeardown(t *testing.T) {
	err := os.RemoveAll(TEST_DIR)
	if err != nil {
		t.Fatal("Expected to remove fixtures, got %s", err)
	}
}
