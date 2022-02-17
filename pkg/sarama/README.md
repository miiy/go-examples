topic

```bash
# 创建topic
# --topic 指定topic名，--replication-factory 指定副本数，--partitions 指定分区数
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --create --topic test2 --replication-factor 1 --partitions 1
Created topic test.

# 查看所有topic
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --list
test

# 查看某个topic详情
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic test
Topic: test	PartitionCount: 1	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	
# 发送消息
$ ./kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test

# 消费消息
$ ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning

# 删除 topic
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic test2

# 修改分区数
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic test --partitions 2
$ ./kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic test
Topic: test	PartitionCount: 2	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	Topic: test	Partition: 1	Leader: 0	Replicas: 0	Isr: 0
```

```bash
go run consumer.go --topic test --offset -2
go run producer.go --topic test
```
