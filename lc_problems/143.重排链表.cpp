// @before-stub-for-debug-begin
#include <vector>
#include <string>
#include "commoncppproblem143.h"

using namespace std;
// @before-stub-for-debug-end

/*
 * @lc app=leetcode.cn id=143 lang=cpp
 *
 * [143] 重排链表
 *
 * https://leetcode.cn/problems/reorder-list/description/
 *
 * algorithms
 * Medium (64.93%)
 * Likes:    1255
 * Dislikes: 0
 * Total Accepted:    261.7K
 * Total Submissions: 402K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
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
#include <unordered_map>

class Solution
{
public:
    void reorderList(ListNode *head)
    {
        if (!head)
            return;

        unordered_map<int, ListNode *> ump;
        int i = 1;
        ListNode *tr = head->next;
        while (tr)
        {
            ump[i++] = tr;
            tr = tr->next;
        }

        i = ump.size(), tr = head;
        int n = ump.size() % 2 ? ump.size() / 2 + 1 : ump.size() / 2;

        while (i > n)
        {
            ump[i]->next = tr->next;
            tr->next = ump[i];
            i--;
            tr = tr->next;
            if (tr == nullptr)
                break;
            tr = tr->next;
        }
    }
};
// @lc code=end
