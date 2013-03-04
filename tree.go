package btree

import "os"

// Tree represents a b-tree structure
type Tree struct {
	Order   uint16   // tree order M
	version uint8    // file format version
	fkey    *os.File // file handle to the key file
	fval    *os.File // file handle to the value file
	root    *Node    // root node
}
