# 长连接和短链接

长连接 ： 长连接，一个连接上可以发送多个的数据包（所以才会发生了粘包和拆包的情况），在连接保持期间，如果没有数据包发送，需要双方发链路检测包

短链接：短链接，就是发送就关了就可以了，没必要去发送一直保持连接的链路检测包。

1. 短连接： web网站的http服务一般都是用短连接
2. 长连接：适用于点对点之间的频繁通讯,如通信、股票、Feed、直播、共享桌面，特别适合于客户端与服务频繁交互的情况下，如实时共享、多人协作等平台
