func verticalTraversal(root *TreeNode) [][]int {
    type Node struct {
        val int
        row int
        col int
    }

    nodes := make([]Node, 0)

    var traverse func(node *TreeNode, row, col int)
    traverse = func(node *TreeNode, row, col int) {
        if node == nil {
            return
        }

        nodes = append(nodes, Node{
            val: node.Val,
            row: row,
            col: col,
        })

        traverse(node.Left, row+1, col-1)
        traverse(node.Right, row+1, col+1)
    }

    traverse(root, 0, 0)

    sort.Slice(nodes, func(i, j int) bool {
        if nodes[i].col != nodes[j].col {
            return nodes[i].col < nodes[j].col
        }
        if nodes[i].row != nodes[j].row {
            return nodes[i].row < nodes[j].row
        }
        return nodes[i].val < nodes[j].val
    })

    result := make([][]int, 0)
    currentCol := math.MinInt

    for _, node := range nodes {
        if node.col != currentCol {
            result = append(result, []int{})
            currentCol = node.col
        }

        result[len(result)-1] = append(result[len(result)-1], node.val)
    }

    return result
}