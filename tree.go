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
	GrowMarker  string
	TrimMarker  string
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

	if t.DisplayRoot {
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
	t.buffer = bytes.NewBuffer([]byte{})

	if t.DisplayRoot {
		t.render(t.Root, "", false)
	} else {
		t.render(t.Root.Limbs, INDENT_BLANK, 1 < len(t.Root.Limbs))
	}

	return t.buffer.String()
}

func (t *Tree) render(limb interface{}, prefix string, hasRemaining bool) {
	switch b := limb.(type) {
	case *Branch:
		t.buffer.WriteString(b.String())
		t.buffer.WriteRune('\n')
		if b.Open {
			if hasRemaining {
				t.render(b.Limbs, prefix+INDENT, hasRemaining)
			} else {
				t.render(b.Limbs, prefix+INDENT_BLANK, hasRemaining)
			}
		}

	case *Leaf:
		t.buffer.WriteString(b.String())
		t.buffer.WriteRune('\n')

	case []Limb:
		for i, s := range b {
			if "" != prefix {
				t.buffer.WriteString(prefix[4:])
				if len(b) == i+1 {
					t.buffer.WriteString(INDENT_END)
				} else {
					t.buffer.WriteString(INDENT_MIDDLE)
				}

			}
			t.render(s, prefix, i < len(b)-1)
		}
	}
}
