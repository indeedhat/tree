package tree

import (
	// "fmt"
	"os"
	"testing"
)

var tree = NewTree()

func TestMain(m *testing.M) {
	tree.Root.Limbs = []Limb{
		&Leaf{
			Key:  "k1",
			Text: "key 1",
		},
		&Branch{
			Key:  "k2",
			Text: "key 2",
			Open: true,
			Limbs: []Limb{
				&Leaf{
					Key:  "k3",
					Text: "Key 3",
				},
				&Branch{
					Key:  "k4",
					Text: "key 4",
					Open: false,
					Limbs: []Limb{
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

	tree.Plant()
	code := m.Run()

	os.Exit(code)
}

func TestTree(t *testing.T) {
	expected := `key 1
- key 2 [3]
  Key 3
  + key 4 [1]
  key 6
`
	if tree.Render() != expected {
		t.Error("Failed to render standard tree")
	}
}

func TestTreeWithRoot(t *testing.T) {
	expected := `- root [2]
  key 1
  - key 2 [3]
    Key 3
    + key 4 [1]
    key 6
`

	tree.DisplayRoot = true
	if tree.Render() != expected {
		t.Error("Failed to render root tree")
	}
}

func TestTreeWithLeftCount(t *testing.T) {
	expected := `- [2] root
  key 1
  - [3] key 2
    Key 3
    + [1] key 4
    key 6
`

	tree.DisplayRoot = true
	tree.CountOnLeft = true
	if tree.Render() != expected {
		t.Error("Failed to render left count tree")
	}
}

func TestTreeWithCustomCharacters(t *testing.T) {
	expected := `^ [2] root
....key 1
....^ [3] key 2
........Key 3
........v [1] key 4
........key 6
`

	tree.DisplayRoot = true
	tree.CountOnLeft = true
	tree.TrimMarker = '^'
	tree.GrowMarker = 'v'
	tree.Indent = "...."

	if tree.Render() != expected {
		t.Error("Failed to render root tree")
	}
}

func TestFindBranch(t *testing.T) {
	branch, ok := tree.Root.Limbs[1].(*Branch)
	if !ok || branch != tree.Find("k2") {
		t.Error("Failed to find branch by key")
	}
}

func TestFindLeaf(t *testing.T) {
	leaf := tree.Root.Limbs[1].(*Branch).Limbs[1].(*Branch).Limbs[0]
	if leaf != tree.Find("k2.k4.k5") {
		t.Error("Failed to find leaf by key")
	}
}

func TestFindRoot(t *testing.T) {
	if tree.Find("") != tree.Root {
		t.Error("Failed to find root by path")
	}
}

func TestDoesntFindBadPath(t *testing.T) {
	if nil != tree.Find("some.bad.path") {
		t.Error("Find returned a value of a bad path")
	}
}

func TestFindByIndexNoRoot(t *testing.T) {
	tree.DisplayRoot = false
	if tree.FindByIndex(0, false) != tree.Root.Limbs[0] {
		t.Error("Failed to find limb by index with hidden root")
	}
}

func TestFindByIndexWithRoot(t *testing.T) {
	tree.DisplayRoot = true
	if tree.FindByIndex(0, false) != tree.Root {
		t.Error("Failed to find limb by index when index is visible")
	}
}

func TestFindByIndexSkipHidden(t *testing.T) {
	tree.DisplayRoot = true

	leaf := tree.Root.Limbs[1].(*Branch).Limbs[2]
	if tree.FindByIndex(5, true) != leaf {
		t.Errorf("Failed to find limb when skipping closed branches %v - %v", leaf, tree.FindByIndex(5, true))
	}
}

func TestFindByIndexIncludeHidden(t *testing.T) {
	tree.DisplayRoot = true

	leaf := tree.Root.Limbs[1].(*Branch).Limbs[1].(*Branch).Limbs[0]
	if tree.FindByIndex(5, false) != leaf {
		t.Errorf("Failed to find limb when including closed branches %v - %v", leaf, tree.FindByIndex(5, false))
	}
}

func TestDoesntFindBadIndex(t *testing.T) {
	if nil != tree.FindByIndex(10, false) {
		t.Error("Find retuned a value for an invalid index")
	}
}
