/*
 * @lc app=leetcode.cn id=1738 lang=golang
 * @lcpr version=30204
 *
 * [1738] 找出第 K 大的异或坐标值
 *
 * https://leetcode.cn/problems/find-kth-largest-xor-coordinate-value/description/
 *
 * algorithms
 * Medium (68.02%)
 * Likes:    132
 * Dislikes: 0
 * Total Accepted:    46.4K
 * Total Submissions: 68.2K
 * Testcase Example:  '[[5,2],[1,6]]\n1'
 *
 * 给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。
 *
 * 矩阵中坐标 (a, b) 的 目标值 可以通过对所有元素 matrix[i][j] 执行异或运算得到，其中 i 和 j 满足 0 <= i <= a <
 * m 且 0 <= j <= b < n（下标从 0 开始计数）。
 *
 * 请你找出 matrix 的所有坐标中第 k 大的目标值（k 的值从 1 开始计数）。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[5,2],[1,6]], k = 1
 * 输出：7
 * 解释：坐标 (0,1) 的目标值是 5 XOR 2 = 7 ，为最大的目标值。
 *
 * 示例 2：
 *
 * 输入：matrix = [[5,2],[1,6]], k = 2
 * 输出：5
 * 解释：坐标 (0,0) 的目标值是 5 = 5 ，为第 2 大的目标值。
 *
 * 示例 3：
 *
 * 输入：matrix = [[5,2],[1,6]], k = 3
 * 输出：4
 * 解释：坐标 (1,0) 的目标值是 5 XOR 1 = 4 ，为第 3 大的目标值。
 *
 * 示例 4：
 *
 * 输入：matrix = [[5,2],[1,6]], k = 4
 * 输出：0
 * 解释：坐标 (1,1) 的目标值是 5 XOR 2 XOR 1 XOR 6 = 0 ，为第 4 大的目标值。
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 1000
 * 0 <= matrix[i][j] <= 10^6
 * 1 <= k <= m * n
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import "slices"

// @lcpr-template-end
// @lc code=start
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	nums := make([]int, m*n)
	// 同一个数异或偶数次为 0
	for i := range matrix {
		for j := range matrix[i] {
			if i != 0 {
				matrix[i][j] ^= matrix[i-1][j]
			}
			if j != 0 {
				matrix[i][j] ^= matrix[i][j-1]
			}
			if i != 0 && j != 0 {
				matrix[i][j] ^= matrix[i-1][j-1]
			}
			nums = append(nums, matrix[i][j])
		}
	}

	// 降序排序
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	return nums[0+k-1]
}

// @lc code=end

/*
// @lcpr case=start
// [[5,2],[1,6]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[5,2],[1,6]]\n2\n
// @lcpr case=end

// @lcpr case=start
// [[5,2],[1,6]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[5,2],[1,6]]\n4\n
// @lcpr case=end

*/
