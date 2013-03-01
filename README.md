btree
=====

**btree** is a disk based [B-tree](http://en.wikipedia.org/wiki/B-tree) implementation in Golang. 

At runtime, all the key nodes are kept in memory and all the values are kept on disk. 
Consequently, the data is stored in two files: key file and value file.

A tree can be in *read-only* mode or *read-write* mode. 

### Initialization

Creating a new tree will return a tree in *read-write* mode. If the data files exist, they will be truncated

    order := 10 // order is the maximum number of child-nodes that each node can have
    tree, err := btree.New("/data/keys.b", "/data/vals.b", order)
    
Opening an existing tree

    // btree.READONLY | btree.READWRITE
    tree, err := btree.Open("/data/keys.b", "/data/vals.b", btree.READONLY)

### Insertion

The tree's insert operation expects a string-typed key and binary value. If the key already exists,
the old value will be replaced with the new one

    key := "name"
    val := []byte("John Appleseed")
    err := tree.Insert(key, val)

In the background, the following actions are performed:

1. Append the new value to the value file (even if the key exists, old values are not removed from disk).
This step gives us byte-offset V of the new value.
2. Transverse the tree to find the key:
    1. If the key exists, replace its value pointer with V from step 1.
    2. If the key does not exist, append a new key to the key file. This gives us the byte-offset K.
