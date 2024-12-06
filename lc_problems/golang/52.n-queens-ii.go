/*
 * @lc app=leetcode.cn id=52 lang=golang
 * @lcpr version=30204
 *
 * [52] N 皇后 II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.45%)
 * Likes:    548
 * Dislikes: 0
 * Total Accepted:    172.6K
 * Total Submissions: 208.2K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 *
 *
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
func totalNQueens(n int) int {
	ans := 0
	colValid := make([]bool, n)
	diagValid := make([]bool, 2*n)
	antiDiagValid := make([]bool, 2*n)

	isValid := func(row, col, preCol int) bool {
		if colValid[col] {
			return false
		}
		// 对角线
		// y - x = b
		if diagValid[col-row+n] {
			return false
		}
		// 斜对角线
		// y + x = b
		if antiDiagValid[row+col] {
			return false
		}
		return true
	}
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row == n {
			ans++
			return
		}

		for i := 0; i < n; i++ {
			if !isValid(row, i, col) {
				continue
			}
			colValid[i] = true
			diagValid[i-row+n] = true
			antiDiagValid[row+i] = true
			dfs(row+1, i)
			colValid[i] = false
			diagValid[i-row+n] = false
			antiDiagValid[row+i] = false
		}

		return
	}

	dfs(0, -1)

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
