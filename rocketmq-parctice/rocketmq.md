# rocketmq 消费丢失问题

## 案例：
> 以下的消息生产均是集群模式
### case1:同consumergroup, 相同topic , 不tag
c1, c2

现象：
c1:
出现消费了tag c2的情况 
```
c1===topic[wstopic1]===tag:[c2]=======body:[p cnt: 1]
```

c2: 
只能消费到部分消息, 缺失了一半


### case2:同consumergroup, 不同topic
c1, c3
现象:
日志开始出现：
```
time="2023-04-07T17:46:07+08:00" level=warning msg="pull message from broker error" broker="10.0.102.10:10911" underlayError="unknown Response Code: 24, remark: the consumer's subscription not exist\nSee http://rocketmq.apache.org/docs/faq/ for further details."

```
同时c1, c2 只能消费到各自topic下的一半消息


### case3:单consumer, 订阅多次不同tag
c4
先订阅t1, 再订阅t2
现象：
只能消费到t2



## 原因
导致上述问题的主要原因在于 rocketmq 的消息存储与消费策略

包含message queue, 以及客户端消费的负载均衡



## 最佳实践：

消费的订阅关系一致：https://rocketmq.apache.org/zh/docs/4.x/bestPractice/07subscribe

多次订阅问题：https://rocketmq.apache.org/zh/docs/4.x/consumer/02push
一个consumer 对一个topic的多次订阅, 只有最后一个的订阅的回调方法会生效
![image-20230407161317184](https://kingstone95.oss-cn-hangzhou.aliyuncs.com/img/image-20230407161317184.png)



