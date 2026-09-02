/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := make(map[int]*Node)

    var traverse func(node *Node) *Node
    traverse = func(node *Node) *Node {
        if node == nil {
            return nil
        }
        if v, exist := visited[node.Val]; exist {
            return v
        }

        newNode := &Node{
            Val: node.Val,
            Neighbors: make([]*Node, 0),
        }
        visited[node.Val] = newNode

        for _, adj := range node.Neighbors {
            newNode.Neighbors = append(newNode.Neighbors, traverse(adj))
        }

        return newNode
    }

    return traverse(node)
}