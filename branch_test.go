package tree

import (
	// "reflect"
	"testing"
)

func TestBranchString(t *testing.T) {
	expected := "+ TestBranch [0]"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchEmptyString(t *testing.T) {
	expected := "+  [0]"
	tree := NewTree()

	branch := &Branch{
		Text: "",
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchStringOpen(t *testing.T) {
	expected := "- TestBranch [0]"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
		Open: true,
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchStringWithChildren(t *testing.T) {
	expected := "+ TestBranch [2]"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
		Limbs: []interface{}{
			&Leaf{},
			&Leaf{},
		},
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchStringWithLeftChildCount(t *testing.T) {
	expected := "+ [2] TestBranch"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
		Limbs: []interface{}{
			&Leaf{},
			&Leaf{},
		},
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.CountOnLeft = true
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchPath(t *testing.T) {
	expected := "branchKey"

	branch := &Leaf{
		Key: expected,
	}

	if branch.Path() != expected {
		t.Error("Branch path method returned the wrong val")
	}
}

func TestBranchNestedPath(t *testing.T) {
	expected := "parentBranch.branchKey"

	parentBranch := &Branch{
		Key: "parentBranch",
	}

	branch := &Branch{
		Key:    "branchKey",
		parent: parentBranch,
	}

	parentBranch.Limbs = append(parentBranch.Limbs, branch)

	if branch.Path() != expected {
		t.Error("Nested branch path method returned the wrong val")
	}
}

func TestBranchToggle(t *testing.T) {
	open := true
	branch := &Branch{}

	if open == branch.Open {
		t.Error("Branch should be closed by default")
	}

	branch.Toggle()
	if open != branch.Open {
		t.Error("Failed to toggle branch to open state")
	}

	branch.Toggle()
	if open == branch.Open {
		t.Error("Failed to toggle branch to closed state")
	}
}

func TestBranchGrowChildren(t *testing.T) {
	open := true
	child := &Branch{}
	branch := &Branch{}
	branch.Limbs = append(branch.Limbs, child)

	if open == child.Open || open == branch.Open {
		t.Error("Initial state of tree is not as expected")
	}

	branch.GrowChildren()
	if open != branch.Open || open != child.Open {
		t.Error("Failed to open self or children")
	}
}

func TestAddChild(t *testing.T) {
	expected := []*Leaf{
		&Leaf{},
		&Leaf{},
		&Leaf{},
		&Leaf{},
		&Leaf{},
	}

	branch := &Branch{}
	branch.Limbs = []interface{}{
		expected[1],
		expected[3],
	}

	branch.AddChild(expected[0], 0)
	if expected[0] != branch.Limbs[0] {
		t.Errorf("failed to add to position 0")
	}

	// 	branch.AddChild(expected[0], -1)
	// 	if !reflect.DeepEqual(expected[3], branch.Limbs[3]) {
	// 		t.Error("failed to add to last position")
	// 	}

	// 	branch.AddChild(expected[0], 2)
	// 	if !reflect.DeepEqual(expected[2], branch.Limbs[2]) {
	// 		t.Error("failed to add to position 0")
	// 	}

	// 	for i, _ := range expected {
	// 		if !reflect.DeepEqual(expected[i], branch.Limbs[i]) {
	// 			t.Error("final slice does not match expected")
	// 		}
	// 	}
}
