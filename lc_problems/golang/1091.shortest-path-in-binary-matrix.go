/*
 * @lc app=leetcode.cn id=1091 lang=golang
 * @lcpr version=30204
 *
 * [1091] 二进制矩阵中的最短路径
 *
 * https://leetcode.cn/problems/shortest-path-in-binary-matrix/description/
 *
 * algorithms
 * Medium (40.82%)
 * Likes:    378
 * Dislikes: 0
 * Total Accepted:    86.5K
 * Total Submissions: 211.8K
 * Testcase Example:  '[[0,1],[1,0]]'
 *
 * 给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
 *
 * 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n -
 * 1)）的路径，该路径同时满足下述要求：
 *
 *
 * 路径途经的所有单元格的值都是 0 。
 * 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
 *
 *
 * 畅通路径的长度 是该路径途经的单元格总数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：grid = [[0,1],[1,0]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 * 输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
 * 输出：4
 *
 *
 * 示例 3：
 *
 * 输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length
 * n == grid[i].length
 * 1 <= n <= 100
 * grid[i][j] 为 0 或 1
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	check := func(x, y int) bool {
		if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 0 {
			return true
		}
		return false
	}

	if grid[0][0] != 0 {
		return -1
	}

	que := [][2]int{{0, 0}}
	for ans := 1; len(que) > 0; ans++ {
		for k := len(que); k > 0; k-- {
			curX, curY := que[0][0], que[0][1]
			que = que[1:]
			// 到达右下角
			if curX == n-1 && curY == n-1 {
				return ans
			}

			for nx := curX - 1; nx <= curX+1; nx++ {
				for ny := curY - 1; ny <= curY+1; ny++ {
					if check(nx, ny) {
						grid[nx][ny] = 1 // 标记走过
						que = append(que, [2]int{nx, ny})
					}
				}
			}

		}
	}

	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[1,1,0],[1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[1,1,0],[1,1,0]]\n
// @lcpr case=end

*/
