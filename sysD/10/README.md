# 一致性哈希

一致性哈希把节点对应的哈希值变成了一个范围，而不再是离散的。在一致性哈希中，把整个哈希值的范围定义得非常大，然后把这个范围分配给现有的节点

假设哈希值范围是从0到100，四个节点为：[0,25)、[25,50)、[50,75)、[75,100]
第二个节点退出，节点变为：[0,25)、[25,75)、[75,100]，只需要把第二节点接管过来即可，其他节点数据未受影响
在第二和第三节点之间增加节点，节点变为：[0,25)、[25,50)、[50,63)、[63,75)、[75,100]，第三节点部分数据未收影响，另一部分数据给新增节点管理，其他节点数据未受影响

其实核心思想就是为了顾及大局，让相邻的节点处理那个去除的节点，这样其它非相邻的节点就不受到影响。
