## shortLink
一个基于`redis`做缓存的短地址生成与访问服务

### v0.1
第一版实现的功能
1. `url -> shortLink`的转换
2. `shortLink`重定向跳转到`url`
3. 短链的过期时间设置与续租
4. 利用redis的`INCR`原子增和`SETNX`预防在分布式多机情况下`url`可能被多次生成短链的问题(主要原因是redis没有强事务)