package trie

/*212. 单词搜索 II
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
输出: ["eat","oath"]
说明:
你可以假设所有输入都由小写字母 a-z 组成。
提示
你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？
散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树.

链接：https://leetcode-cn.com/problems/word-search
先用words构造Trie
再遍历看board能否走到叶节点（DFS）
*/
var dx []int = []int{-1, 1, 0, 0}
var dy []int = []int{0, 0, -1, 1}

func findWords(board [][]byte, words []string) []string {
	return nil
}
