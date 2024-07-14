package main

import (
	"strings"

	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

const treeAllCode = `[green]package[white] main

[green]import[white] [red]"github.com/blacknon/mview"[white]

[green]func[white] [yellow]main[white]() {
	$$$

	root := mview.[yellow]NewTreeNode[white]([red]"Root"[white]).
		[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"First child"[white]).
			[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"Grandchild A"[white])).
			[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"Grandchild B"[white]))).
		[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"Second child"[white]).
			[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"Grandchild C"[white])).
			[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"Grandchild D"[white]))).
		[yellow]AddChild[white](mview.[yellow]NewTreeNode[white]([red]"Third child"[white]))

	tree.[yellow]SetRoot[white](root).
		[yellow]SetCurrentNode[white](root)

	mview.[yellow]NewApplication[white]().
		[yellow]SetRoot[white](tree, true).
		[yellow]Run[white]()
}`

const treeBasicCode = `tree := mview.[yellow]NewTreeView[white]()`

const treeTopLevelCode = `tree := mview.[yellow]NewTreeView[white]().
		[yellow]SetTopLevel[white]([red]1[white])`

const treeAlignCode = `tree := mview.[yellow]NewTreeView[white]().
		[yellow]SetAlign[white](true)`

const treePrefixCode = `tree := mview.[yellow]NewTreeView[white]().
		[yellow]SetGraphics[white](false).
		[yellow]SetTopLevel[white]([red]1[white]).
		[yellow]SetPrefixes[white]([][green]string[white]{
			[red]"[red[]* "[white],
			[red]"[darkcyan[]- "[white],
			[red]"[darkmagenta[]- "[white],
		})`

type node struct {
	text     string
	expand   bool
	selected func()
	children []*node
}

var (
	tree          = mview.NewTreeView()
	treeNextSlide func()
	treeCode      = mview.NewTextView()
)

var rootNode = &node{
	text: "Root",
	children: []*node{
		{text: "Expand all", selected: func() { tree.GetRoot().ExpandAll() }},
		{text: "Collapse all", selected: func() {
			for _, child := range tree.GetRoot().GetChildren() {
				child.CollapseAll()
			}
		}},
		{text: "Hide root node", expand: true, children: []*node{
			{text: "Tree list starts one level down"},
			{text: "Works better for lists where no top node is needed"},
			{text: "Switch to this layout", selected: func() {
				tree.SetAlign(false)
				tree.SetTopLevel(1)
				tree.SetGraphics(true)
				tree.SetPrefixes(nil)
				treeCode.SetText(strings.Replace(treeAllCode, "$$$", treeTopLevelCode, -1))
			}},
		}},
		{text: "Align node text", expand: true, children: []*node{
			{text: "For trees that are similar to lists"},
			{text: "Hierarchy shown only in line drawings"},
			{text: "Switch to this layout", selected: func() {
				tree.SetAlign(true)
				tree.SetTopLevel(0)
				tree.SetGraphics(true)
				tree.SetPrefixes(nil)
				treeCode.SetText(strings.Replace(treeAllCode, "$$$", treeAlignCode, -1))
			}},
		}},
		{text: "Prefixes", expand: true, children: []*node{
			{text: "Best for hierarchical bullet point lists"},
			{text: "You can define your own prefixes per level"},
			{text: "Switch to this layout", selected: func() {
				tree.SetAlign(false)
				tree.SetTopLevel(1)
				tree.SetGraphics(false)
				tree.SetPrefixes([]string{"[red]* ", "[darkcyan]- ", "[darkmagenta]- "})
				treeCode.SetText(strings.Replace(treeAllCode, "$$$", treePrefixCode, -1))
			}},
		}},
		{text: "Basic tree with graphics", expand: true, children: []*node{
			{text: "Lines illustrate hierarchy"},
			{text: "Basic indentation"},
			{text: "Switch to this layout", selected: func() {
				tree.SetAlign(false)
				tree.SetTopLevel(0)
				tree.SetGraphics(true)
				tree.SetPrefixes(nil)
				treeCode.SetText(strings.Replace(treeAllCode, "$$$", treeBasicCode, -1))
			}},
		}},
		{text: "Next slide", selected: func() { treeNextSlide() }},
	}}

// TreeView demonstrates the tree view.
func TreeView(nextSlide func()) (title string, info string, content mview.Primitive) {
	treeNextSlide = nextSlide
	tree.SetBorder(true)
	tree.SetTitle("TreeView")

	// Add nodes.
	var add func(target *node) *mview.TreeNode
	add = func(target *node) *mview.TreeNode {
		node := mview.NewTreeNode(target.text)
		node.SetSelectable(target.expand || target.selected != nil)
		node.SetExpanded(target == rootNode)
		node.SetReference(target)
		if target.expand {
			node.SetColor(tcell.ColorLimeGreen.TrueColor())
		} else if target.selected != nil {
			node.SetColor(tcell.ColorRed.TrueColor())
		}
		for _, child := range target.children {
			node.AddChild(add(child))
		}
		return node
	}
	root := add(rootNode)
	tree.SetRoot(root)
	tree.SetCurrentNode(root)
	tree.SetSelectedFunc(func(n *mview.TreeNode) {
		original := n.GetReference().(*node)
		if original.expand {
			n.SetExpanded(!n.IsExpanded())
		} else if original.selected != nil {
			original.selected()
		}
	})

	treeCode.SetWrap(false)
	treeCode.SetDynamicColors(true)
	treeCode.SetText(strings.Replace(treeAllCode, "$$$", treeBasicCode, -1))
	treeCode.SetPadding(1, 1, 2, 0)

	flex := mview.NewFlex()
	flex.AddItem(tree, 0, 1, true)
	flex.AddItem(treeCode, codeWidth, 1, false)

	return "TreeView", "", flex
}
