/*
 * @lc app=leetcode.cn id=3240 lang=golang
 * @lcpr version=30204
 *
 * [3240] 最少翻转次数使二进制矩阵回文 II
 *
 * https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/description/
 *
 * algorithms
 * Medium (31.24%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    6.6K
 * Total Submissions: 13.3K
 * Testcase Example:  '[[1,0,0],[0,1,0],[0,0,1]]'
 *
 * 给你一个 m x n 的二进制矩阵 grid 。
 *
 * 如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。
 *
 * 你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。
 *
 * 请你返回 最少 翻转次数，使得矩阵中 所有 行和列都是 回文的 ，且矩阵中 1 的数目可以被 4 整除 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,0,0],[0,1,0],[0,0,1]]
 *
 * 输出：3
 *
 * 解释：
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,1],[0,1],[0,0]]
 *
 * 输出：2
 *
 * 解释：
 *
 *
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1],[1]]
 *
 * 输出：2
 *
 * 解释：
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m * n <= 2 * 10^5
 * 0 <= grid[i][j] <= 1
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0

	// 四个镜像位置
	for i, row := range grid[:m/2] {
		r := grid[m-1-i]
		for j, x := range row[:n/2] {
			cnt1 := x + row[n-1-j] + r[n-1-j] + r[j]
			ans += min(cnt1, 4-cnt1)
		}
	}

	// 行列均为 奇数	中心必须为 0
	if m%2 == 1 && n%2 == 1 {
		ans += grid[m/2][n/2]
	}

	// 分别计算 行 列 为奇数时 为满足 1 的个数为 4 的倍数时的操作
	diff, cnt1 := 0, 0
	if m%2 == 1 {
		row := grid[m/2]
		for i, x := range row[:n/2] {
			if x != row[n-1-i] {
				diff++
			} else {
				cnt1 += x * 2
			}
		}
	}

	if n%2 == 1 {
		for i := 0; i < m/2; i++ {
			if grid[i][n/2] != grid[m-1-i][n/2] {
				diff++
			} else {
				cnt1 += grid[i][n/2] * 2
			}
		}
	}

	if diff > 0 {
		ans += diff
	} else {
		ans += cnt1 % 4
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[1,0,0],[0,1,0],[0,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,1],[0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1],[1]]\n
// @lcpr case=end

*/
