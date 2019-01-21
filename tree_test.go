package btree

import (
	"strconv"
	"sync"
	"testing"
)

var tree *BTree
var ids = [...]int{1, 2, 8, 798, 1560, 10000}

// TestPut ...
func TestIntPut(t *testing.T) {
	tree = &BTree{}
	for i := 1; i <= ids[len(ids)-1]; i++ {
		v := strconv.Itoa(i)
		tree.Put(i, v)
	}

	s := tree.Size()
	if s != ids[len(ids)-1] {
		t.Errorf("Inserted -%d-, put function not working", s)
	}

}

// TestPutConcurrency ...
func TestIntPutConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	stree := &BTree{}

	wg.Add(ids[len(ids)-1])
	for i := 1; i <= ids[len(ids)-1]; i++ {
		v := strconv.Itoa(i)
		go func(i int) {
			defer wg.Done()
			stree.Put(i, v)
		}(i)
	}

	wg.Wait()

	s := tree.Size()
	if s != ids[len(ids)-1] {
		t.Errorf("Inserted -%d-, put function not working", s)
	}
}

// TestFind ...
func TestIntFind(t *testing.T) {
	for i := 0; i < len(ids); i++ {
		_, found := tree.Find(ids[i])
		if !found {
			t.Errorf("Could not find key: %d, Find function not working", ids[i])
		}
	}
}

// TestMax ...
func TestIntMax(t *testing.T) {
	r := tree.Max()
	if r.Key != ids[len(ids)-1] {
		t.Errorf("Invalid Max value")
	}
}

// TestMin ...
func TestIntMin(t *testing.T) {
	r := tree.Min()
	if r.Key != ids[0] {
		t.Errorf("Invalid Min value")
	}
}

// TestRemove using concurrency
func TestIntRemove(t *testing.T) {
	for i := 0; i < len(ids); i++ {
		var n *Node
		n = tree.Remove(ids[i])
		if n == nil {
			t.Errorf("Could not remove key: %d, Remove function not working", ids[i])
		}
	}

	// tree.TraverseInOrder(tree.root, func(n *Node) { fmt.Println(n.Key) })
}

// TestIntSave ...
func TestIntSave(t *testing.T) {
	err := tree.Commit()
	if err != nil {
		t.Error(err)
	}
}

// TestIntLoad ...
func TestIntLoad(t *testing.T) {
	err := tree.Load()
	if err != nil {
		t.Error(err)
	}
}

// TestSize ...
func TestIntSize(t *testing.T) {
	s := tree.Size()
	if s != (ids[len(ids)-1]) {
		t.Errorf("Incorrect size -%d-", s)
	}
}
