package main

type TreeNode struct {
	Value  string
	Parent *TreeNode
}

func findNearestCommonAncestor(node1, node2 *TreeNode) *TreeNode {
	if node1 == node2 {
		return node1
	}
	depth1 := findDepth(node1)
	depth2 := findDepth(node2)
	for depth1 > depth2 {
		node1 = node1.Parent
		depth1--
	}
	for depth2 > depth1 {
		node2 = node2.Parent
		depth2--
	}
	
	for node1 != node2 {
		node1 = node1.Parent
		node2 = node2.Parent
	}
	return node1
}

func findDepth(node *TreeNode) int {
	depth := 0
	for node != nil {
		depth++
		node = node.Parent
	}
	return depth
}

func main() {
	nodeR := &TreeNode{Value: "R", Parent: nil}
	nodeE := &TreeNode{Value: "E", Parent: nodeR}
	nodeF := &TreeNode{Value: "F", Parent: nodeE}
	nodeA := &TreeNode{Value: "A", Parent: nodeF}
	nodeB := &TreeNode{Value: "B", Parent: nodeE}

	commonAncestor := findNearestCommonAncestor(nodeA, nodeB)
	if commonAncestor != nil {
		println("The nearest common ancestor is node:", commonAncestor.Value)
	} else {
		println("No common ancestor found.")
	}
}
