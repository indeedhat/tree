package tree

type Leaf struct {
	Key  string
	Text string

	parent *Branch
}

// build string path to the current leaf
func (l *Leaf) Path() string {
	path := ""

	if nil != l.parent {
		path += (*l.parent).Path()
		path += "."
	}

	return path + l.Key
}

// build string representation of the leaf
func (l *Leaf) String() string {
	return l.Text
}
