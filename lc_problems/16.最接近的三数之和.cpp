/*
 * @lc app=leetcode.cn id=16 lang=cpp
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (44.85%)
 * Likes:    1455
 * Dislikes: 0
 * Total Accepted:    490.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 * 
 * 返回这三个数的和。
 * 
 * 假定每组输入只存在恰好一个解。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 * 
 * 
 */

// @lc code=start
class Solution { 
public:
    int threeSumClosest(vector<int>& nums, int target) {
        sort(nums.begin(),nums.end());

        int sz=nums.size(),minDif=INT_MAX,s=0,res=0;// 
        for(int i=0;i<sz-2;i++){
            // 剪枝
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }

            int j=i+1,k=sz-1;
            
            // 双指针
            while(j<k){
                s=nums[i]+nums[j]+nums[k];

                if(s==target) return s;

                // 分别用两重条件判断  确保所有情况都考虑到 
                if(s>target){
                    if(s-target<minDif){
                        minDif=s-target;res=s;
                    }
                    k--;
                }
                else{
                    if(target-s<minDif){
                        minDif=target-s;res=s;
                    }
                    j++;
                }
            }
        }
        return res;
    }
};
// @lc code=end

