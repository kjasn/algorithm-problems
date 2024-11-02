/*
 * @lc app=leetcode.cn id=3226 lang=golang
 * @lcpr version=30204
 *
 * [3226] 使两个整数相等的位更改次数
 *
 * https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/description/
 *
 * algorithms
 * Easy (64.14%)
 * Likes:    16
 * Dislikes: 0
 * Total Accepted:    10.1K
 * Total Submissions: 14.3K
 * Testcase Example:  '13\n4'
 *
 * 给你两个正整数 n 和 k。
 *
 * 你可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。
 *
 * 返回使得 n 等于 k 所需要的更改次数。如果无法实现，返回 -1。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： n = 13, k = 4
 *
 * 输出： 2
 *
 * 解释：
 * 最初，n 和 k 的二进制表示分别为 n = (1101)2 和 k = (0100)2，
 *
 * 我们可以改变 n 的第一位和第四位。结果整数为 n = (0100)2 = k。
 *
 *
 * 示例 2：
 *
 *
 * 输入： n = 21, k = 21
 *
 * 输出： 0
 *
 * 解释：
 * n 和 k 已经相等，因此不需要更改。
 *
 *
 * 示例 3：
 *
 *
 * 输入： n = 14, k = 13
 *
 * 输出： -1
 *
 * 解释：
 * 无法使 n 等于 k。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n, k <= 10^6
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import (
	"math/bits"
)

// @lcpr-template-end
// @lc code=start
func minChanges(n int, k int) int {
	/**
	ans := 0
	// 每次操作只能将 n 为 1 的二进制位变成 0
	// 找 n 与 k 二进制位不同的位进行比较
	for max(n, k) != 0 {
		if k%2 == 1 {
			if n%2 != 1 {
				return -1
			}
		} else if n%2 == 1 { // n为1，k不为1
			ans++
		}
		k /= 2
		n /= 2
	}
	return ans
	*/

	// 位运算
	// k 中 1 的位置为 n 的子集， 否则无法将 n 变成 k
	// 操作次数为 n 与 k 中 1 的数量差
	if k&n != k {
		return -1
	}
	return bits.OnesCount(uint(k ^ n))
}

// @lc code=end

/*
// @lcpr case=start
// 13\n4\n
// @lcpr case=end

// @lcpr case=start
// 21\n21\n
// @lcpr case=end

// @lcpr case=start
// 14\n13\n
// @lcpr case=end

*/
