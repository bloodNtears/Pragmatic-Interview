package main

type TreeNode struct {
	Value  string
	Parent *TreeNode
}

func isDescendant(node, target *TreeNode) bool {
	if node == nil {
		return false
	}
	if node == target {
		return true
	}
	return isDescendant(node.Parent, target)
}

func findNearestCommonAncestor(node1, node2 *TreeNode) *TreeNode {
	if node1 == nil || node2 == nil {
		return nil
	}

	currentNode := node1
	for currentNode != nil {
		if isDescendant(node2, currentNode) {
			return currentNode
		}
		currentNode = currentNode.Parent
	}

	return nil
}

func main() {
	nodeR := &TreeNode{Value: "R", Parent: nil}
	nodeE := &TreeNode{Value: "E", Parent: nodeR}
	nodeF := &TreeNode{Value: "F", Parent: nodeE}
	nodeA := &TreeNode{Value: "A", Parent: nodeF}
	nodeB := &TreeNode{Value: "B", Parent: nodeE}
	nodeG := &TreeNode{Value: "G", Parent: nodeB}
	nodeK := &TreeNode{Value: "K", Parent: nodeG}

	commonAncestor := findNearestCommonAncestor(nodeA, nodeK)
	if commonAncestor != nil {
		println("The nearest common ancestor is node:", commonAncestor.Value)
	} else {
		println("No common ancestor found.")
	}
}
