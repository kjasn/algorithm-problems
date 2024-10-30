/*
 * @lc app=leetcode.cn id=2596 lang=cpp
 * @lcpr version=21913
 *
 * [2596] 检查骑士巡视方案
 */

// @lc code=start
class Solution {
private:
    bool can_arrive(auto& p, int x, int y) { // 看官方题解，也可以直接判断坐标之差的乘积是否等于2 更好
        if(abs(x - p.first) == 2 && abs(y - p.second) == 1) return true;
        else if(abs(x - p.first) == 1 && abs(y - p.second) == 2) return true;
        else return false;

    }

public:
    bool checkValidGrid(vector<vector<int>>& grid) {
        if(grid[0][0]) return false;// 从左上角出发，从下标0出发，否则false！！！

        typedef pair<int, int> PII;
        int n = grid.size();
        unordered_map<int, PII> m;
        for(int i = 0; i < n; i++) {
            for(int j = 0; j < n; j++) {
                m[grid[i][j]] = {i, j};
            }
        }

        int cur_x = 0, cur_y = 0; 
        for(int i = 1; i < n * n; i++) {
            auto p = m[i];
            if(!can_arrive(p,cur_x, cur_y)) return false;
            cur_x = p.first, cur_y = p.second;
        }   
        return true;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,3,6],[5,8,1],[2,7,4]]\n
// @lcpr case=end

 */

