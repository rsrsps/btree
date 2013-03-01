btree
=====

**btree** is a disk based [B-tree](http://en.wikipedia.org/wiki/B-tree) implementation in Golang. 

At runtime, all the key nodes are kept in memory and all the values are kept on disk. 
Consequently, the data is stored in two files: key file and value file.

A tree can be in *read-only* mode or *read-write* mode. 

Creating a new tree will return a tree in *read-write* mode. If the data files exist, they will be truncated:

    order := 10 // order is the maximum number of child-nodes that each node can have
    tree, err := btree.New("/data/keys.b", "/data/vals.b", order)
    
Opening an existing tree:

    // btree.READONLY | btree.READWRITE
    tree, err := btree.Open("/data/keys.b", "/data/vals.b", btree.READONLY)
