package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewTree()
	tree.Root.Limbs = []interface{}{
		&Leaf{
			Key:  "k1",
			Text: "key 1",
		},
		&Branch{
			Key:  "k2",
			Text: "key 2",
			Open: true,
			Limbs: []interface{}{
				&Leaf{
					Key:  "k3",
					Text: "Key 3",
				},
				&Branch{
					Key:  "k4",
					Text: "key 4",
					Open: false,
					Limbs: []interface{}{
						&Leaf{
							Key:  "k5",
							Text: "key 5",
						},
					},
				},
				&Leaf{
					Key:  "k6",
					Text: "key 6",
				},
			},
		},
	}
	tree.CountOnLeft = true
	tree.DisplayRoot = true
	tree.Root.TrimChildren()

	tree.Plant()
	fmt.Println(tree.Render())
}
