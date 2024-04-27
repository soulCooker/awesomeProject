208. 实现 Trie (前缀树)
中等
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
 

示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
 

提示：

1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次

思路：
需要一个初始节点为左Trie树的根
考虑一个单词的插入步骤，从根开始查找子节点列表是否包含值等于剩余单词单词首字母的子节点，如果找到，则将剩余单词的开头index++对子节点递归处理。
停止条件：剩余单词长度为0，或者找不到值相同的子节点。后者说明查找失败，将剩余单词的后缀添加到当前节点上。前者说明前缀存在，更新当前节点为单词的末尾。

查找步骤：和插入步骤类型，区别在停止条件的处理，剩余单词长度为0且当前节点为单词末尾时，返回true，其他停止条件返回false。
前缀查找步骤：和插入步骤类型，区别在停止条件的处理，只在查找失败时返回true。

伪代码：
定义Trie节点
    值
    子节点Map[子节点值]子节点指针
    是否单词末尾

创建初始节点为Trie树的根。

插入步骤：
树root初始为Trie树的根。
初始化待匹配字符串起始下标
遍历当前单词的所有字符
    取当前字符对应在Trie树根的子节点
    子节点存在
        树root，更新为子节点
    子节点不存在
        停止循环

如果待匹配字符串起始下标<字符长度
    遍历待匹配字符串单词
        用字符创建节点
        将添加到root节点的子节点字典中
        更新root为当前节点
更新当前root为单词末尾
返回

查找步骤：
匹配步骤同上

如果待匹配字符串起始下标==字符长度且root为单词末端
    返回true
否则
    返回false

查找前缀步骤：
匹配步骤同上

如果待匹配字符串起始下标==字符长度
    返回true
否则
    返回false

优化点：
只保括小写字符，可以用长度为26的数组来表示子节点列表
插入方法的查找和插入可以在一个循环里解决
