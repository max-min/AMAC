## 如何计算时间和空间的复杂度。

公式：`T(n) = O(fn(n))` 平时我们看到的样子是这样的`O(n)`

### 时间复杂度

求解的时候要注意一个原则：去掉低阶，去掉常数，去掉最高阶的常数，例如：`2n^2 + 2n + 1` 它的时间复杂度就是 `O(n^2)`
所以说啊，时间复杂度就是一个估计值，举个例子
```go
for i := 0;i < n ;i++ {
  for j := i;j <n;j++ {

  }
}
```
外层循环那么肯定是N 因为是循环那么那么的中间肯定是`x`，所以说外层是N，内层，因为根据i的变化而变化，比如第一次循环就是n次，第二次就是n-1次，第三次就是n-2次那么结果就是`(n^2-n)*0.5` 所以最后取大头就等于`n^2`

列举几个平时经常看到的时间空间复杂度，按照速度顺序列举，越靠前，时间越少，速度越快。

`1 > logn > n > nlogn > n^2 > n^3 > 2^n > n! > n^n`

第二种计算方式就是嵌套相乘法：

只要是嵌套，那么就使用内部乘以外部就ok了，例如nlogn就是一个logn的算法执行了n倍。

我们来具体解释一下，为什么二分法是logn具体来说是log2n

二分法的最后一步就是只剩下了1，也就是找到了那个数值，加入说将这个过程倒过来呢？
```go
1 2 4 8

//也就是

2^0 2^1 2^2 2^3... 2^x = n
```
也就是说 执行了x次才能等于n ，那么二分法只不过是每次都除以2而已，次数是一样的，所以x = log2n 也就是 logn

nlogn呢？更好理解， 让循环n次来执行二分法，那么这个算法就是nlogn因为是嵌套关系，所以用乘法。

### 最好，最坏，加权平均，均摊时间复杂度
>  加权的意思就是找到某次的发生的概率然后乘以它自己的事件，然后所有事件加在一起就ok了。

举个例子，遍历数组中的数据,

- 最好：这个数据刚好在0 那么你的时间复杂度永远是1了

- 最差：这个数据不存在，那么你的时间复杂度就是n了

- 加权平均：算法是这样的，假如数组中有n个数据，要查找的数据在这个数组里面的概率是0.5，那么在每个地方的概率都是n，那么
乘起来就是1/2n 那么 我们可以这么算

```go
//不在数组和在数组最后一个 都是 1/2n *n
1/2n *1 + 1/2n *2 + ... 1/2n *n + 1/2n *n=  1/2n (1+2+3....n+n) = (1+n)n/4n + 1/2 = n 所以它的加权平均时间复杂度就是n
```

最后一个有意思了，叫做

- 均摊时间复杂度，

有个算法 前n次计算都是O(1) 但是第O(n+1)次是n，那么这个n就分摊到前n次就ok了，也就是 忽略不计。所以这种方法就是一个大概的估计

举个例子:

```go
var sl  = make([]int,10)
var i = 0
func Add(ele int){
if len(slice) < i {
le := len(slice) * 2
slm  := make([]int,le)
for k,v := range sl{
  slm[k] = v
}
sl = slm
}
sl [i] = ele
i++
}
```
首先，它最好的结果就是不超过这个len，那么就是1呗，最坏就是超过了这个len,那么需要循环n次进行copy，所以最好是1，最坏就是n
那么加权平均时间复杂度是多少呢？

`1/n * 1 + 1/n * 1 + ...1/n *n = 1/n *(n+n) = 2`所以加权平均时间复杂度是1

我们想象均摊法，前n-1次都是自由加入所以是1 只有最后一次是n，那么把n分摊到每一次中不就是 每次都是2而已嘛，所以还是1 ，所以 这个算法的均摊时间复杂度就是1

### 空间复杂度

我们常见的空间复杂度就是O(1)、O(n)、O(n2 )

空间复杂度就是在操作某个算法的时候需要额外消耗多少的内存空间。具体的可以举个例子

```go
a := make([]int,10)
```
给定一个切片，让你去这个切片的前9位，那么这个空间复杂度是多少？
就是`O(1)`了，因为你并没有开辟新的空间，如果是递归呢？

其实也很容易理解，空间复杂度就是`O(n)`举个例子

```go
func A(n int){
if n <=0 {
  return
}

  A(n-1)
}
```
这个我们可以看到循环了，知道n=0才可以出栈，所以栈空间而言，向上堆砌了N次,所以它的空间复杂度就是`O(n)` 那么时间复杂度是多少?
这个算法中的时间复杂度也是`O(n)` 因为n不同循环的次数也不同，n等于1循环1次n等于10循环10次

与此同时 `make([]int,n)` 就变成了 o(n)了，假如是 `make(nint,n^2)` 显而易见，那就是o(n^2)了。
