/*
 * @lc app=leetcode.cn id=1031 lang=golang
 * @lcpr version=30204
 *
 * [1031] 两个非重叠子数组的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/description/
 *
 * algorithms
 * Medium (66.01%)
 * Likes:    293
 * Dislikes: 0
 * Total Accepted:    29.3K
 * Total Submissions: 44.4K
 * Testcase Example:  '[0,6,5,2,2,5,1,9,4]\n1\n2'
 *
 * 给你一个整数数组 nums 和两个整数 firstLen 和 secondLen，请你找出并返回两个非重叠 子数组 中元素的最大和，长度分别为
 * firstLen 和 secondLen 。
 *
 * 长度为 firstLen 的子数组可以出现在长为 secondLen 的子数组之前或之后，但二者必须是不重叠的。
 *
 * 子数组是数组的一个 连续 部分。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
 * 输出：20
 * 解释：子数组的一种选择中，[9] 长度为 1，[6,5] 长度为 2。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
 * 输出：29
 * 解释：子数组的一种选择中，[3,8,1] 长度为 3，[8,9] 长度为 2。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
 * 输出：31
 * 解释：子数组的一种选择中，[5,6,0,9] 长度为 4，[0,3,8] 长度为 3。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= firstLen, secondLen <= 1000
 * 2 <= firstLen + secondLen <= 1000
 * firstLen + secondLen <= nums.length <= 1000
 * 0 <= nums[i] <= 1000
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

// @lcpr-template-end
// @lc code=start
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	sum := make([]int, n+1)

	// 前缀和
	for i := range nums {
		sum[i+1] += sum[i] + nums[i]
	}

	ans := 0
	// 假设连个数组分别为 A B
	// 一次遍历 记录前一个子数组的 最大值
	solve := func(l1, l2 int) {
		maxA := 0
		for i := l1 + l2; i <= n; i++ {
			maxA = max(maxA, sum[i-l2]-sum[i-l2-l1])
			ans = max(ans, maxA+sum[i]-sum[i-l2])
		}
	}

	// 分别计算 A前B后 A后B前
	solve(firstLen, secondLen) // A B
	solve(secondLen, firstLen) // B A

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [0,6,5,2,2,5,1,9,4]\n1\n2\n
// @lcpr case=end

// @lcpr case=start
// [3,8,1,3,2,1,8,9,0]\n3\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,1,5,6,0,9,5,0,3,8]\n4\n3\n
// @lcpr case=end

*/
