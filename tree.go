package btree

import (
	"fmt"
	"sync"
)

type Tree struct {
	root *Node
	lock sync.RWMutex
}

func (t *Tree) Put(key int, value string) {

	t.lock.Lock()
	defer t.lock.Unlock()

	newNode := &Node{
		Key:   key,
		Value: value,
		Left:  nil,
		Right: nil,
	}

	if t.root == nil {
		t.root = newNode
		return
	}

	putNode(t.root, newNode)
}

func (t *Tree) Find(key int) (string, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return findNode(t.root, key)
}

func (t *Tree) Max() *Node {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t == nil {
		return nil
	}

	node := t.root
	for {
		if node.Right == nil {
			return node
		}
	}
}

func (t *Tree) Min() *Node {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t == nil {
		return nil
	}

	node := t.root
	for {
		if node.Left == nil {
			return node
		}
	}
}

func (t *Tree) Remove(key int, parent *Node) {
	t.lock.Lock()
	defer t.lock.Unlock()
	removeNode(t.root, key)
}

func putNode(n *Node, newNode *Node) {
	switch {
	case newNode.Key < n.Key:
		//fmt.Printf("LEFT: %d\n", n.Key)
		if n.Left == nil {
			n.Left = newNode
		} else {
			putNode(n.Left, newNode)
		}
	default:
		//fmt.Printf("RIGHT: %d\n", n.Key)
		if n.Right == nil {
			n.Right = newNode
		} else {
			putNode(n.Right, newNode)
		}
	}
}

func findNode(n *Node, key int) (string, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case key < n.Key:
		return findNode(n.Left, key)
	case key > n.Key:
		return findNode(n.Right, key)
	}
	return n.Value, true
}

func removeNode(n *Node, key int) *Node {
	if n == nil {
		return nil
	}

	switch {
	case key < n.Key:
		n.Left = removeNode(n.Left, key)
		return n
	case key > n.Key:
		n.Right = removeNode(n.Right, key)
		return n
	}

	if n.Left == nil && n.Right == nil {
		n = nil
		return nil
	}

	if n.Left == nil {
		n = n.Right
		return n
	}
	if n.Right == nil {
		n = n.Left
		return n
	}

	smallestRightValue := n.Right

	for {
		if smallestRightValue != nil && smallestRightValue.Left != nil {
			smallestRightValue = smallestRightValue.Left
		} else {
			break
		}
	}

	n.Key = smallestRightValue.Key
	n.Value = smallestRightValue.Value
	n.Right = removeNode(n.Right, n.Key)
	return n
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Key)
		stringify(n.Right, level)
	}
}
