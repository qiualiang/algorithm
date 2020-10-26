package trie

/*208. 实现 Trie (前缀树)
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:
你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
*/

type TrieNode struct {
	val      byte
	isWord   bool
	children []*TrieNode
}

func NewNode(c byte) *TrieNode {
	node := new(TrieNode)
	node.val = c
	node.children = make([]*TrieNode, 26)
	node.isWord = false
	return node
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var c byte = ' '
	node := NewNode(c)
	return Trie{root: node}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	ws := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if ws.children[c-'a'] == nil {
			ws.children[c-'a'] = NewNode(c)
		}
		ws = ws.children[c-'a']
	}
	ws.isWord = true

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ws := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if ws.children[c-'a'] == nil {
			return false
		}
		ws = ws.children[c-'a']
	}
	return ws.isWord

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	ws := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if ws.children[c-'a'] == nil {
			return false
		} else {
			ws = ws.children[c-'a']
		}
	}
	return true

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
