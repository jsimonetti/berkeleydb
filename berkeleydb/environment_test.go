package berkeleydb

import(
	"testing"
	"os"
)

const TEST_DIR = "./TEST_ENV"

func TestNewEnvironment(t *testing.T) {
	_, err := os.Stat(TEST_DIR)
	if err != nil && os.IsNotExist(err) {
		e := os.Mkdir(TEST_DIR, os.ModeDir)
		if e != nil {
			t.Fatal("Failed to create directory: %s", e)
		}
	}

	_, err = NewEnvironment()

	if err != nil {
		t.Error("Expected environment, got %s", err)
	}
	
}

func TestTeardown(t *testing.T) {
	err := os.RemoveAll(TEST_DIR)
	if err != nil {
		t.Fatal("Expected to remove fixtures, got %s", err)
	}
}
