package btree

import (
	"strconv"
	"testing"
)

var tree *Tree

func TestPut(t *testing.T) {
	tree = &Tree{}

	for i := 0; i < 1000; i++ {
		v := strconv.Itoa(i)
		tree.Put(i, v)
	}

	//tree.String()
}

func TestFind(t *testing.T) {
	_, r1found := tree.Find(2)
	_, r2found := tree.Find(102)
	_, r3found := tree.Find(632)
	_, r4found := tree.Find(999)

	if !r1found || !r2found || !r3found || !r4found {
		t.Errorf("Key not found, find not working")
	}
}
