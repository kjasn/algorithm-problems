/*
 * @lc app=leetcode.cn id=1962 lang=cpp
 * @lcpr version=30112
 *
 * [1962] 移除石子使总数最小
 */


// @lcpr-template-start
using namespace std;
#include<bits/stdc++.h>
// @lcpr-template-end
// @lc code=start
class Solution {
public:
    int minStoneSum(vector<int>& piles, int k) {
        int sum = accumulate(piles.begin(),piles.end(),0);
        priority_queue<int> pq;
        for(int i = 0; i < piles.size(); i++){
            pq.push(piles[i]);
        }

        while(k--){
            int tp = pq.top();
            if(!tp) break;
            
            sum -= tp / 2;
            tp = (tp + 1) / 2;
            
            pq.pop();
            pq.push(tp);
        }
        return sum;
    }
};
// @lc code=end




// @lcpr case=start
// [5,4,9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [4,3,6,7]\n3\n
// @lcpr case=end


