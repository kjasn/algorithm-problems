/*
 * @lc app=leetcode.cn id=2033 lang=golang
 * @lcpr version=30204
 *
 * [2033] 获取单值网格的最小操作数
 *
 * https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/description/
 *
 * algorithms
 * Medium (45.09%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    8K
 * Total Submissions: 17.7K
 * Testcase Example:  '[[2,4],[6,8]]\n2'
 *
 * 给你一个大小为 m x n 的二维整数网格 grid 和一个整数 x 。每一次操作，你可以对 grid 中的任一元素 加 x 或 减 x 。
 *
 * 单值网格 是全部元素都相等的网格。
 *
 * 返回使网格化为单值网格所需的 最小 操作数。如果不能，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：grid = [[2,4],[6,8]], x = 2
 * 输出：4
 * 解释：可以执行下述操作使所有元素都等于 4 ：
 * - 2 加 x 一次。
 * - 6 减 x 一次。
 * - 8 减 x 两次。
 * 共计 4 次操作。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：grid = [[1,5],[2,3]], x = 1
 * 输出：5
 * 解释：可以使所有元素都等于 3 。
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：grid = [[1,2],[3,4]], x = 2
 * 输出：-1
 * 解释：无法使所有元素相等。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10^5
 * 1 <= m * n <= 10^5
 * 1 <= x, grid[i][j] <= 10^4
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import "math"

// @lcpr-template-end
// @lc code=start
func minOperations(grid [][]int, x int) int {
	const INF int = 1e4 + 10
	mn, mx := INF, 0
	rem := grid[0][0] % x // 余数

	st := make([]int, INF)
	for _, row := range grid {
		for _, val := range row {
			if val%x != rem { // 余数不同
				return -1
			}
			tmp := val / x // x 的倍数
			st[tmp]++
			mx = max(mx, tmp)
			mn = min(mn, tmp)
		}
	}

	ans := math.MaxInt
	abs := func(a int) int {
		if a < 0 {
			a = -a
		}
		return a
	}

	// 枚举给定数据范围内 x 的所有倍数
	for i := mn; i <= mx; i++ {
		cnt := 0
		for j := mn; j <= mx; j++ {
			if j == i {
				continue
			}
			cnt += abs(i-j) * st[j]
		}
		ans = min(ans, cnt)
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [[2,4],[6,8]]\n2\n
// @lcpr case=end

// @lcpr case=start
// [[1,5],[2,3]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,4]]\n2\n
// @lcpr case=end

*/
