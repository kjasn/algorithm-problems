/*
 * @lc app=leetcode.cn id=3242 lang=golang
 * @lcpr version=30204
 *
 * [3242] 设计相邻元素求和服务
 *
 * https://leetcode.cn/problems/design-neighbor-sum-service/description/
 *
 * algorithms
 * Easy (82.01%)
 * Likes:    11
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 11.2K
 * Testcase Example:  '["NeighborSum","adjacentSum","adjacentSum","diagonalSum","diagonalSum"]\n' +
  '[[[[0,1,2],[3,4,5],[6,7,8]]],[1],[4],[4],[8]]'
 *
 * 给你一个 n x n 的二维数组 grid，它包含范围 [0, n^2 - 1] 内的不重复元素。
 *
 * 实现 neighborSum 类：
 *
 *
 * neighborSum(int [][]grid) 初始化对象。
 * int adjacentSum(int value) 返回在 grid 中与 value 相邻的元素之和，相邻指的是与 value
 * 在上、左、右或下的元素。
 * int diagonalSum(int value) 返回在 grid 中与 value 对角线相邻的元素之和，对角线相邻指的是与 value
 * 在左上、右上、左下或右下的元素。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 *
 * ["neighborSum", "adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"]
 *
 * [[[[0, 1, 2], [3, 4, 5], [6, 7, 8]]], [1], [4], [4], [8]]
 *
 * 输出： [null, 6, 16, 16, 4]
 *
 * 解释：
 *
 *
 *
 *
 * 1 的相邻元素是 0、2 和 4。
 * 4 的相邻元素是 1、3、5 和 7。
 * 4 的对角线相邻元素是 0、2、6 和 8。
 * 8 的对角线相邻元素是 4。
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：
 *
 * ["neighborSum", "adjacentSum", "diagonalSum"]
 *
 * [[[[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]]], [15],
 * [9]]
 *
 * 输出： [null, 23, 45]
 *
 * 解释：
 *
 *
 *
 *
 * 15 的相邻元素是 0、10、7 和 6。
 * 9 的对角线相邻元素是 4、12、14 和 15。
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= n == grid.length == grid[0].length <= 10
 * 0 <= grid[i][j] <= n^2 - 1
 * 所有 grid[i][j] 值均不重复。
 * adjacentSum 和 diagonalSum 中的 value 均在范围 [0, n^2 - 1] 内。
 * 最多会调用 adjacentSum 和 diagonalSum 总共 2 * n^2 次。
 *
 *
*/

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
type NeighborSum struct {
	nums [][]int
	pos  map[int][]int
}

func Constructor(grid [][]int) NeighborSum {
	ret := new(NeighborSum)
	ret.pos = make(map[int][]int)
	for i, row := range grid {
		for j, col := range row {
			ret.pos[col] = []int{i, j}
		}
	}
	ret.nums = grid
	return *ret
}

func (this *NeighborSum) AdjacentSum(value int) int {
	x, y := 0, 0
	if v, ok := this.pos[value]; ok {
		x, y = v[0], v[1]
	} else {
		return -1
	}
	ans := 0
	if x > 0 {
		ans += this.nums[x-1][y]
	}

	if x < len(this.nums)-1 {
		ans += this.nums[x+1][y]
	}

	if y > 0 {
		ans += this.nums[x][y-1]
	}
	if y < len(this.nums[0])-1 {
		ans += this.nums[x][y+1]
	}

	return ans
}

func (this *NeighborSum) DiagonalSum(value int) int {
	x, y := 0, 0
	if v, ok := this.pos[value]; ok {
		x, y = v[0], v[1]
	} else {
		return -1
	}
	ans := 0
	if x > 0 {
		if y > 0 {
			ans += this.nums[x-1][y-1]
		}
		if y < len(this.nums[0])-1 {
			ans += this.nums[x-1][y+1]
		}
	}

	if x < len(this.nums)-1 {
		if y > 0 {
			ans += this.nums[x+1][y-1]
		}
		if y < len(this.nums[0])-1 {
			ans += this.nums[x+1][y+1]
		}
	}

	return ans
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
// @lc code=end

/*
// @lcpr case=start
// ["neighborSum", "adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"][[[[0, 1, 2], [3, 4, 5], [6, 7, 8]]], [1], [4], [4], [8]]\n
// @lcpr case=end

// @lcpr case=start
// ["neighborSum", "adjacentSum", "diagonalSum"][[[[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]]], [15], [9]]\n
// @lcpr case=end

*/
