# 僵尸进程和孤儿进程

僵尸进程：也就是说父进程需要调用wait或者waitpid的时候才能让子进程退出。但是父进程忘了，所以就没有退出了，导致流离失所僵尸了

孤儿进程：也就是说父进程退出了，然后子进程还在运行，它就是孤儿进程

其实看起来僵尸进程和孤儿进程很像，都是父进程不管了，其实不是，僵尸是因为父进程对它的一个抛弃，父进程还在啊，而孤儿进程就是父进程死了，
无奈它成孤儿了，所以如果想回收僵尸进程，就是杀掉父进程，僵尸进程变成孤儿进程，然后让init进程收养这个孤儿进程即可。所以说
上面那个孤儿进程的也是一样的，只要是孤儿进程一律由init进程收养。
