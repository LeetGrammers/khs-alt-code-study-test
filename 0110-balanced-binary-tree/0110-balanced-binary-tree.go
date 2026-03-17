/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, answer := dfs(root)
    return answer
}

func dfs(node *TreeNode) (int, bool) {
    if node == nil {
        return 0, true
    }
    left, cl := dfs(node.Left)
    right, cr := dfs(node.Right)
    if cl == false || cr == false {
        return 0, false
    }

    return max(left, right) + 1, abs(left - right) < 2
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}