/*
 * @lc app=leetcode.cn id=1921 lang=cpp
 * @lcpr version=21913
 *
 * [1921] 消灭怪物的最大数量
 */

// @lc code=start
class Solution {
public:
    int eliminateMaximum(vector<int>& dist, vector<int>& speed) {
        int sz = dist.size(), res = 0;

        vector<int> t;
        for(int i = 0; i < sz; i++) {
            int tmp = dist[i] / speed[i];
            if(dist[i] % speed[i]) tmp++;// 未除尽y
            t.emplace_back(tmp);
        }

        // 贪心
        // 按到达时间排序 然后遍历
        sort(t.begin(), t.end());
        int timer = 0;
        for(int& x : t) {
            if(timer < x) {
                res++;timer++;
            }else break;
        }

        return res;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [1,3,4]\n[1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,3]\n[1,1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n[5,3,2]\n
// @lcpr case=end

 */

