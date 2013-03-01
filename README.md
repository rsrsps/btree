btree
=====

**btree** is a disk based [B-tree](http://en.wikipedia.org/wiki/B-tree) implementation in Golang. 

At runtime, all the key nodes are kept in memory and all the values are kept on disk. 
Consequently, the data is split into two files: a key file and a value file.

A tree can be in *read-only* mode or *read-write* mode. 

### [Initialization](https://github.com/3fps/btree/wiki/Initialization)

Creating a new tree will return a tree in *read-write* mode. If the data files exist, they will be truncated

    order := 10 // order is the maximum number of child-nodes that each node can have
    tree, err := btree.New("/data/keys.b", "/data/vals.b", order)
    
    
Opening an existing tree

    // btree.READONLY | btree.READWRITE
    tree, err := btree.Open("/data/keys.b", "/data/vals.b", btree.READONLY)

**Do not forget to close the tree**
    
    err := tree.Close()

### [Insertion](https://github.com/3fps/btree/wiki/Insertion)

The tree's insert operation expects a string-typed key and binary value. If the key already exists,
the old value will be replaced with the new one. The `Insert` function returns the byte-offset location of the new value.

    key := "name"
    val := []byte("John Appleseed")
    valLoc, err := tree.Insert(key, val)

### [Find](https://github.com/3fps/btree/wiki/Find)
To look for a value in the tree, a case-sensitive key is required. If the key does not exist (or soft deleted), it'll return `nil`:

    var val []byte = tree.Find("name")
    if val != nil {
      fmt.Println(string(val))
    }

### [Deletion](https://github.com/3fps/btree/wiki/Deletion)
Because key deletion tend to require a lot of change in the tree structure, this implementation only support soft-delete. Soft-delete marks a key as deleted, but it still remains in the tree structure.

    err := tree.SoftDelete("name")
