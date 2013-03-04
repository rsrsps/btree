package btree

import (
	"testing"
	"os/user"
	"path/filepath"
	"os"
)

// TestCreateNwTree tests creating a new non-existing tree
func TestCreateNewTree(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	keyPath := filepath.Join(usr.HomeDir, "btree_TestCreateNewTree_key.b")
	valPath := filepath.Join(usr.HomeDir, "btree_TestCreateNewTree_val.b")
	defer os.Remove(keyPath)
	defer os.Remove(valPath)

	tree := &Tree{Order: 16}
	err = tree.Open(keyPath, valPath, READWRITE)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// test key header content
	fkey, err := os.Open(keyPath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer fkey.Close()

	buff := make([]byte, HEADER_SIZE)

	n, err := fkey.ReadAt(buff, 0)
	if n != int(HEADER_SIZE) {
		t.Errorf("Key header file only has %d bytes instead of %d as expected.", n, HEADER_SIZE)
		t.FailNow()
	}

	if buff[0] != 0x12 && buff[1] != 0x00 {
		t.Error("The first two bytes of key header must be 0x1200")
	}
	if buff[2] != tree.version {
		t.Errorf("Key header version is %d instead of %d as expected.", buff[2], tree.version)
	}
	if buff[4] != byte(tree.Order) {
		t.Errorf("Key header order is %d instead of %d as expected.", buff[4], tree.Order)
	}

	// test value header content
	fval, err := os.Open(valPath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer fval.Close()

	buff = make([]byte, HEADER_SIZE)

	n, err = fval.ReadAt(buff, 0)
	if n != int(HEADER_SIZE) {
		t.Errorf("Key header file only has %d bytes instead of %d as expected.", n, HEADER_SIZE)
		t.FailNow()
	}

	if buff[0] != 0x12 && buff[1] != 0x10 {
		t.Error("The first two bytes of value header must be 0x1210")
	}
	if buff[2] != tree.version {
		t.Errorf("Value header version is %d instead of %d as expected.", buff[2], tree.version)
	}
}