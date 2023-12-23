/*
 * @lc app=leetcode.cn id=1423 lang=cpp
 * @lcpr version=30111
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxScore(vector<int>& cardPoints, int k) {
        int allSum = 0;
        for(int& x : cardPoints) allSum += x;
        // cout<<allSum<<endl;

        // 前缀和
        int sz = cardPoints.size();
        vector<int> s(sz + 1);
        for(int i = 0; i < sz; i++) s[i + 1] = s[i] + cardPoints[i];

        int curSum = 0;
        for(int i = 0; i < k + 1; i++) {
            int j = sz - k - 1 + i;
            curSum = min(curSum, s[j + 1] - s[i]);
        }

        return allSum - curSum;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4,5,6,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [9,7,7,9,7,7,9]\n7\n
// @lcpr case=end

// @lcpr case=start
// [1,1000,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,79,80,1,1,1,200,1]\n3\n
// @lcpr case=end

 */

