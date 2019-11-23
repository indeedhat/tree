package tree

import (
	"bytes"
	// "unicode/utf8"
)

type Position int

const (
	Left Position = iota
	Right
)

type IndentStyle struct {
	Spacer   string
	Vertical string
	First    string
	Middle   string
	Last     string
}

type ToggleStyle struct {
	Open     string
	Close    string
	Position Position
}

type CountStyle struct {
	Visible  bool
	Template string
	Position Position
}

type Renderer struct {
	DisplayRoot bool
	Indent      IndentStyle
	Toggle      ToggleStyle
	Count       CountStyle

	buffer *bytes.Buffer
}

func (r *Renderer) render(limb interface{}, prefix string, depth int, isLastEntry bool) {
	switch limb := limb.(type) {
	case *Branch:
		r.buffer.WriteString(r.Indent.First)
		r.buffer.WriteString(limb.String())
		r.buffer.WriteRune('\n')

		if !limb.Open {
			return
		}

		if isLastEntry {
			r.render(limb.Limbs, prefix+r.Indent.Vertical, depth+1, isLastEntry)
			// if 0 == depth {
			// } else {
			// 	r.render(limb.Limbs, prefix+r.Indent.Spacer, depth+1, isLastEntry)
			// }
		} else {
			if 0 == depth {
				r.render(limb.Limbs, prefix+r.Indent.Vertical, depth+1, isLastEntry)
			} else {
				r.render(limb.Limbs, prefix+r.Indent.Spacer, depth+1, isLastEntry)
			}
		}

	case *Leaf:
		r.buffer.WriteString(limb.String())
		r.buffer.WriteRune('\n')

	case []Limb:
		for i, entry := range limb {
			if "" != prefix {
				r.buffer.WriteString(prefix[len(r.Indent.Vertical):])

				if len(limb) == i+1 {
					r.buffer.WriteString(r.Indent.Last)
				} else {
					r.buffer.WriteString(r.Indent.Middle)
				}
			}

			r.render(entry, prefix, depth, i+1 < len(limb))
		}
	}
}

// func (r *Renderer) render(limb interface{}, prefix string, depth int, hasRemaining bool) {
// 	switch b := limb.(type) {
// 	case *Branch:
// 		r.buffer.WriteString(r.Indent.First)
// 		r.buffer.WriteString(b.String())
// 		r.buffer.WriteRune('\n')

// 		if b.Open {
// 			if hasRemaining {
// 				r.render(b.Limbs, prefix+r.Indent.Vertical, depth+1, hasRemaining)
// 			} else {
// 				r.render(b.Limbs, prefix+r.Indent.Spacer, depth+1, hasRemaining)
// 			}
// 		}

// 	case *Leaf:
// 		r.buffer.WriteString(b.String())
// 		r.buffer.WriteRune('\n')

// 	case []Limb:
// 		for i, s := range b {
// 			if "" != prefix && (r.DisplayRoot || r.Indent.Spacer != prefix) {
// 				// if len(prefix) < len(r.Indent.Vertical) || r.DisplayRoot {
// 				// 	r.buffer.WriteString(prefix[len(r.Indent.Spacer):])
// 				// } else {
// 				// 	r.buffer.WriteString(prefix[len(r.Indent.Vertical):])
// 				// }
// 				r.buffer.WriteString(prefix[utf8.RuneCountInString(r.Indent.Spacer):])

// 				if len(b) == i+1 {
// 					r.buffer.WriteString(r.Indent.Last)
// 				} else {
// 					r.buffer.WriteString(r.Indent.Middle)
// 				}
// 			}

// 			r.render(s, prefix, depth, i < len(b)-1)
// 		}
// 	}
// }

func DefaultRenderer() *Renderer {
	return &Renderer{
		DisplayRoot: true,
		Indent: IndentStyle{
			Spacer:   "    ",
			Vertical: " │  ",
			Middle:   " ├─ ",
			Last:     " └─ ",
			First:    "",
		},
		Toggle: ToggleStyle{
			Open:     "[+]",
			Close:    "[-]",
			Position: Left,
		},
		Count: CountStyle{
			Visible:  true,
			Template: "(%d)",
			Position: Right,
		},
	}
}

func MinimalRenderer() *Renderer {
	return &Renderer{
		Indent: IndentStyle{
			Spacer:   "  ",
			Vertical: "  ",
			Middle:   "  ",
			Last:     "  ",
			First:    "",
		},
		Toggle: ToggleStyle{
			Open:     "+",
			Close:    "-",
			Position: Left,
		},
		Count: CountStyle{
			Visible: false,
		},
	}
}
