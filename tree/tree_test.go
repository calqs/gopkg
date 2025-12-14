package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testLeaf struct {
	id string
}

func (l testLeaf) GetID() string {
	return l.id
}

func Test_I_Can_Find_Leaves(t *testing.T) {
	root := NewNode(testLeaf{"1"}).AddNodes(
		NewNode(testLeaf{"2"}).AddNodes(
			NewNode(testLeaf{"3"}).AddLeaves(
				testLeaf{"4"},
				testLeaf{"4.1"},
			),
			NewNode(testLeaf{"3.1"}),
		),
		NewNode(testLeaf{"2.1"}),
	)

	leaf, err := root.FindLeaf("4")
	assert.Nil(t, err)
	assert.Equal(t, "4", leaf.GetID())
	assert.Equal(t, "4.1", root.FindNode("4").Parent.FindNode("4.1").Item.GetID())
	assert.Equal(t, "2.1", root.FindNode("4").Parent.Parent.Parent.FindNode("2.1").Item.GetID())
	assert.Nil(t, root.FindNode("4").Parent.Parent.Parent.Parent)
}
