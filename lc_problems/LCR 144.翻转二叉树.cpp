/*
 * @lc app=leetcode.cn id=LCR 144 lang=cpp
 * @lcpr version=30204
 *
 * [LCR 144] 翻转二叉树
 *
 * https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/description/
 *
 * algorithms
 * Easy (79.45%)
 * Likes:    388
 * Dislikes: 0
 * Total Accepted:    393.3K
 * Total Submissions: 495K
 * Testcase Example:  '[5,7,9,8,3,2,4]'
 *
 * 给定一棵二叉树的根节点 root，请左右翻转这棵二叉树，并返回其根节点。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [5,7,9,8,3,2,4]
 * 输出：[5,9,7,4,2,3,8]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/
 *
 *
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
public:
    TreeNode* flipTree(TreeNode* root) {
        if (root != nullptr) {
            flipTree(root->left);
            flipTree(root->right);
            auto tmp = root->left;
            root->left = root->right;
            root->right = tmp;
        }

        return root;
    }
};
// @lc code=end


/*
// @lcpr case=start
// [5,7,9,8,3,2,4]\n
// @lcpr case=end

 */

