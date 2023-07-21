package main

type TreeNode struct {
	Value  string
	Parent *TreeNode
}

func findNearestCommonAncestor(node1, node2 *TreeNode) *TreeNode {

	isRoot := func(node *TreeNode) bool {
		return node.Parent == nil
	}

	for !isRoot(node1) {
		currentNode := node2
		for !isRoot(currentNode) {
			if currentNode.Parent == node1.Parent {
				return currentNode.Parent
			}
			currentNode = currentNode.Parent
		}
		node1 = node1.Parent
	}

	return nil
}

func main() {
	nodeR := &TreeNode{Value: "R", Parent: nil}
	//nodeG := &TreeNode{Value: "G", Parent: nodeR}
	nodeE := &TreeNode{Value: "E", Parent: nodeR}
	nodeF := &TreeNode{Value: "F", Parent: nodeE}
	nodeA := &TreeNode{Value: "A", Parent: nodeF}
	nodeB := &TreeNode{Value: "B", Parent: nodeE}
	//nodeS := &TreeNode{Value: "S", Parent: nil}

	commonAncestor := findNearestCommonAncestor(nodeA, nodeB)
	if commonAncestor != nil {
		println("The nearest common ancestor is node:", commonAncestor.Value)
	} else {
		println("No common ancestor found.")
	}
}
