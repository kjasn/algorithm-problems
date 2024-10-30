/*
 * @lc app=leetcode.cn id=18 lang=cpp
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (36.85%)
 * Likes:    1651
 * Dislikes: 0
 * Total Accepted:    475.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 * 
 * 
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 * 
 * 
 * 你可以按 任意顺序 返回答案 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 
 * 
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        sort(nums.begin(),nums.end());
        int sz=nums.size();
        vector<vector<int>> res;
        // 剪枝
        if(sz<4||(long long)nums[0]+nums[1]+nums[2]+nums[3]>target) return res;

        for(int i=0;i<sz-3;i++){
            if (i > 0 && nums[i] == nums[i - 1]) continue;

            for(int j=i+1;j<sz-2;j++){
                if (j > i+1 && nums[j] == nums[j - 1]) continue;

                int k=j+1,p=sz-1;
                while(k<p){
                    if (k > j+1 && nums[k] == nums[k - 1]) {k++;continue;}
                    long long t = (long long)nums[i]+nums[j]+nums[k]+nums[p];// 记得要显示转型 否则溢出
                    // cout<<t<<endl;
                    if(t==target) {
                        // cout<<"yes "<<i<<j<<k<<p<<endl;
                        res.emplace_back(vector<int>{nums[i],nums[j],nums[k],nums[p]});k++;
                    }
                    else if(t>target) p--;
                    else k++;
                }
            }
        }

        return res;
    }
};
// @lc code=end

