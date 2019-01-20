package btree

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}