package main

type TreeNode struct {
	Value  string
	Parent *TreeNode
}

func findNearestCommonAncestor(node1, node2 *TreeNode) *TreeNode {
	ptr1, ptr2 := node1, node2

	for ptr1 != ptr2 {
		ptr1 = getNextNode(ptr1, node2)
		ptr2 = getNextNode(ptr2, node1)
	}

	return ptr1
}

func getNextNode(currentNode, targetNode *TreeNode) *TreeNode {
	if currentNode == nil {
		return targetNode
	}
	return currentNode.Parent
}

func main() {
	nodeR := &TreeNode{Value: "R", Parent: nil}
	nodeE := &TreeNode{Value: "E", Parent: nodeR}
	nodeF := &TreeNode{Value: "F", Parent: nodeE}
	nodeA := &TreeNode{Value: "A", Parent: nodeF}
	nodeB := &TreeNode{Value: "B", Parent: nodeE}
	//nodeS := &TreeNode{Value: "S", Parent: nil}

	commonAncestor := findNearestCommonAncestor(nodeA, nodeB)
	if commonAncestor != nil {
		println("The nearest common ancestor is node:", commonAncestor.Value)
	} else {
		println("No common ancestor found (or common ancestor is tree's root).")
	}
}
