/*
 * @lc app=leetcode.cn id=1813 lang=golang
 * @lcpr version=30204
 *
 * [1813] 句子相似性 III
 *
 * https://leetcode.cn/problems/sentence-similarity-iii/description/
 *
 * algorithms
 * Medium (41.58%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    26.4K
 * Total Submissions: 63.5K
 * Testcase Example:  '"My name is Haley"\n"My Haley"'
 *
 * 给定两个字符串 sentence1 和 sentence2，每个表示由一些单词组成的一个句子。句子是一系列由 单个 空格分隔的
 * 单词，且开头和结尾没有多余空格。每个单词都只包含大写和小写英文字母。
 *
 * 如果两个句子 s1 和 s2 ，可以通过往其中一个句子插入一个任意的句子（可以是空句子）而得到另一个句子，那么我们称这两个句子是 相似的
 * 。注意，插入的句子必须与现有单词用空白隔开。
 *
 * 比方说，
 *
 *
 * s1 = "Hello Jane" 与 s2 = "Hello my name is Jane"，我们可以往 s1 中 "Hello" 和 "Jane"
 * 之间插入 "my name is" 得到 s2 。
 * s1 = "Frog cool" 与 s2 = "Frogs are cool" 不是相似的，因为尽管往 s1 中插入 "s are"，它没有与
 * "Frog" 用空格隔开。
 *
 *
 * 给你两个句子 sentence1 和 sentence2 ，如果 sentence1 和 sentence2 是 相似 的，请你返回 true
 * ，否则返回 false 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：sentence1 = "My name is Haley", sentence2 = "My Haley"
 *
 * 输出：true
 *
 * 解释：可以往 sentence2 中 "My" 和 "Haley" 之间插入 "name is" ，得到 sentence1 。
 *
 *
 *
 * 示例 2：
 *
 * 输入：sentence1 = "of", sentence2 = "A lot of words"
 *
 * 输出：false
 *
 * 解释：没法往这两个句子中的一个句子只插入一个句子就得到另一个句子。
 *
 *
 *
 * 示例 3：
 *
 * 输入：sentence1 = "Eating right now", sentence2 = "Eating"
 *
 * 输出：true
 *
 * 解释：可以往 sentence2 的结尾插入 "right now" 得到 sentence1 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= sentence1.length, sentence2.length <= 100
 * sentence1 和 sentence2 都只包含大小写英文字母和空格。
 * sentence1 和 sentence2 中的单词都只由单个空格隔开。
 *
 *
 */

// @lcpr-template-start
package algorithmProblems

import "strings"

// @lcpr-template-end
// @lc code=start
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")

	// swap
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	// 分别从首位比较
	i, j := 0, 0
	n, m := len(s1), len(s2)
	for i < n && s1[i] == s2[i] {
		i++
	}
	for j < n && s1[n-1-j] == s2[m-1-j] {
		j++
	}

	// 相同单词数之和
	return i+j >= n
}

// @lc code=end

/*
// @lcpr case=start
// "My name is Haley"\n"My Haley"\n
// @lcpr case=end

// @lcpr case=start
// "of"\n"A lot of words"\n
// @lcpr case=end

// @lcpr case=start
// "Eating right now"\n"Eating"\n
// @lcpr case=end

*/
