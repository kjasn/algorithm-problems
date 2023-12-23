/*
 * @lc app=leetcode.cn id=2454 lang=cpp
 * @lcpr version=30111
 *
 * [2454] 下一个更大元素 IV
 */


// @lcpr-template-start
using namespace std;
#include<bits/stdc++.h>
#define PII pair<int, int>
// @lcpr-template-end
// @lc code=start
class Solution {
private:
    bool cmp(PII a, PII b){
        if(a.first > b.first) return true;
        else if(a.first == b.first){
            return a.second < b.second;
        }
    }

public:
    vector<int> secondGreaterElement(vector<int>& nums) {
        int sz = nums.size();
        vector<int> res(sz, -1);
        if(sz < 2) return res;

        vector<PII> v;
        for(int i = 0; i < sz; i++){
            v.push_back({nums[i], i});
        }
        sort(v.begin(), v.end(), cmp);

        for(int i = 0; i < sz; i++){
            




        return res;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [2,4,0,9,6]\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n
// @lcpr case=end

 */
