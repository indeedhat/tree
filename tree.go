package tree

import (
	"bytes"
	"strings"
)

type Limb interface {
	Path() string
	String() string
}

type Tree struct {
	GrowMarker  rune
	TrimMarker  rune
	Indent      string
	Root        *Branch
	CountOnLeft bool
	DisplayRoot bool

	buffer *bytes.Buffer
}

// Create a new tree with the default values
func NewTree() *Tree {
	tree := &Tree{
		TrimMarker: TRIM_MARKER,
		GrowMarker: GROW_MARKER,
		Indent:     INDENT,
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
	if t.DisplayRoot {
		t.Root.Key = "root"
	}

	t.plant(t.Root, nil)
}

// Find a limb by its key path
func (t *Tree) Find(keyPath string) *Limb {
	return find(t.Root, keyPath)
}

func find(subject interface{}, keyPath string) *Limb {
	switch s := subject.(type) {
	case Branch:
	case *Branch:
		if s.Path() == keyPath {
			return subject.(*Limb)
		}
		return find(s.Limbs, keyPath)

	case Leaf:
	case *Leaf:
		if s.Path() == keyPath {
			return subject.(*Limb)
		}

	case []interface{}:
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
func (t *Tree) FindByIndex(index int, skipClosed bool) *Limb {
	var subject interface{}

	if t.DisplayRoot {
		subject = t.Root
	} else {
		subject = t.Root.Limbs
	}

	found, _ := findByIndex(subject, index, skipClosed)
	return found
}

func findByIndex(subject interface{}, index int, skipClosed bool) (*Limb, int) {
	switch s := subject.(type) {
	case Branch:
	case *Branch:
		if 0 == index {
			return subject.(*Limb), 0
		}

		if skipClosed && !s.Open {
			return nil, index - 1
		}

		return findByIndex(s.Limbs, index-1, skipClosed)

	case Leaf:
	case *Leaf:
		if 0 == index {
			return subject.(*Limb), 0
		}

		return nil, index - 1

	case []interface{}:
		var found *Limb

		for _, e := range s {
			found, index = findByIndex(e, index, skipClosed)

			if nil != found {
				return found, 0
			}
		}
	}

	return nil, -1
}

// Render the tree to string
func (t *Tree) Render() string {
	t.buffer = bytes.NewBuffer([]byte{})

	if t.DisplayRoot {
		t.render(t.Root, 0)
	} else {
		t.render(t.Root.Limbs, 0)
	}

	return t.buffer.String()
}

func (t *Tree) render(limb interface{}, depth int) {
	switch b := limb.(type) {
	case Branch:
	case *Branch:
		t.buffer.WriteString(b.String())
		t.buffer.WriteRune('\n')
		if b.Open {
			t.render(b.Limbs, depth+1)
		}

	case Leaf:
	case *Leaf:
		t.buffer.WriteString(b.String())
		t.buffer.WriteRune('\n')

	case []interface{}:
		for _, s := range b {
			if 0 < depth {
				t.buffer.WriteString(strings.Repeat(t.Indent, depth))
			}
			t.render(s, depth)
		}
	}
}

func (t *Tree) plant(limb interface{}, parent *Branch) {
	switch b := limb.(type) {
	case Branch:
	case *Branch:
		b.tree = t
		b.parent = parent
		t.plant(b.Limbs, b)

	case Leaf:
	case *Leaf:
		b.parent = parent

	case []interface{}:
		for _, s := range b {
			t.plant(s, parent)
		}
	}
}
