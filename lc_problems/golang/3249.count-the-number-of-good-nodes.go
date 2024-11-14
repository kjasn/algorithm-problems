/*
 * @lc app=leetcode.cn id=3249 lang=golang
 * @lcpr version=30204
 *
 * [3249] 统计好节点的数目
 *
 * https://leetcode.cn/problems/count-the-number-of-good-nodes/description/
 *
 * algorithms
 * Medium (52.40%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    13.1K
 * Total Submissions: 21.4K
 * Testcase Example:  '[[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]'
 *
 * 现有一棵 无向 树，树中包含 n 个节点，按从 0 到 n - 1 标记。树的根节点是节点 0 。给你一个长度为 n - 1 的二维整数数组
 * edges，其中 edges[i] = [ai, bi] 表示树中节点 ai 与节点 bi 之间存在一条边。
 *
 * 如果一个节点的所有子节点为根的 子树 包含的节点数相同，则认为该节点是一个 好节点。
 *
 * 返回给定树中 好节点 的数量。
 *
 * 子树 指的是一个节点以及它所有后代节点构成的一棵树。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
 *
 * 输出：7
 *
 * 说明：
 *
 * 树的所有节点都是好节点。
 *
 *
 * 示例 2：
 *
 *
 * 输入：edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]
 *
 * 输出：6
 *
 * 说明：
 *
 * 树中有 6 个好节点。上图中已将这些节点着色。
 *
 *
 * 示例 3：
 *
 *
 * 输入：edges =
 * [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]
 *
 * 输出：12
 *
 * 解释：
 *
 * 除了节点 9 以外其他所有节点都是好节点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 10^5
 * edges.length == n - 1
 * edges[i].length == 2
 * 0 <= ai, bi < n
 * 输入确保 edges 总表示一棵有效的树。
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
func countGoodNodes(edges [][]int) int {
	ans := 0
	n := len(edges) + 1
	tree := make([][]int, n)
	for _, v := range edges {
		a, b := v[0], v[1]
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	var dfs func(int, int) int
	dfs = func(root, father int) int {
		if len(tree[root]) == 0 {
			ans++
			return 1
		}

		tot, sz0 := 1, 0 // root 树的所有子节点个数  第一棵子树的子节点个数
		flag := true
		for _, b := range tree[root] {
			if b == father {
				continue
			}
			sz := dfs(b, root)
			if sz0 == 0 {
				sz0 = sz
			} else if sz != sz0 {
				flag = false
			}
			tot += sz
		}
		if flag {
			ans++
		}
		return tot
	}

	dfs(0, -1)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]\n
// @lcpr case=end

*/
