package tree

import (
	"testing"
)

func TestLeafString(t *testing.T) {
	expected := "TestLeaf"

	leaf := &Leaf{
		Text: expected,
	}

	if leaf.String() != expected {
		t.Error("Leaf string method returned the wrong val")
	}
}

func TestLeafEmptyString(t *testing.T) {
	expected := ""

	leaf := &Leaf{
		Text: expected,
	}

	if leaf.String() != expected {
		t.Error("Leaf string method returned the wrong val")
	}
}

func TestLeafPath(t *testing.T) {
	expected := "leafkey"

	leaf := &Leaf{
		Key: expected,
	}

	if leaf.Path() != expected {
		t.Error("Leaf path method returned the wrong val")
	}
}

func TestLeafNestedPath(t *testing.T) {
	expected := "parentBranch.leafKey"

	branch := &Branch{
		Key: "parentBranch",
	}

	leaf := &Leaf{
		Key:    "leafKey",
		parent: branch,
	}

	branch.Limbs = append(branch.Limbs, leaf)

	if leaf.Path() != expected {
		t.Error("Nested leaf path method returned the wrong val")
	}
}
