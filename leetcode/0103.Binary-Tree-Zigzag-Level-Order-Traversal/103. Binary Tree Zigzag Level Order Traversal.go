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

// dfs
func zigzagLevelOrder2(root *TreeNode) [][]int {
	var dfs func(root *TreeNode, level int)
	res := [][]int{}
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) > level {
			if level&1 == 0 {
				res[level] = append(res[level], root.Val)
			} else {
				res[level] = append([]int{root.Val}, res[level]...)
			}

		} else {
			res = append(res, []int{root.Val})
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}

// bfs
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		lay := make([]int, 0, size)
		if depth&1 == 1 {
			for i := size - 1; i >= 0; i-- {
				lay = append(lay, q[i].Val)
			}
		} else {
			for i := 0; i < size; i++ {
				lay = append(lay, q[i].Val)
			}
		}
		res = append(res, lay)
		q = q[size:]
		depth++
	}
	return res
}
