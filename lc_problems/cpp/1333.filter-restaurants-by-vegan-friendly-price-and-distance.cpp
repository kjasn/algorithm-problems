/*
 * @lc app=leetcode.cn id=1333 lang=cpp
 * @lcpr version=21917
 *
 * [1333] 餐厅过滤器
 */

// @lc code=start
// 简单模拟
#define PII pair<int, int>

class Solution {
public:
    vector<int> filterRestaurants(vector<vector<int>>& restaurants, int veganFriendly, int maxPrice, int maxDistance) {
        vector<PII> v;
        int sz = restaurants.size();
        for(int i = 0; i < sz; i++) {
            if(restaurants[i][2] >= veganFriendly && restaurants[i][3] <= maxPrice && restaurants[i][4] <= maxDistance) {
                v.push_back({restaurants[i][1], restaurants[i][0]});
            }
        }
        sort(v.rbegin(), v.rend());

        vector<int> res;
        for(auto& x : v) res.emplace_back(x.second); 
        return res;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n1\n50\n10\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n0\n50\n10\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n0\n30\n3\n
// @lcpr case=end

 */

