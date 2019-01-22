package btree

import (
	"errors"
	"sync"
)

// BTree structure
type BTree struct {
	root *Node
	lock sync.RWMutex
	size int
}

// Put Inserts a new node on the tree
func (t *BTree) Put(key int, value string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	newNode := &Node{
		Key:   key,
		Value: value,
		Left:  nil,
		Right: nil,
	}

	// Increment tree size
	t.size++

	// If root node is empty, this new node becomes the root.
	if t.root == nil {
		t.root = newNode
	} else {
		t.root.put(newNode)
	}
}

// Find returns node searching by Key
func (t *BTree) Find(key int) (interface{}, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.root == nil {
		return "", false
	}

	return t.root.find(key)
}

// Max returns the node with max value
func (t *BTree) Max() *Node {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.root == nil {
		return nil
	}

	return t.root.max()
}

// Min returns the node with min value
func (t *BTree) Min() *Node {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t == nil {
		return nil
	}

	return t.root.min()
}

// LeftHeight ...
func (t *BTree) LeftHeight() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	i := 0
	n := t.root
	for n != nil {
		n = n.Left
		i++
	}
	return i
}

// RightHeight ...
func (t *BTree) RightHeight() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	i := 0
	n := t.root
	for n != nil {
		n = n.Right
		i++
	}
	return i
}

// Remove removes a node searching by Key
func (t *BTree) Remove(key int) *Node {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.root == nil {
		return nil
	}

	fakeParent := &Node{Right: t.root}
	n := t.root.remove(key, fakeParent)

	if fakeParent.Right == nil {
		t.root = nil
	}

	return n
}

// TraverseInOrder goes through all the nodes in order
func (t *BTree) TraverseInOrder(n *Node, f func(*Node)) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if n == nil {
		return
	}

	t.TraverseInOrder(n.Left, f)
	f(n)
	t.TraverseInOrder(n.Right, f)
}

// Commit saves the tree to the file storage
func (t *BTree) Commit() error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t == nil || t.root == nil {
		return errors.New("Can not save an empty tree")
	}

	var err error
	var done = make(chan bool)
	go func(err error) {
		err = Serialize(t.root)
		done <- true
	}(err)

	if err != nil {
		return err
	}
	<-done

	t.root = nil

	return nil
}

// Load loads in memory the stored tree
func (t *BTree) Load() error {
	t.lock.Lock()
	defer t.lock.Unlock()

	var err error
	var done = make(chan bool)
	go func(err error) {
		err = Deserialize(t.root)
		done <- true
	}(err)

	if err != nil {
		return err
	}
	<-done

	return nil
}

// Size prints the tree
func (t *BTree) Size() int {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.size
}
