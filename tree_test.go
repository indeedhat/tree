package tree

import (
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
					Open: true,
					Limbs: []Limb{
						&Leaf{
							Key:  "k5",
							Text: "key 5",
						},
						&Leaf{
							Key:  "k6",
							Text: "key 6",
						},
						&Branch{
							Key:  "k9",
							Text: "key 9",
							Open: true,
							Limbs: []Limb{
								&Leaf{
									Key:  "k10",
									Text: "key 10",
								},
								&Leaf{
									Key:  "k11",
									Text: "key 11",
								},
								&Branch{
									Key:  "k12",
									Text: "key 12",
									Open: false,
									Limbs: []Limb{
										&Leaf{
											Key:  "k13",
											Text: "key 13",
										},
									},
								},
							},
						},
					},
				},
				&Leaf{
					Key:  "k7",
					Text: "key 7",
				},
			},
		},
		&Leaf{
			Key:  "k8",
			Text: "key 8",
		},
	}

	tree.Plant()
	code := m.Run()

	os.Exit(code)
}

func TestTree(t *testing.T) {
	expected := `[-] root (3)
 ├─ key 1
 ├─ [-] key 2 (3)
 │   ├─ Key 3
 │   ├─ [-] key 4 (3)
 │   │   ├─ key 5
 │   │   ├─ key 6
 │   │   └─ [-] key 9 (3)
 │   │       ├─ key 10
 │   │       ├─ key 11
 │   │       └─ [+] key 12 (1)
 │   └─ key 7
 └─ key 8
`
	tree.DisplayRoot = true
	if tree.Render() != expected {
		t.Error("Failed to render standard tree")
	}
}

func TestTreeNoRoot(t *testing.T) {
	expected := ` ├─ key 1
 ├─ [-] key 2 (3)
 │   ├─ Key 3
 │   ├─ [-] key 4 (3)
 │   │   ├─ key 5
 │   │   ├─ key 6
 │   │   └─ [-] key 9 (3)
 │   │       ├─ key 10
 │   │       ├─ key 11
 │   │       └─ [+] key 12 (1)
 │   └─ key 7
 └─ key 8
`

	tree.DisplayRoot = false
	if tree.Render() != expected {
		t.Error("Failed to render no root tree")
	}
}

func TestTreeWithLeftCount(t *testing.T) {
	expected := `[-] (3) root
 ├─ key 1
 ├─ [-] (3) key 2
 │   ├─ Key 3
 │   ├─ [-] (3) key 4
 │   │   ├─ key 5
 │   │   ├─ key 6
 │   │   └─ [-] (3) key 9
 │   │       ├─ key 10
 │   │       ├─ key 11
 │   │       └─ [+] (1) key 12
 │   └─ key 7
 └─ key 8
`

	tree.DisplayRoot = true
	tree.CountOnLeft = true
	if tree.Render() != expected {
		t.Error("Failed to render left count tree")
	}
}

func TestTreeWithCustomCharacters(t *testing.T) {
	expected := `^ root (3)
 ├─ key 1
 ├─ ^ key 2 (3)
 │   ├─ Key 3
 │   ├─ ^ key 4 (3)
 │   │   ├─ key 5
 │   │   ├─ key 6
 │   │   └─ ^ key 9 (3)
 │   │       ├─ key 10
 │   │       ├─ key 11
 │   │       └─ v key 12 (1)
 │   └─ key 7
 └─ key 8
`

	tree.DisplayRoot = true
	tree.TrimMarker = "^"
	tree.GrowMarker = "v"
	tree.CountOnLeft = false

	if tree.Render() != expected {
		t.Error("Failed to render custom marker tree tree")
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
	if tree.FindByIndex(11, true) != leaf {
		t.Errorf("Failed to find limb when skipping closed branches %v - %v", leaf, tree.FindByIndex(11, true))
	}
}

func TestFindByIndexSkipHiddenNoRoot(t *testing.T) {
	tree.DisplayRoot = false

	leaf := tree.Root.Limbs[1].(*Branch).Limbs[2]
	if tree.FindByIndex(10, true) != leaf {
		t.Errorf("Failed to find limb when skipping closed branches %v - %v", leaf, tree.FindByIndex(10, true))
	}
}

func TestFindByIndexIncludeHidden(t *testing.T) {
	tree.DisplayRoot = true

	leaf := tree.Root.Limbs[1].(*Branch).Limbs[1].(*Branch).Limbs[2].(*Branch).Limbs[2].(*Branch).Limbs[0]
	if tree.FindByIndex(11, false) != leaf {
		t.Errorf("Failed to find limb when including closed branches %v - %v", leaf, tree.FindByIndex(11, false))
	}
}

func TestFindByIndexIncludeHiddenNoRoot(t *testing.T) {
	tree.DisplayRoot = false

	leaf := tree.Root.Limbs[1].(*Branch).Limbs[1].(*Branch).Limbs[2].(*Branch).Limbs[2].(*Branch).Limbs[0]
	if tree.FindByIndex(10, false) != leaf {
		t.Errorf("Failed to find limb when including closed branches %v - %v", leaf, tree.FindByIndex(10, false))
	}
}

func TestDoesntFindBadIndex(t *testing.T) {
	if nil != tree.FindByIndex(30, false) {
		t.Error("Find retuned a value for an invalid index")
	}
}
