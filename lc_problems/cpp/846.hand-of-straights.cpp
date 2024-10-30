// @lcpr-before-debug-begin




// @lcpr-before-debug-end

/*
 * @lc app=leetcode.cn id=846 lang=cpp
 * @lcpr version=21913
 *
 * [846] 一手顺子
 */
using namespace std;
#include <iostream>
#include <map>

#include <vector>
// @lc code=start
class Solution {
public:
    bool isNStraightHand(vector<int>& hand, int groupSize) {
        int sz = hand.size();
        if(sz % groupSize) return false;

        map<int, int> m;// 存储 牌的值 和 数量
        for(int& x : hand) m[x]++;

        // 遍历排序好的牌
        for(auto& x : m) {
            int n = x.second;
            for(int i = 1; i < groupSize; i++) {
                if(m.count(x.first + i) && m[x.first + i] >= n) m[x.first + i] -= n;// 找下一张牌
                else return false;
                if(!m[x.first + i]) m.erase(x.first + i);
            }
            
        }
        return true;
        
    }
};
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,6,2,3,4,7,8]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n4\n
// @lcpr case=end

 */

