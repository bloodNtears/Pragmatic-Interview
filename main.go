package main

type TreeNode struct {
	Value  string
	Parent *TreeNode
}

func findNearestCommonAncestor(node1, node2 *TreeNode) *TreeNode {
	visited := make(map[*TreeNode]bool)

	for node1 != nil {
		visited[node1] = true
		node1 = node1.Parent
	}
	
	for node2 != nil {
		if visited[node2] {
			return node2
		}
		node2 = node2.Parent
	}

	return nil
}


func main() {
	nodeR := &TreeNode{
		Value:  "R",
		Parent: nil,
	}
	nodeE := &TreeNode{
		Value:  "E",
		Parent: nodeR,
	}
	nodeF := &TreeNode{
		Value:  "F",
		Parent: nodeE,
	}
	nodeA := &TreeNode{
		Value:  "A",
		Parent: nodeF,
	}
	//nodeG := &TreeNode{
	//	Value: "G",
	//	Parent: nodeF,
	//}
	nodeB := &TreeNode{
		Value:  "B",
		Parent: nodeE,
	}

	commonAncestor := findNearestCommonAncestor(nodeA, nodeB)
	if commonAncestor != nil {
		println("The nearest common ancestor is node:", commonAncestor.Value)
	} else {
		println("No common ancestor found.")
	}
}
