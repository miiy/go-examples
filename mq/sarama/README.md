## 启动关闭

```bash
## 启动 kafka：先启动Zookeeper集群，然后启动Kafka
./zkServer.sh start
./kafka-server-start.sh -daemon ../config/server.properties
# 关闭 kafka：所有kafka节点都关闭后在停止zookeeper
./kafka-server-stop.sh
./zkServer.sh stop
```

## topic

```bash
# 查看所有topic
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --list

# 创建topic
# --topic 操作的 topic 名称
# --partitions 分区数
# --replication-factory 分区副本数
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --create --partitions 1 --replication-factor 1 --topic test
Created topic test.

# 查看某个topic详情
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic test
Topic: test	TopicId: bH61zwaDTrS_56mIgekW4g	PartitionCount: 1	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0

# 修改分区数
# 注意：分区数只能增加不能减少
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic test --partitions 2
# 再次查看topic详情
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic test
Topic: test	TopicId: bH61zwaDTrS_56mIgekW4g	PartitionCount: 2	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	Topic: test	Partition: 1	Leader: 0	Replicas: 0	Isr: 0

# 删除 topic
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic test
```

## producer

```bash
# 发送消息
$ ./kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test
```

## consumer

```bash
# 消费消息
# --from-beginning 从头开始消费
$ ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
```

```bash
go run consumer.go --topic test --offset -2
go run producer/main.go --brokers=localhost:9092
```
