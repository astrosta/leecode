/*
	二叉树的最近公共祖先
	解题思路：遍历二叉树，未找到对应子节点返回nil，当左右子节点都不为nil时，此节点为二叉树的最近公共祖先
*/

package _36

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}
