# trie 树

trie 中文 前缀树 字典树。

字典树其实原理很简单，它可是说是多叉树，因为你并不知道到底有几个叉，它的根节点是没有任何意义的，因为你在根节点下面是各种字符，这些字符是
一个单词的开始，举个例子：

```go
                H           S

          e       i     a     i
        l             y         t
      l
    o      
```
所以这种结构就是trie树，也就是前缀树，比如说我们现在给出前缀是h开头的，那么我们可以吧h开头的所有单词全部匹配过来
这里就是 Hello 和 Hi 这也是前缀树的一大用途

**在trie树中最常见的操作就两种1 将字符加入trie中，2 按照一个前缀去查询一组单词。**

如果在一个字符串中查找到某一个匹配值，其实大可以使用字符串匹配算法，或者使用红黑树，甚至跳表 哈希表都可以，因为前缀树要浪费了很多空间
但是为什么我们要使用前缀树呢？因为给定一个前缀然后找到相关的匹配 这一点非常的棒，前缀树就是搞这个的，
