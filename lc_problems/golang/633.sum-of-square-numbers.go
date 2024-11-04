/*
 * @lc app=leetcode.cn id=633 lang=golang
 * @lcpr version=30204
 *
 * [633] 平方数之和
 *
 * https://leetcode.cn/problems/sum-of-square-numbers/description/
 *
 * algorithms
 * Medium (37.96%)
 * Likes:    487
 * Dislikes: 0
 * Total Accepted:    161.4K
 * Total Submissions: 419.4K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：c = 5
 * 输出：true
 * 解释：1 * 1 + 2 * 2 = 5
 *
 *
 * 示例 2：
 *
 * 输入：c = 3
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= c <= 2^31 - 1
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import "math"

// @lcpr-template-end
// @lc code=start
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c/2; a++ {
		b := int(math.Sqrt(float64(c - a*a)))
		if a*a+b*b == c {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
