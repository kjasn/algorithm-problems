/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并 K 个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (57.84%)
 * Likes:    2566
 * Dislikes: 0
 * Total Accepted:    683.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 * 
 * 
 * 示例 2：
 * 
 * 输入：lists = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 输入：lists = [[]]
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if(!lists.size()) return nullptr;
        while(lists.size() != 1) {
            // 两两合并
            ListNode *h = new ListNode(-1);
            ListNode *cp = h;
            auto l1 = lists[0], l2 = lists[1];
            while(l1 || l2) {   // 换个方法些，不用判断哪个遍历完
                int v1 = l1 ? l1 -> val : 1e5;
                int v2 = l2 ? l2 -> val : 1e5;
                
                if (v1 < v2) {
                    h -> next = l1;
                    if(l1) l1 = l1 -> next;
                }else {
                    h -> next = l2;
                    if(l2) l2 = l2 -> next;
                }
                h = h -> next;
            }

            lists.erase(lists.begin(), lists.begin() + 2);
            lists.emplace_back(cp -> next);
        }

        return lists[0];
    }
};
// @lc code=end

