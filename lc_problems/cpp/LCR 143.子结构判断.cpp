// @lcpr-before-debug-begin




// @lcpr-before-debug-end

/*
 * @lc app=leetcode.cn id=LCR 143 lang=cpp
 * @lcpr version=30204
 *
 * [LCR 143] 子结构判断
 *
 * https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/description/
 *
 * algorithms
 * Medium (46.35%)
 * Likes:    820
 * Dislikes: 0
 * Total Accepted:    335.9K
 * Total Submissions: 724.7K
 * Testcase Example:  '[1,2,3,4]\n[3]'
 *
 * 给定两棵二叉树 tree1 和 tree2，判断 tree2 是否以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
 * 注意，空树 不会是以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 *
 *
 *
 * 输入：tree1 = [1,7,5], tree2 = [6,1]
 * 输出：false
 * 解释：tree2 与 tree1 的一个子树没有相同的结构和节点值。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：tree1 = [3,6,7,1,8], tree2 = [6,1]
 * 输出：true
 * 解释：tree2 与 tree1 的一个子树拥有相同的结构和节点值。即 6 - > 1。
 *
 *
 *
 * 提示：
 *
 * 0 <= 节点个数 <= 10000
 *
 */


 // @lcpr-template-start
#include<bits/stdc++.h>
using namespace std;
// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
private:
    bool dfs(TreeNode* a, TreeNode* b) {
        if (b == nullptr) {
            return true;
        }
        if (a == nullptr || a->val != b->val) {
            return false;
        }

        return dfs(a->left, b->left) && dfs(a->right, b->right);
    }
public:
    bool isSubStructure(TreeNode* A, TreeNode* B) {
        if (A == nullptr || B == nullptr) {
            return false;
        }
        return dfs(A, B) || isSubStructure(A->left, B) || isSubStructure(A->right, B);
    }
};
// @lc code=end


// @lcpr-div-debug-arg-start
// funName=isSubStructure
// paramTypes= ["number[]","number[]"]
// @lcpr-div-debug-arg-end




/*
// @lcpr case=start
// [1,7,5]\n[6,1]\n
// @lcpr case=end

// @lcpr case=start
// [3,6,7,1,8]\n[6,1]\n
// @lcpr case=end

 */

