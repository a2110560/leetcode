package tree

type Node struct {
	Val   interface{}
	Right *Node
	Left  *Node
}

func NewNode(val interface{}) *Node {
	node := new(Node)
	node.Val = val
	return node
}

func (n *Node) AddLeftNode(l *Node) {
	n.Left = l
}

func (n *Node) AddRightNode(r *Node) {
	n.Right = r
}

func GenerateTree() {
	/*
					3
			7				1
		2		6				9
			5		11		4
	*/
	root := NewNode(3)
	node7 := NewNode(7)
	node1 := NewNode(1)
	node2 := NewNode(2)
	node6 := NewNode(6)
	node9 := NewNode(9)
	node5 := NewNode(5)
	node11 := NewNode(11)
	node4 := NewNode(4)

	root.AddLeftNode(node7)
	root.AddRightNode(node1)
	node7.AddLeftNode(node2)
	node7.AddRightNode(node6)
	node6.AddLeftNode(node5)
	node6.AddRightNode(node11)
	node1.AddRightNode(node9)
	node9.AddLeftNode(node4)
}
