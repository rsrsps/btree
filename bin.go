package btree

// putUint16 puts an 16 bit unsigned number into the first 2 bytes of the slice
func putUint16(b []byte, n uint16) {
	b[0] = byte((n & 0xFF00) >> 8)
	b[1] = byte((n & 0x00FF))
}

// putInt64 puts a 16 bit signed number into the first 8 bytes of the slice
func putInt64(b []byte, n int64) {
	for i := 7; i >= 0; i-- {
		b[i] = byte(n & 0xFF)
		n >>= 8
	}
}
