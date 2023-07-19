package main

type TreeNode struct {
	Value    string
	Children []*TreeNode
}

func findNearestCommonAncestor(root *TreeNode, node1, node2 *TreeNode) *TreeNode {
	var pathToNode1, pathToNode2 []*TreeNode

	findPathToNode(root, node1, &pathToNode1)
	findPathToNode(root, node2, &pathToNode2)

	var commonAncestor *TreeNode

	for i := 0; i < len(pathToNode1) && i < len(pathToNode2); i++ {
		if pathToNode1[i] != pathToNode2[i] {
			break
		}
		commonAncestor = pathToNode1[i]
	}

	return commonAncestor
}

func findPathToNode(root, target *TreeNode, path *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root)

	if root == target {
		return true
	}

	for _, child := range root.Children {
		if findPathToNode(child, target, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func main() {
	root := &TreeNode{Value: "R"}
	root.Children = []*TreeNode{
		{Value: "C", Children: []*TreeNode{
			{Value: "F", Children: []*TreeNode{
				{Value: "A"},
				{Value: "G"},
			}},
			{Value: "B"},
		}},
		{Value: "E"},
		{Value: "D"},
	}

	nodeA := root.Children[0].Children[0].Children[0]
	nodeB := root.Children[0].Children[1]
	// nodeG := root.Children[0].Children[0].Children[1]
	// println(nodeA.Value, nodeB.Value, nodeG.Value)
	// println(root.Value, root.Children[0].Children[1].Value)

	commonAncestor := findNearestCommonAncestor(root, nodeA, nodeB)
	if commonAncestor != nil {
		println("The nearest common ancestor is node:", commonAncestor.Value)
	} else {
		println("No common ancestor found.")
	}
}
