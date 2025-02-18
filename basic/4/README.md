# 树，二叉树

什么是树结构，这里面二叉树是一个代表，

有几个概念 层 高度 深度 层 = 深度+1 深度从上往下数 从0开始 是深度；高度是 从下往上数，并且从0开始。

## 这里列举几种常见的二叉树

- 二叉树

- 二叉查找树

- 平衡二叉查找树

- 红黑树

- 递归树

- 完全平衡二叉树（堆）

- 多叉树

- B+树

- B-

- B*

总之 树的分类非常的多，而且树也是 “非线性数据结构”，还有一个 “非线性数据结构”的是 图。

## 二叉树

所谓二叉树，就是最多有两个叉子的树状结构，通常 二叉树也是最常见的数据结构。

二叉树有 满二叉树：意思就是全部都很规矩，是满的二叉树，完全二叉树：最后的叶子节点全部在左边然后除了叶子节点的其它节点要求是满二叉树。

## 如何储存二叉树

- 指针

- 数组

其实 真正的数据结构 计算机里只有这俩 不连续和连续。

### 指针储存

```go

type Node struct {
  left *Node // 左节点的指针
  right *Node // 右节点的指针
  value interface{}// 节点的value内容
}

type Tree struct {
  root *Node// 根节点的指针。g
  deepth int // tree的高度

}

```
## 数组储存

基于数组的那种二叉树的方式，都是需要计算出来那个叉是多少，我们用到的技术是角标的计算。

我们把根节点设置为 1，那么左边的子节点 就是 2*i = 2 右边的就是 2 * i + 1 = 3，继续
左边的那个子节点的子节点左边也是 2 * i = 4 右边的那个是 2 * i + 1 = 5 右边的那个子节点 3 * 2 = 6 右边 7

所以我们现在得到的节点信息就是 1 2 3 4 5 6 7 所以我们的计算公式就是 `左边：2 * i； 右边： 2 * i + 1；i 是它的父节点 `

所以完全二叉树可以不良妃数组容量的情况下使用数组储存，这也就是为什么 完全二叉树要用数组储存比用指针储存的原因，堆这种数据结构就是
一种完全二叉树，所以，堆通常使用数组来储存。

## 二叉树的遍历

- 前序遍历 先打印节点再打印子左节点再右子节点
- 中序遍历 先打印左子节点在打印节点再打印右子节点
- 后序遍历 先打印左子节点再打印右子节点再打印节点
