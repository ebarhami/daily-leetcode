func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    
    idx, val := 0, postorder[len(postorder)-1]
    for inorder[idx] != val {
        idx++
    }

    node := &TreeNode{
        Val: postorder[len(postorder)-1],
        Left: buildTree(inorder[:idx], postorder[:idx]),
        Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
    }
    return node
}

/*
inorder -> left + root + right
postorder -> left + right + root

*/