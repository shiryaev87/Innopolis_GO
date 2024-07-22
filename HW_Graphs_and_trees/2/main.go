package main

import (
	"fmt"
	"sort"
)

type Node struct {
	values   []int
	children []*Node
}

type TwoThreeTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{
		values:   []int{value},
		children: nil,
	}
}

func NewTwoThreeTree() *TwoThreeTree {
	return &TwoThreeTree{root: nil}
}

func (tree *TwoThreeTree) Search(value int) bool {
	return search(tree.root, value)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if contains(node.values, value) {
		return true
	}

	for i, v := range node.values {
		if value < v {
			if i < len(node.children) {
				return search(node.children[i], value)
			}
			return false
		}
	}

	if len(node.children) > len(node.values) {
		return search(node.children[len(node.values)], value)
	}
	return false
}

func (tree *TwoThreeTree) Insert(value int) {
	if tree.root == nil {
		tree.root = NewNode(value)
	} else {
		newRoot, newChild := insert(tree.root, value)
		if newChild != nil {
			newRootNode := &Node{
				values:   []int{newRoot},
				children: []*Node{tree.root, newChild},
			}
			tree.root = newRootNode
		}
	}
}

func insert(node *Node, value int) (int, *Node) {
	if node.children == nil {
		node.values = append(node.values, value)
		sort.Ints(node.values)
		if len(node.values) < 3 {
			return 0, nil
		} else {
			return split(node)
		}
	}

	for i, v := range node.values {
		if value < v {
			newRoot, newChild := insert(node.children[i], value)
			if newChild != nil {
				node.values = append(node.values, newRoot)
				node.children = append(node.children[:i+1], node.children[i:]...)
				node.children[i+1] = newChild
				sort.Ints(node.values)
				if len(node.values) < 3 {
					return 0, nil
				} else {
					return split(node)
				}
			}
			return 0, nil
		}
	}

	newRoot, newChild := insert(node.children[len(node.values)], value)
	if newChild != nil {
		node.values = append(node.values, newRoot)
		node.children = append(node.children, newChild)
		sort.Ints(node.values)
		if len(node.values) < 3 {
			return 0, nil
		} else {
			return split(node)
		}
	}
	return 0, nil
}

func split(node *Node) (int, *Node) {
	middle := node.values[1]
	newNode := &Node{
		values:   []int{node.values[2]},
		children: nil,
	}
	node.values = []int{node.values[0]}
	if node.children != nil {
		newNode.children = append(newNode.children, node.children[2:]...)
		node.children = append(node.children[:2])
	}
	return middle, newNode
}

func contains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func (tree *TwoThreeTree) Delete(value int) {
	if tree.root != nil {
		tree.root, _ = delete(tree.root, value)
		if len(tree.root.values) == 0 && len(tree.root.children) > 0 {
			tree.root = tree.root.children[0]
		}
	}
}

func delete(node *Node, value int) (*Node, bool) {
	if node == nil {
		return nil, false
	}

	index := findIndex(node.values, value)
	if index < len(node.values) && node.values[index] == value {
		if node.children == nil {
			node.values = append(node.values[:index], node.values[index+1:]...)
			return node, len(node.values) == 0
		}

		largest := findMax(node.children[index])
		node.values[index] = largest
		//var childNode *Node
		var merged bool
		node.children[index], merged = delete(node.children[index], largest)
		if merged {
			node = merge(node, index)
		}
		return node, false
	}

	if node.children == nil {
		return node, false
	}

	//var childNode *Node
	var merged bool
	if index < len(node.children) {
		node.children[index], merged = delete(node.children[index], value)
	}
	if merged {
		node = merge(node, index)
	}
	return node, false
}

func findIndex(values []int, value int) int {
	for i, v := range values {
		if value <= v {
			return i
		}
	}
	return len(values)
}

func findMax(node *Node) int {
	for len(node.children) > 0 {
		node = node.children[len(node.children)-1]
	}
	return node.values[len(node.values)-1]
}

func merge(node *Node, index int) *Node {
	if index < len(node.values) {
		node.children[index].values = append(node.children[index].values, node.values[index])
		node.children[index].values = append(node.children[index].values, node.children[index+1].values...)
		node.children[index].children = append(node.children[index].children, node.children[index+1].children...)
		node.values = append(node.values[:index], node.values[index+1:]...)
		node.children = append(node.children[:index+1], node.children[index+2:]...)
	} else {
		index--
		node.children[index].values = append(node.children[index].values, node.values[index])
		node.children[index].values = append(node.children[index].values, node.children[index+1].values...)
		node.children[index].children = append(node.children[index].children, node.children[index+1].children...)
		node.values = append(node.values[:index], node.values[index+1:]...)
		node.children = append(node.children[:index+1], node.children[index+2:]...)
	}
	if len(node.values) == 0 {
		return node.children[0]
	}
	return node
}

// PrintTree outputs the tree structure
func (tree *TwoThreeTree) PrintTree() {
	printNode(tree.root, 0)
}

func printNode(node *Node, level int) {
	if node != nil {
		fmt.Printf("%s%v\n", string(' '+level*2), node.values)
		for _, child := range node.children {
			printNode(child, level+1)
		}
	}
}

func main() {
	tree := NewTwoThreeTree()
	values := []int{10, 20, 5, 6, 15, 30, 25}

	for _, v := range values {
		tree.Insert(v)
	}

	tree.PrintTree()
	fmt.Println()

	fmt.Println(tree.Search(15))  // true
	fmt.Println(tree.Search(100)) // false

	tree.Delete(6)
	tree.PrintTree()
	fmt.Println()

	tree.Delete(20)
	tree.PrintTree()
	fmt.Println()

	tree.Delete(100)
	tree.PrintTree()
	fmt.Println()
}
