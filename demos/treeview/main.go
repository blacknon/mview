// Demo code for the TreeView primitive.
package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

// Show a navigable tree view of the current directory.
func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	rootDir := "."
	root := mview.NewTreeNode(rootDir)
	root.SetColor(tcell.ColorRed.TrueColor())
	tree := mview.NewTreeView()
	tree.SetRoot(root)
	tree.SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *mview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := mview.NewTreeNode(file.Name())
			node.SetReference(filepath.Join(path, file.Name()))
			node.SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen.TrueColor())
			}
			target.AddChild(node)
		}
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *mview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	app.SetRoot(tree, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
