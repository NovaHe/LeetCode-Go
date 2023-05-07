package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 优化版本，先右子树后左子树
func rightSideView(root *TreeNode) []int {
	var dfs func(root *TreeNode, level int)
	res := []int{}
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) == level {
			res = append(res, root.Val)
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)

	}
	dfs(root, 0)
	return res
}

func rightSideView1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}
	return res
}
