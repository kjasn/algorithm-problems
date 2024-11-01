/*
 * @lc app=leetcode.cn id=2344 lang=golang
 * @lcpr version=30204
 *
 * [2344] 使数组可以被整除的最少删除次数
 *
 * https://leetcode.cn/problems/minimum-deletions-to-make-array-divisible/description/
 *
 * algorithms
 * Hard (53.80%)
 * Likes:    21
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 18.1K
 * Testcase Example:  '[2,3,2,4,3]\n[9,6,9,3,15]'
 *
 * 给你两个正整数数组 nums 和 numsDivide 。你可以从 nums 中删除任意数目的元素。
 *
 * 请你返回使 nums 中 最小 元素可以整除 numsDivide 中所有元素的 最少 删除次数。如果无法得到这样的元素，返回 -1 。
 *
 * 如果 y % x == 0 ，那么我们说整数 x 整除 y 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,3,2,4,3], numsDivide = [9,6,9,3,15]
 * 输出：2
 * 解释：
 * [2,3,2,4,3] 中最小元素是 2 ，它无法整除 numsDivide 中所有元素。
 * 我们从 nums 中删除 2 个大小为 2 的元素，得到 nums = [3,4,3] 。
 * [3,4,3] 中最小元素为 3 ，它可以整除 numsDivide 中所有元素。
 * 可以证明 2 是最少删除次数。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [4,3,6], numsDivide = [8,2,6,10]
 * 输出：-1
 * 解释：
 * 我们想 nums 中的最小元素可以整除 numsDivide 中的所有元素。
 * 没有任何办法可以达到这一目的。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length, numsDivide.length <= 10^5
 * 1 <= nums[i], numsDivide[i] <= 10^9
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import (
	"slices"
)

// @lcpr-template-end
// @lc code=start
func minOperations(nums []int, numsDivide []int) int {
	// a > b
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// 求最大公因数
	g := 0
	for _, num := range numsDivide {
		g = gcd(num, g)
	}

	slices.Sort(nums)
	for i, v := range nums {
		if g%v == 0 {
			return i
		}
	}

	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,2,4,3]\n[9,6,9,3,15]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,6]\n[8,2,6,10]\n
// @lcpr case=end

*/
