# 索引

索引是存储引擎用于快速找到记录的一种数据结构

其实 skip list中对于链表进行改造加入的那些个数据结构就是加入的索引。索引是对于整个数据的一种映射。

索引是在存储引擎层实现的，而不是在服务器层实现的，所以不同存储引擎具有不同的索引类型和实现。

当然B+ tree中也是索引，跟b tree不同b+tree中储存的都是指针，而btree中储存的是真实的值，所以
b+ 是b tree的pro版。
