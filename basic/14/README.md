# 贪心算法思想

贪心算法思想我理解就是每次选择都选择一条你认为最好的最优秀的路，所以你很贪心啊。这就是非常非常通俗的贪心算法思想，
因为贪心算法毕竟是一种算法思想而已，所以怎么理解都是可以的。

举个例子，现在有10个苹果大小从大到小排列，有15个孩子，大小从大到小排列，而且他们吃的饭量还不一样，该怎么分才能尽量多满足指标呢？
这个指标就是尽量的多满足孩子，因为孩子多苹果少嘛，所以该怎么完成指标就是这次的算法问题。

这个问题就是一种贪心算法，具体分析，首先孩子多，那么我们只能最多选择10个孩子，因为苹果也是大小不一样的，但是最小的孩子是所有苹果中他最容易
满足的，这就好办了，我们就可以选择最优解，也就是每次都选择最小的孩子然后按照他的最小需求选择苹果，这种算法就是贪心算法--- 每次都选择最优解
然后取满足整个事件。整个事件都是**每次**都选择**最优解** 非常的贪心，就是贪心算法。

区间覆盖问题：有几个区间，他们可能存在重合，现在让你选出尽量多的区间，并且这些区间不重合（端点重合不算）
例如说有 [1,8] [2,3] [7,9] [8,20] [4,6] [11,12]

选出来的区间就应该是 选左边要求比前一次选择的右端大 但是整个区间又要求最够length小，所以右端要求足够小
 [2,3] [4,6]  [7,9] [11,12]  [18, 8]  [8 ,20] 跨度太大，选出来的区间数量过少。

 这种题的关键就是端点不重合，并且要求length足够小（最小），如果不是最小的length那么就不符合规定。

 ## 霍夫曼变化和贪心算法

 霍夫曼编码统计字符出现的频率，然后选择使用二进制给予不同的编码长度，以此来解决比如之前的都使用3位或者几位等等这种现象。
 我们统计出来频率，按照贪心算法，把频率最高的使用尽量少的位数，频率低的使用尽量多的位数，这样有效的减少存储空间。
