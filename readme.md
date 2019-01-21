# Golang BTree

Binary Search Tree implemented with concurrency safe on Golang.

- Concurrency safe
- Serialization
- Persistent storage



## Structs

#### Binary tree

```go
type BTree struct {
	root *Node
	lock sync.RWMutex
	size int
}
```

#### Node

```go
type Node struct {
	Key           int
	Value         interface{}
	Left          *Node
	Right         *Node
}
```



## Usage

```go
package main

import (
	"strconv"
    "github.com/xfrr/btree"
)

func main() {
    // Create tree
	var tree = &BTree{}
    
    ...
}
```

#### Put

Insert key-value.

``` go
for i := 1; i < 1000; i++ {
	v := strconv.Itoa(i)
	tree.Put(i, v)
}
```

#### Find

Search and return the node searching by key.

```go
node := tree.Find(132)
```

#### Remove

Delete and return the node searching by key.

```go
node := tree.Remove(132)
```

#### Commit

Serialize tree and save it to disk.

```go
err := tree.Commit()
```

#### Load

Load tree from disk.

```go
err := tree.Load()
```

#### Max

Returns the node with max key value. 

```go
max := tree.Max()
```

#### Min

Returns the node with min key value. 

```go
min := tree.Min()
```

#### TraverseInOrder

Iterate over all nodes in order

```go
tree.TraverseInOrder(tree.root, func(n *Node) {
    ... print node
})
```





## Unit Test

- Put
- Find
- Remove
- Max
- Min
- Size
- ...



Command to execute

```shell
go test -v
```

