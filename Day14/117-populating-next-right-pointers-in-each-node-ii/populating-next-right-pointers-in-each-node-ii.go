/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {return nil}
	q := make([]*Node, 0)
    q = append(q, root)

    for len(q) > 0 {
        nNodeInTheSameLevel := len(q)
        var prev *Node
        for i:=0;i<nNodeInTheSameLevel;i++{
            fr := q[0]
            q = q[1:]

            if prev != nil {
                prev.Next = fr
            }

            if fr.Left != nil {
                q = append(q, fr.Left)
            }
            if fr.Right != nil {
                q = append(q, fr.Right)
            }
            prev = fr
        }
    }
    return root
}