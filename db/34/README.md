# 事务

事务就是一组原子性的 SQL 查询，或者说一个独立的工作单元。事务内的语句，要么全部执行成功，要么全部执行失败

所以事务中有 回滚，也就是说如果执行失败那就返回之前执行的那个地方。如果成功就成功就ok了。事务在nosql中应用比较少，
著名的MongoDB甚至都无法直接支持事务，如果是转账这种特别需要事务的地方，使用nosql就不行了，因为一笔钱必须是
付过钱了还是没有付过钱，这个都是需要特别的声明，不能含糊。
