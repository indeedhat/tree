package tree

import (
	"testing"
)

func TestBranchString(t *testing.T) {
	expected := "[+] TestBranch (0)"
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
	expected := "[+]  (0)"
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
	expected := "[-] TestBranch (0)"
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
	expected := "[+] TestBranch (2)"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
		Limbs: []Limb{
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
	expected := "[+] (2) TestBranch"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
		Limbs: []Limb{
			&Leaf{},
			&Leaf{},
		},
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Renderer.Count.Position = Left
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchStringWithNoCount(t *testing.T) {
	expected := "[+] TestBranch"
	tree := NewTree()

	branch := &Branch{
		Text: "TestBranch",
		Limbs: []Limb{
			&Leaf{},
			&Leaf{},
		},
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Renderer.Count.Visible = false
	tree.Plant()

	if branch.String() != expected {
		t.Errorf("Branch string method returned the wrong val: %s - %s", expected, branch.String())
	}
}

func TestBranchStringWithToggleOnRight(t *testing.T) {
	expected := "TestBranch [+]"
	tree := NewTree()

	tree.Renderer.Toggle.Position = Right
	branch := &Branch{
		Text: "TestBranch",
		Limbs: []Limb{
			&Leaf{},
			&Leaf{},
		},
	}

	tree.Root.Limbs = append(tree.Root.Limbs, branch)
	tree.Renderer.Count.Visible = false
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

func TestBranchTrimChildren(t *testing.T) {
	open := true
	child := &Branch{Open: open}
	branch := &Branch{Open: open}
	branch.Limbs = append(branch.Limbs, child)

	if open != child.Open || open != branch.Open {
		t.Error("Initial state of tree is not as expected")
	}

	branch.TrimChildren()
	if open == branch.Open || open == child.Open {
		t.Error("Failed to open self or children")
	}
}

func TestAddChild(t *testing.T) {
	expected := []Limb{
		&Leaf{},
		&Leaf{},
		&Leaf{},
		&Leaf{},
		&Leaf{},
	}

	branch := &Branch{}
	branch.Limbs = []Limb{
		expected[1],
		expected[3],
	}

	branch.AddChild(expected[0], 0)
	if expected[0] != branch.Limbs[0] {
		t.Errorf("failed to add to position 0")
	}

	branch.AddChild(expected[4], -1)
	if expected[4] != branch.Limbs[3] {
		t.Error("failed to add to last position")
	}

	branch.AddChild(expected[2], 2)
	if expected[2] != branch.Limbs[2] {
		t.Error("failed to add to position 2")
	}
}

func TestRemoveChild(t *testing.T) {
	expected := []Limb{
		&Leaf{},
		&Leaf{},
		&Leaf{},
		&Leaf{},
		&Leaf{},
	}

	branch := &Branch{}
	branch.Limbs = append(branch.Limbs, expected...)

	branch.RemoveChild(2)
	branch.RemoveChild(0)
	branch.RemoveChild(2)

	if 2 != len(branch.Limbs) {
		t.Error("Failed to remove children")
	}

	if expected[1] != branch.Limbs[0] || expected[3] != branch.Limbs[1] {
		t.Error("Failed to remove the correct children")
	}
}
