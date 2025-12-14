package tree

import (
	"errors"
)

type Leaf interface {
	GetID() string
}

type Node[LeafT Leaf] struct {
	Parent *Node[LeafT]
	Item   LeafT
	Leaves []*Node[LeafT]
}

func NewNode[LeafT Leaf](
	item LeafT,
	leaves ...LeafT,
) *Node[LeafT] {
	node := &Node[LeafT]{
		Item: item,
	}
	return node.AddLeaves(leaves...)
}
func (n *Node[LeafT]) Attach(parent *Node[LeafT]) *Node[LeafT] {
	if parent == nil {
		return n
	}
	parent.AddNodes(n)
	n.Parent = parent
	return n
}

func (n *Node[LeafT]) AddLeaves(leaf ...LeafT) *Node[LeafT] {
	for _, l := range leaf {
		n.Leaves = append(n.Leaves, NewNode[LeafT](l).Attach(n))
	}
	return n
}

func (n *Node[LeafT]) AddNodes(nodes ...*Node[LeafT]) *Node[LeafT] {
	for _, l := range nodes {
		l.Parent = n
		n.Leaves = append(n.Leaves, l)
	}
	return n
}

func (n *Node[LeafT]) FindLeaf(id string) (LeafT, error) {
	node := n.FindNode(id)
	if node == nil {
		var def LeafT
		return def, errors.New("could not find a node")
	}
	return node.Item, nil
}

func (n *Node[LeafT]) FindNode(id string) *Node[LeafT] {
	if n.Item.GetID() == id {
		return n
	}
	for _, l := range n.Leaves {
		if l.Item.GetID() == id {
			return l
		}
		res := l.FindNode(id)
		if res != nil {
			return res
		}
	}
	return nil
}

func (n *Node[LeafT]) FindParent(id string) *Node[LeafT] {
	if n.Parent == nil {
		return nil
	}
	return n.Parent.FindParent(id)
}
