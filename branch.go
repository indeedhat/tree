package tree

import (
	"bytes"
	"fmt"
)

type Branch struct {
	Limbs []Limb
	Open  bool
	Key   string
	Text  string

	parent *Branch
	tree   *Tree
}

// build up a string path to the current limb
func (b *Branch) Path() string {
	path := ""

	if nil != b.parent {
		path += (*b.parent).Path()

		if "" != path {
			path += "."
		}
	}

	return path + b.Key
}

// return the string representation of the limb
func (b *Branch) String() string {
	var buffer bytes.Buffer

	buffer.WriteRune(pickBranchToggleState(b, b.tree))
	buffer.WriteRune(' ')
	if b.tree.CountOnLeft {
		buffer.WriteString(fmt.Sprintf("[%d] ", len(b.Limbs)))
	}

	buffer.WriteString(b.Text)
	if !b.tree.CountOnLeft {
		buffer.WriteString(fmt.Sprintf(" [%d]", len(b.Limbs)))
	}

	return buffer.String()
}

// toogle the open state of the branch
func (b *Branch) Toggle() {
	b.Open = !b.Open
}

// add a child limb at the given index
// -1 = append
func (b *Branch) AddChild(limb Limb, index int) {
	if -1 == index || index >= len(b.Limbs) {
		b.Limbs = append(b.Limbs, limb)
	} else {
		rest := b.Limbs[index:]
		b.Limbs = append(b.Limbs[:index], limb)
		b.Limbs = append(b.Limbs, rest...)
	}

	b.tree.plant(limb, b)
}

// remove a child limb at the given index
func (b *Branch) RemoveChild(index int) {
	b.Limbs = append(b.Limbs[:index], b.Limbs[index+1:]...)
}

// grow self and all child branches
func (b *Branch) GrowChildren() {
	b.Open = true
	toggleLimb(b.Limbs, true)
}

// trim self and all child branches
func (b *Branch) TrimChildren() {
	b.Open = false
	toggleLimb(b.Limbs, false)
}

func pickBranchToggleState(b *Branch, t *Tree) rune {
	if b.Open {
		return t.TrimMarker
	}

	return t.GrowMarker
}

func toggleLimb(l interface{}, open bool) {
	switch b := l.(type) {
	case *Branch:
		b.Open = open
		toggleLimb(b.Limbs, open)
	case []Limb:
		for _, e := range b {
			toggleLimb(e, open)
		}
	}
}
