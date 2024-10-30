/*
 * @lc app=leetcode.cn id=1191 lang=cpp
 *
 * [1191] K 次串联后最大子数组之和
 *
 * https://leetcode.cn/problems/k-concatenation-maximum-sum/description/
 *
 * algorithms
 * Medium (27.01%)
 * Likes:    112
 * Dislikes: 0
 * Total Accepted:    8.5K
 * Total Submissions: 31.3K
 * Testcase Example:  '[1,2]\n3'
 *
 * 给定一个整数数组 arr 和一个整数 k ，通过重复 k 次来修改数组。
 * 
 * 例如，如果 arr = [1, 2] ， k = 3 ，那么修改后的数组将是 [1, 2, 1, 2, 1, 2] 。
 * 
 * 返回修改后的数组中的最大的子数组之和。注意，子数组长度可以是 0，在这种情况下它的总和也是 0。
 * 
 * 由于 结果可能会很大，需要返回的 10^9 + 7 的 模 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = [1,2], k = 3
 * 输出：9
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = [1,-2,1], k = 5
 * 输出：2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：arr = [-1,-2], k = 7
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 
 * 1 <= arr.length <= 10^5
 * 1 <= k <= 10^5
 * -10^4 <= arr[i] <= 10^4
 * 
 * 
 */

// @lc code=start
#define MOD 10e9+7

class Solution {
public:
    int kConcatenationMaxSum(vector<int>& arr, int k) {
        int dp[2]={0},sz=arr.size();// 用滚动数组dp  dp表示遍历到当前位置时以该位置结尾的子数组的最大和
        
        // 初始化
        dp[0]=0>arr[0]? 0:arr[0];

        for(int i=1;i<k*sz;i++){
            dp[i%2]=max(arr[i%3] , (dp[(i-1)%2]+arr[i%3])%MOD);
        }

        return max(dp[0],dp[1]);
    }
};
// @lc code=end

