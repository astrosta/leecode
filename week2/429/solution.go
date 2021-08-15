/*
	N 叉树的层序遍历
	解题思路： 树的遍历可用递归实现，子节点的个数即为递归的次数，递归函数需用level参数记录递归深度，递归深度即为树的层级数
*/

package _29

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	//根节点层级为0
	childrenTraversal(root, 0, &result)

	return result
}

func childrenTraversal(root *Node, level int, result *[][]int) {
	if root == nil {
		return
	}

	//对二维切片扩容
	if len(*result) < level+1 {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val)

	//层级递增，然后遍历子节点
	level++
	for _, node := range root.Children {
		childrenTraversal(node, level, result)
	}

	return
}
