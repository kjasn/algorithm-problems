// @before-stub-for-debug-begin
#include <vector>
#include <string>
#include "commoncppproblem849.h"

using namespace std;
// @before-stub-for-debug-end

/*
 * @lc app=leetcode.cn id=849 lang=cpp
 *
 * [849] 到最近的人的最大距离
 *
 * https://leetcode.cn/problems/maximize-distance-to-closest-person/description/
 *
 * algorithms
 * Medium (44.22%)
 * Likes:    239
 * Dislikes: 0
 * Total Accepted:    32K
 * Total Submissions: 65.3K
 * Testcase Example:  '[1,0,0,0,1,0,1]'
 *
 * 给你一个数组 seats 表示一排座位，其中 seats[i] = 1 代表有人坐在第 i 个座位上，seats[i] = 0 代表座位 i
 * 上是空的（下标从 0 开始）。
 * 
 * 至少有一个空座位，且至少有一人已经坐在座位上。
 * 
 * 亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。
 * 
 * 返回他到离他最近的人的最大距离。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：seats = [1,0,0,0,1,0,1]
 * 输出：2
 * 解释：
 * 如果亚历克斯坐在第二个空位（seats[2]）上，他到离他最近的人的距离为 2 。
 * 如果亚历克斯坐在其它任何一个空位上，他到离他最近的人的距离为 1 。
 * 因此，他到离他最近的人的最大距离是 2 。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：seats = [1,0,0,0]
 * 输出：3
 * 解释：
 * 如果亚历克斯坐在最后一个座位上，他离最近的人有 3 个座位远。
 * 这是可能的最大距离，所以答案是 3 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：seats = [0,1]
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * seats[i] 为 0 或 1
 * 至少有一个 空座位
 * 至少有一个 座位上有人
 * 
 * 
 */

// @lc code=start

class Solution {
public:
    int maxDistToClosest(vector<int>& seats) {
        int dis = 1, sz = seats.size();
        int p1 = -1,p2 = -1;// -1 标识为初始化
        for(int i = 0; i < sz; i++) {
            if(seats[i]) {
                if(p2 != -1) {
                    p1 = p2, p2 = i;
                    dis = max(dis, (p2 - p1) / 2);
                }
                else {
                    p2 = i;dis = max(dis, p2);
                }
            }
        }

        if(p1 == -1) return max(p2, sz - 1 - p2);// 只有一人
        return max(dis, sz - 1 - p2);// 处理右边缘无人的情况
    }
};
// @lc code=end

