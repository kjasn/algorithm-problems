/*
 * @lc app=leetcode.cn id=863 lang=golang
 * @lcpr version=30204
 *
 * [863] 二叉树中所有距离为 K 的结点
 *
 * https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/description/
 *
 * algorithms
 * Medium (61.46%)
 * Likes:    712
 * Dislikes: 0
 * Total Accepted:    60.1K
 * Total Submissions: 97.7K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n2'
 *
 * 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k ，返回到目标结点 target 距离为 k
 * 的所有结点的值的数组。
 *
 * 答案可以以 任何顺序 返回。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
 * 输出：[7,4,1]
 * 解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1
 *
 *
 * 示例 2:
 *
 * 输入: root = [1], target = 1, k = 3
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * 节点数在 [1, 500] 范围内
 * 0 <= Node.val <= 500
 * Node.val 中所有值 不同
 * 目标结点 target 是树上的结点。
 * 0 <= k <= 1000
 *
 *
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	parents := make(map[int]*TreeNode)

	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}

	// 找到每个结点的父结点
	findParents(root)

	var findAnswer func(*TreeNode, int, int)
	findAnswer = func(node *TreeNode, from, depth int) {
		if node == nil {
			return
		}

		// 深度为 k，添加
		if depth == k {
			ans = append(ans, node.Val)
		}
		// 向下找子结点
		if node.Left != nil && node.Left.Val != from {
			findAnswer(node.Left, node.Val, depth+1)
		}
		if node.Right != nil && node.Right.Val != from {
			findAnswer(node.Right, node.Val, depth+1)
		}
		// 向上扩展
		if parent, ok := parents[node.Val]; ok && parent.Val != from {
			findAnswer(parent, node.Val, depth+1)
		}
	}

	findAnswer(target, -1, 0)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n3\n
// @lcpr case=end

*/
