package btree

const VERSION uint8 = 1 // current version

const (
	READONLY  int = 1 // read-only mode
	READWRITE int = 2 // read-write mode
)

const HEADER_SIZE int64 = 512 // file header size
