# Tree
[![Actions Status](https://github.com/indeedhat/tree/workflows/Go/badge.svg)](https://github.com/indeedhat/tree/actions)


- [x] hack together basic rendering
- [x] build path
    - [x] link limbs to parent
- [x] find limb by index
- [x] find limb by key path
- [x] move toggle indicator to String method on branch
- [x] show child count on branch 
- [x] allow toggle indicator to be on left or right of text
- [ ] design a nicer way of initiating tree
- [ ] normalise the api (everything trees)
- [.] pull render element out into its own loop instead of building in String method
- [x] fix bug with find by index (show vs hide root)
- [x] branch/tree methods
    - [x] expand children
    - [x] collapse children
    - [x] add child
    - [x] remove child
- [x] all the tests
    - [x] Tree
        - [x] Plant
        - [x] Find
        - [x] FindByIndex
        - [x] Render
    - [x] Branch
        - [x] Path
        - [x] String
        - [x] Toggle
        - [x] AddChild
        - [x] RemoveChild
        - [x] GrowChildren
        - [x] TrimChildren
    - [x] Leaf
        - [x] Path
        - [x] String

