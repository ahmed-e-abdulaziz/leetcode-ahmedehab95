/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil { return nil }
	copyMap := map[*Node]*Node{} // Keep track of copy nodes
	createCopy(node, copyMap)
	return copyMap[node]
}

func createCopy(node *Node, copyMap map[*Node]*Node) {
	copyNode, ok := copyMap[node] // get the copy node to add its neighbors then
	if !ok { // if the copy doesn't exist create it                
		copyNode = &Node{Val: node.Val}
		copyMap[node] = copyNode
	}
	for _, n := range node.Neighbors { 
		if _, ok := copyMap[n]; !ok {
			createCopy(n, copyMap) // For each neighbor that's not visited, do dfs to create the copy
		}
		copyNode.Neighbors = append(copyNode.Neighbors, copyMap[n]) // Add the copy neighbor (copyMap[n]) from map 
	}
}