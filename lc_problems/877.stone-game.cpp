/*
 * @lc app=leetcode.cn id=877 lang=cpp
 * @lcpr version=21913
 *
 * [877] 石子游戏
 */
using namespace std;
#include <iostream>

#include <vector>
// @lc code=start
class Solution {
public:
    bool stoneGame(vector<int>& piles) {
        int sz = piles.size();
        int round = 1;// 游戏进行的轮数
        int scA = 0, scB = 0;// 分数
        for(int i = 0, j = sz - 1; i <= j;){
            int front = 0, back = 0;// 表示从头或尾取，可以 多赢得 的分
            if(i != sz - 1) front = piles[i] - piles[i + 1];
            if(j != 0) back = piles[j] - piles[j - 1];

            int tmpSc = 0;
            if(back > front) {
                tmpSc = piles[j];j--;
            }else{
                tmpSc = piles[i];i++;
            }
            round % 2? scA += tmpSc : scB += tmpSc;
        }

        return scA > scB;
    }
};
// @lc code=end



/*
// @lcpr case=start
// [5,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,7,2,3]\n
// @lcpr case=end

 */

