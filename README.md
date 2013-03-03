btree
=====

**btree** is a disk-based [B-tree](http://en.wikipedia.org/wiki/B-tree) implementation written in Golang. 

During runtime, all the key nodes are kept in memory and all the values are kept on disk. 
Consequently, the data is split into two files: a key file and a value file.

A tree can be in *read-only* mode or *read-write* mode. 

### [Initialization](https://github.com/3fps/btree/wiki/Initialization)

A tree needs to be open before being used. If the files do exist, they will be created with the given order. If the files are not in proper format, an error will return. Otherwise, the tree header and all the keys will be loaded into memory.

    // btree.READONLY | btree.READWRITE
    tree = &btree.Tree{Order: 128}
    tree, err := btree.Open("/data/keys.b", "/data/vals.b", btree.READONLY)

**Remember to close the tree**
    
    err := tree.Close()

### [Insertion](https://github.com/3fps/btree/wiki/Insertion)

The tree's insert operation expects a string-typed key and binary value. If the key already exists,
the old value will be replaced with the new one. The `Insert` function returns the byte-offset location of the new value.

    key := "name"
    val := []byte("John Appleseed")
    valLoc, err := tree.Insert(key, val)

### [Find](https://github.com/3fps/btree/wiki/Find)
To find a value in the tree, a case-sensitive key is required. If the key does not exist (or has been soft-deleted), it will return `nil`:

    var val []byte = tree.Find("name")
    if val != nil {
      fmt.Println(string(val))
    }

### [Deletion](https://github.com/3fps/btree/wiki/Deletion)
Because key deletion relies on tree restructuring, this implementation only supports soft-delete. Soft-delete marks a key as deleted, but the key remains in the tree structure.

    err := tree.SoftDelete("name")
