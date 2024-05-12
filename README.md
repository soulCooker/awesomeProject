# 社招求职攻略

## 简历准备

框架：联系方式、学历、技能总结、工作履历（总结所有的工作经历，可以快速获取你在不同公司的工作年限、负责过的业务方向等信息）
项目经历编写原则：写项目使用的通用技术栈，得到的可量化成果。项目一般会考察你负责的部分以及在其中的职责，实现难点&&解决方案，如果有高并发、交易场景可以重点介绍，可以从稳定性、可扩展性挖掘项目可以说的点。
针对简历中提到的技能总结和技术栈可能会被问到的问题，开始进行面试准备。

## 算法和数据结构&&代码能力

<https://leetcode.cn/> 热题100
开始写代码前，要构思好解题思路，并推理论证起正确性后，再开始编写。
编写完初版后，检查考虑所有边界情况，确定可以运行且边界情况都处理后，再通知面试官进行检查，尽量避免低级错误降低通过率。
小tip面试前提前用网页编辑器进行编程，适应没有代码自动补全和编译检查的编写环境，养成严谨的编码习惯。

## CS基础知识

### 计算机网络

 tcp/ip
 http
 学习资料：<https://xiaolincoding.com/network/3_tcp/tcp_feature.html#%E6%85%A2%E5%90%AF%E5%8A%A8>

### 操作系统

 进程通信、调度算法、进程资源-段页式内存虚拟空间、套接字
 多线程并发控制 锁类型 死锁
 系统调用、异常和中断

## 设计模式

 掌握概念、有实践经验。
 看书：《head first Java》

## 编程语言

### Java

 JUC 锁的使用和优化
 学习资料：<https://juejin.cn/post/6844903758690779149#heading-2、https://www.baeldung.com/java-concurrent-locks>
 JVM垃圾回收调优
 开发框架：Spring全家桶：Spring框架、Springboot、SpringCloud
 学习资料：Spring可以看官方网站的文档和分享视频

### Go

 slice结构，线程不安全map原理
 sync包 锁、线程安全map原理
 使用channel实现并行式网页爬虫
 学习资料：goland的官方网站，<https://golang.design/go-questions/map/extend/>

## 微服务&&云原生

### docker容器化技术

 容器概念、资源隔离原理
 学习资料：docker官方网站+实践

### RPC框架

 原理：序列化、NIO/Netty

### 服务治理

 服务发现、负载均衡、故障隔离-熔断、降级、重试
 服务可观测性：分布式日志链路、埋点metrics与监控报警

## 消息中间件

### Kafka

 队列分区、磁盘顺序写入优化、零拷贝技术、
 学习资料：<https://it-blog-cn.com/blogs/qmq/source.html#%E5%9B%9B%E3%80%81%E9%9B%B6%E6%8B%B7%E8%B4%9D%E6%8A%80%E6%9C%AF>

### rocketmq

 延迟消息实现 延迟消息队列+定时轮训
 学习资料：<https://juejin.cn/post/6997370258507956232>
 事务实现、半消息队列
 学习资料：<https://www.cnblogs.com/dennyzhangdd/p/14572024.html>

## 数据库

### MySql

 mysql XA：<https://dev.mysql.com/doc/refman/8.0/en/xa.html>

### Redis

#### 基本数据结构

#### 集群模式

主从复制：replicaof命令 sync和psync命令，全量同步和写命令同步
哨兵模式：非cluster的高可用方案，提供主从节点监控、异常通知、自动主从切换、配置管理（主节点服务发现）能力
Cluster模式：水平扩展 哈希算法分片、哈希槽位（16k个）、高可用（主从自动切换、脑裂问题）
学习资料：<https://www.youtube.com/watch?v=3WOfXRjYnGA> <https://redis.io/docs/latest/operate/oss_and_stack/management/sentinel/>

### ES

### Mongodb

## 分布式系统

### 分布式事务

 Seta <https://seata.apache.org/zh-cn/docs/user/quickstart/>

### 分布式共识算法

 Raft <https://raft.github.io/> <https://www.usenix.org/conference/atc14/technical-sessions/presentation/ongaro>