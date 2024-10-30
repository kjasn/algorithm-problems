/*
 * @lc app=leetcode.cn id=LCP 50 lang=cpp
 * @lcpr version=21913
 *
 * [LCP 50] 宝石补给
 */

// @lc code=start
class Solution {
public:
    int giveGem(vector<int>& gem, vector<vector<int>>& operations) {
        for(auto& x : operations) {
            int n = gem[x[0]] / 2;
            gem[x[0]] -= n;
            gem[x[1]] += n;
        }

        int mn = 10000, mx = 0;
        for(auto& x : gem) {
            mx = max(mx, x);
            mn = min(mn, x);
        }
        return mx - mn;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [3,1,2]\n[[0,2],[2,1],[2,0]]`>\n
// @lcpr case=end

// @lcpr case=start
// [100,0,50,100]\n[[0,2],[0,1],[3,0],[3,0]]`>\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0]\n[[1,2],[3,1],[1,2]]`>\n
// @lcpr case=end

 */

