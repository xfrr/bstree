package btree

// Node ...
type Node struct {
	Key   int
	Value interface{}
	Left  *Node
	Right *Node
}

func (n *Node) put(newNode *Node) {
	switch {
	case newNode.Key < n.Key:
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.put(newNode)
		}
	default:
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.put(newNode)
		}
	}
}

func (n *Node) find(key int) (interface{}, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case key < n.Key:
		return n.Left.find(key)
	case key > n.Key:
		return n.Right.find(key)
	}
	return n.Value, true
}

func (n *Node) remove(key int, parent *Node) *Node {

	if n == nil {
		return nil
	}

	switch {
	case key < n.Key:
		return n.Left.remove(key, n)
	case key > n.Key:
		return n.Right.remove(key, n)
	}

	if n.Left == nil && n.Right == nil {
		n.replace(parent, nil)
		return n
	}

	if n.Left == nil {
		n.replace(parent, n.Right)
		return n
	}

	if n.Right == nil {
		n.replace(parent, n.Left)
		return n
	}

	replacement := n.Left.max()

	n.Key = replacement.Key
	n.Value = replacement.Value

	return replacement.remove(replacement.Key, parent)
}

func (n *Node) replace(parent *Node, newNode *Node) {

	if n == nil {
		return
	}

	if n == parent.Left {
		parent.Left = newNode
		return
	}

	parent.Right = newNode
}

func (n *Node) smallestRightNode() *Node {
	for {
		if n != nil && n.Left != nil {
			n = n.Left
		} else {
			break
		}
	}
	return n
}

func (n *Node) max() *Node {
	if n == nil {
		return nil
	}

	if n.Right == nil {
		return n
	}

	return n.Right.max()
}

func (n *Node) min() *Node {
	if n == nil {
		return nil
	}

	if n.Left == nil {
		return n
	}

	return n.Left.min()
}

func (n *Node) balance(s *Node) {
	// TODO: BALANCE TREE
}
