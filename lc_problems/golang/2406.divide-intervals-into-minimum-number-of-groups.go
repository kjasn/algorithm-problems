/*
 * @lc app=leetcode.cn id=2406 lang=golang
 * @lcpr version=30204
 *
 * [2406] 将区间分为最少组数
 *
 * https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/description/
 *
 * algorithms
 * Medium (49.33%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    12.4K
 * Total Submissions: 25.2K
 * Testcase Example:  '[[5,10],[6,8],[1,5],[2,3],[1,10]]'
 *
 * 给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示 闭 区间 [lefti,
 * righti] 。
 *
 * 你需要将 intervals 划分为一个或者多个区间 组 ，每个区间 只 属于一个组，且同一个组中任意两个区间 不相交 。
 *
 * 请你返回 最少 需要划分成多少个组。
 *
 * 如果两个区间覆盖的范围有重叠（即至少有一个公共数字），那么我们称这两个区间是 相交 的。比方说区间 [1, 5] 和 [5, 8] 相交。
 *
 *
 *
 * 示例 1：
 *
 * 输入：intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
 * 输出：3
 * 解释：我们可以将区间划分为如下的区间组：
 * - 第 1 组：[1, 5] ，[6, 8] 。
 * - 第 2 组：[2, 3] ，[5, 10] 。
 * - 第 3 组：[1, 10] 。
 * 可以证明无法将区间划分为少于 3 个组。
 *
 *
 * 示例 2：
 *
 * 输入：intervals = [[1,3],[5,6],[8,10],[11,13]]
 * 输出：1
 * 解释：所有区间互不相交，所以我们可以把它们全部放在一个组内。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * 1 <= lefti <= righti <= 10^6
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import (
	"container/heap"
	"slices"
	"sort"
)

// @lcpr-template-end
// @lc code=start
func minGroups(intervals [][]int) int {
	// 左端点升序
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	h := hp{}
	for _, v := range intervals {
		if h.Len() == 0 || v[0] <= h.IntSlice[0] { // 堆为空 或 不同组
			heap.Push(&h, v[1])
		} else { // 同组 修改堆顶
			h.IntSlice[0] = v[1]
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}

// define min-heap
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// @lc code=end

/*
// @lcpr case=start
// [[5,10],[6,8],[1,5],[2,3],[1,10]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[5,6],[8,10],[11,13]]\n
// @lcpr case=end

*/
