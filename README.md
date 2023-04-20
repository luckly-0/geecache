这个项目是我跟着极客兔兔7天系列完成的，大家想看更好的可以去极客兔兔7天系列看博客，仅供学习使用
## geecache基本介绍

- GeeCache 是一个基于 Go 语言实现的分布式缓存系统，它是对单机缓存的扩展，可以将缓存数据分布到多台服务器上，提高缓存的容量和性能。
- GeeCache 的核心思想是使用一致性哈希算法（Consistent Hashing）实现缓存数据的分布式存储和访问，将缓存数据按照哈希值分配到不同的节点中，从而实现了数据的负载均衡和容错能力。

- GeeCache 的设计思路简洁、可扩展性强，是学习分布式缓存系统实现的一个很好的案例。

## 本人博客
([CCSU__LRF的博客_CSDN博客-golang项目 - geecache,计算机基础复习,codeforce领域博主](https://blog.csdn.net/csxylrf?spm=1011.2415.3001.5343))欢迎关注
## 笔记 

- [【go项目-geecache】动手写分布式缓存 - day1 - 实现LRU算法](https://blog.csdn.net/csxylrf/article/details/130189041?spm=1001.2014.3001.5501)
- [【go项目-geecache】动手写分布式缓存 - day2 - 单机并发缓存](https://blog.csdn.net/csxylrf/article/details/130210651?spm=1001.2014.3001.5501)
- [【go项目-geecache】动手写分布式缓存 - day3 - HTTP 服务端](https://blog.csdn.net/csxylrf/article/details/130222918?spm=1001.2014.3001.5501)
- [【go项目-geecache】动手写分布式缓存 - day4 - 一致性哈希(hash)](https://blog.csdn.net/csxylrf/article/details/130225636?spm=1001.2014.3001.5501)
- [【go项目-geecache】动手写分布式缓存 - day5 - 分布式节点](https://blog.csdn.net/csxylrf/article/details/130231834?spm=1001.2014.3001.5501)
- [【go项目-geecache】动手写分布式缓存 - day6 - 防止缓存击穿](https://blog.csdn.net/csxylrf/article/details/130244567?spm=1001.2014.3001.5501)
- [【go项目-geecache】动手写分布式缓存 - day7 - 使用 Protobuf 通信](https://blog.csdn.net/csxylrf/article/details/130257794?spm=1001.2014.3001.5501)


