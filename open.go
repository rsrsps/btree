package btree

import "os"

// Open prepares the tree with the specify mode.
// If the mode is READWRITE and the files do not exist, they will be created.
// If the mode is READONLY and either file does not exist or not in the proper format, it'll return an error.
func (this *Tree) Open(keyPath, valPath string, mode int) error {
	if mode == READWRITE {
		if !fileExists(keyPath) && !fileExists(valPath) {
			return this.createTree(keyPath, valPath)
		}
	}

	return nil

}

// fileExists returns true iff the file in the given path exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// createTree creates two files and writes default header information to each file
// and returns a tree structure
func (this *Tree) createTree(keyPath, valPath string) (err error) {
	this.fkey, err = os.Create(keyPath)
	if err != nil {
		return
	}

	this.fval, err = os.Create(valPath)
	if err != nil {
		return
	}

	this.version = VERSION
	err = this.writeFullKeyHeader()
	if err != nil {
		return
	}
	err = this.writeFullValHeader()
	return
}

// writeFullKeyHeader writes the whole key header to file
func (this *Tree) writeFullKeyHeader() error {
	buff := make([]byte, HEADER_SIZE, HEADER_SIZE)
	buff[0], buff[2] = 0x12, this.version
	putUint16(buff[3:5], this.Order)
	if this.root != nil {
		putInt64(buff[5:13], this.root.loc)
	}
	_, err := this.fkey.WriteAt(buff, 0)
	return err
}

// writeFullValHeader writes the whol value header to file
func (this *Tree) writeFullValHeader() error {
	buff := make([]byte, HEADER_SIZE, HEADER_SIZE)
	buff[0], buff[1], buff[2] = 0x12, 0x10, this.version
	_, err := this.fval.WriteAt(buff, 0)
	return err
}
