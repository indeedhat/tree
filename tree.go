package tree

import (
	"bytes"
	// "fmt"
	// "strings"
)

type Limb interface {
	Path() string
	String() string
}

type Tree struct {
	Root     *Branch
	Renderer *Renderer
}

// Create a new tree with the default values
func NewTree() *Tree {
	tree := &Tree{
		Renderer: DefaultRenderer(),
		Root: &Branch{
			Key:  "",
			Text: "root",
			Open: true,
		},
	}

	return tree
}

// Initialize the tree and attach parent pointers to the limbs
func (t *Tree) Plant() {
	t.plant(t.Root, nil)
}

func (t *Tree) plant(limb interface{}, parent *Branch) {
	switch b := limb.(type) {
	case *Branch:
		b.tree = t
		if parent != b.parent {
			b.parent = parent
		}

		t.plant(b.Limbs, b)

	case *Leaf:
		b.parent = parent

	case []Limb:
		for _, s := range b {
			t.plant(s, parent)
		}
	}
}

// Find a limb by its key path
func (t *Tree) Find(keyPath string) Limb {
	if "" == keyPath {
		return t.Root
	}

	return find(t.Root, keyPath)
}

func find(subject interface{}, keyPath string) Limb {
	switch s := subject.(type) {
	case *Branch:
		if s.Path() == keyPath {
			return subject.(Limb)
		}
		return find(s.Limbs, keyPath)

	case *Leaf:
		if s.Path() == keyPath {
			return subject.(Limb)
		}

	case []Limb:
		for _, e := range s {
			if l := find(e, keyPath); nil != l {
				return l
			}
		}
	}

	return nil
}

// find a limb by its index
// optionally dont count the children of a closed branch to count
// against the index
func (t *Tree) FindByIndex(index int, skipClosed bool) Limb {
	var subject interface{}

	if t.Renderer.DisplayRoot {
		subject = t.Root
	} else {
		subject = t.Root.Limbs
		index += 1
	}

	found, _ := findByIndex(subject, index, skipClosed)
	return found
}

func findByIndex(subject interface{}, index int, skipClosed bool) (Limb, int) {

	switch s := subject.(type) {
	case *Branch:
		if 0 == index {
			return subject.(Limb), 0
		}

		if skipClosed && !s.Open {
			return nil, index
		}

		return findByIndex(s.Limbs, index, skipClosed)

	case *Leaf:
		if 0 == index {
			return subject.(Limb), 0
		}

		return nil, index

	case []Limb:
		var found Limb
		for _, e := range s {
			found, index = findByIndex(e, index-1, skipClosed)

			if nil != found {
				return found, 0
			}
		}
	}

	return nil, index
}

// Render the tree to string
func (t *Tree) Render() string {
	t.Renderer.buffer = bytes.NewBuffer([]byte{})

	if t.Renderer.DisplayRoot {
		t.Renderer.render(t.Root, "", 0, false)
	} else {
		t.Renderer.render(t.Root.Limbs, "", 0, 1 <= len(t.Root.Limbs))
	}

	return t.Renderer.buffer.String()
}
