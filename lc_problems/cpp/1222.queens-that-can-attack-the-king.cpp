/*
 * @lc app=leetcode.cn id=1222 lang=cpp
 * @lcpr version=21913
 *
 * [1222] 可以攻击国王的皇后
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> queensAttacktheKing(vector<vector<int>>& queens, vector<int>& king) {
        // 参考灵神题解
        int is_queen[8][8]{};// 数据范围最大 8*8
        for(auto& q : queens) {
            is_queen[q[0]][q[1]] = true;
        }

        int dir[8][2]{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, 
        {1, 1}, {1, -1}, {-1, 1}, {-1, -1}};

        vector<vector<int>> res;
        for(auto& d : dir) {
            int x = king[0] + d[0];
            int y = king[1] + d[1];
            while(x >= 0 && x < 8 && y >= 0 && y < 8) {
                if(is_queen[x][y]) {
                    res.push_back({x, y});break;
                }
                x += d[0];y += d[1];
            }
        }

        return res;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]]\n[0,0]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]]\n[3,3]\n
// @lcpr case=end

// @lcpr case=start
// [[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]]\n[3,4]\n
// @lcpr case=end

 */

