## gokafka 

gokafka是一个非JVM的kafka客户端程序，用于多kafka集群管理的命令行客户端，旨在帮助开发者、运维快速进行维护和管理多个kafka集群。

gokafka采用`sarama`库来实现日常kafka集群的基础demo功能。

**kafka-sdk相关**

[confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)

[sarama](https://github.com/Shopify/sarama)

前者是confluent公司开源的kafka-go的sdk，后者是Shopify公司开源的sdk。


**todo list**

- [ ] 指定topic增加分区(TestAddPartitions)
- [ ] 制定topic修改topic级别的参数，比如保留时间等等(TestUpdateTopicConfig()
- [ ] 删除topic(TestDeleteTopic)
- [ ] 获取消费者组的offset信息
- [ ] 

**快速开始**

```
$ git clone https://github.com/goops-top/gokafka.git 
$ cd gokafka 
$ go build -o build/goops-kafka

$ ./build/goops-kafka
goops-kafka: A kafka tools with golang that can operate the kafka for describe,create,update and so on.

 Note: Without the jvm,so you must be specify the [--broker or --cluster and the Value must be in --config entry.]

Usage:
  gokafka [command]

Available Commands:
  consumer    consumer  a topic message data with specified kafka-cluster.
  create      create the kafka topic with some base params in specify kafka-cluster.
  describe    describe the kafka some info (cluster,topic,broker,loginfo)
  help        Help about any command
  init        init the gokafka some default config.
  list        list the kafka some info (cluster,topic,broker)
  producer    producer  a topic message data with specified kafka-cluster.
  version     print version info of goops-kafka tool

Flags:
      --broker string    指定broker地址
      --cluster string   指定集群
      --config string    指定配置文件(default is $HOME/.goops-kafka)
  -h, --help             help for gokafka

Use "gokafka [command] --help" for more information about a command.

# 使用make编译
# 会自动生成mac和linux发行版本的二进制文件
$ make default

```

**简单使用**

```
# 查看工具版本
$ ./build/gokafka.mac version
Version: 0.0.1
GitBranch: master
CommitId: 445f935
Build Date: 2020-06-26T18:49:48+0800
Go Version: go1.14
OS/Arch: darwin/amd64

# 初始化配置文件
# gokafka 采用配置文件的方式快速管理多个kafka集群
$ ./build/gokafka.mac  init
gokafka config init ok.

# 在用户家目录下生成集群配置文件
$ cat ~/.goops-kafka

app: gokafka
spec:
  clusters:
  - name: test-kafka
    version: V2_5_0_0
    brokers:
    - 10.0.0.1:9092
    - 10.0.0.2:9092
    - 10.0.0.3:9092
  - name: dev-kafka
    version: V1_0_0_0
    brokers:
    - 192.168.0.22:9092
    - 192.168.0.23:9092
    - 192.168.0.24:9092

# 也可以使用--config 来指定集群配置文件
$ ./build/gokafka.mac  --config ./kafka-cluster.yaml version
Version: 0.0.1
GitBranch: master
CommitId: 445f935
Build Date: 2020-06-26T18:53:25+0800
Go Version: go1.14
OS/Arch: darwin/amd64 

# 查看配置文件中的集群详情
$ ./build/gokafka.mac  list cluster
cluster:test-kafka version:V2_5_0_0 connector_brokers:[10.0.0.1:9092]
cluster:dev-kafka version:V1_0_0_0 connector_brokers:[192.168.0.22:9092]

./build/gokafka.mac  list cluster --config ./kafka-cluster.yaml
cluster:log-kafka version:V2_5_0_0 connector_brokers:[log-kafka-1.bgbiao.cn:9092]

# 注意: 当使用集群配置文件管理集群时，需要使用--cluster全局参数指定操作的目标集群
# 当然也可以直接指定broker进行操作
$ ./build/gokafka.mac  --cluster dev-kafka describe broker
controller: 3
brokers num: 3
broker list: [2 1 3]
id:2		 broker:192.168.0.22:9092
id:1		 broker:192.168.0.23:9092
id:3		 broker:192.168.0.24:9092

$ ./build/gokafka.mac  --broker 10.0.0.1:9092  describe broker
controller: 3
brokers num: 3
broker list: [2 1 3]
id:2		 broker:192.168.0.22:9092
id:1		 broker:192.168.0.23:9092
id:3		 broker:192.168.0.24:9092


```


**常用功能**

```
# 1.创建topic
# 可以指定分区和副本数
$ ./build/gokafka.mac  --cluster dev-kafka create --topic test-bgbiao-1
true


# 2.查看topic 
$ ./build/gokafka.mac  --cluster dev-kafka list topic --topic-list test-bgbiao-1
Topic:test-bgbiao-1	PartNum:3	Replicas:3	Config:
Topic-Part:test-bgbiao-1-2	ReplicaAssign:[1 3 2]
Topic-Part:test-bgbiao-1-1	ReplicaAssign:[2 1 3]
Topic-Part:test-bgbiao-1-0	ReplicaAssign:[3 2 1]

# 3.获取topic的分区和副本状态
$ ./build/gokafka.mac  --cluster dev-kafka describe  topic --topic-list test-bgbiao-1
Topic-Part:test-bgbiao-1-2	Leader:1	Replicas:[1 3 2]	ISR:[1 3 2]	OfflineRep:[]
Topic-Part:test-bgbiao-1-1	Leader:2	Replicas:[2 1 3]	ISR:[2 1 3]	OfflineRep:[]
Topic-Part:test-bgbiao-1-0	Leader:3	Replicas:[3 2 1]	ISR:[3 2 1]	OfflineRep:[]

# 4.向topic中生产消息
$ ./build/gokafka.mac  --cluster dev-kafka producer --topic test-bgbiao-1 --msg "Hello, BGBiao."
INFO[0000] Produce msg:Hello, BGBiao. to topic:test-bgbiao-1
INFO[0000] topic:test-bgbiao-1 send ok with offset:0

$ ./build/gokafka.mac  --cluster dev-kafka producer --topic test-bgbiao-1 --msg "Nice to meet you."
INFO[0000] Produce msg:Nice to meet you. to topic:test-bgbiao-1
INFO[0000] topic:test-bgbiao-1 send ok with offset:0

# 5.消费消息(gokafka中创建了一个默认的消费组用来预览消息)


# 在终端进行实时消费，ctrl+c 可取消
# 默认会创建一个Gokafka 的消费者组；并且使用从消息的头部开始消费(相当于官方消费者工具的--begin)
# 可手工指定消费者组和消费的位置
## --group 指定消费者组
## --offset 指定消费位置(latest|earliest) 

$ ./build/gokafka.mac  --cluster dev-kafka consumer --topic test-bgbiao-1
INFO[0000] Sarama consumer up and running!...
INFO[0004] part:2 offset:0
msg: Nice to meet you.
INFO[0004] part:1 offset:0
msg: Hello, BGBiao.

# 指定消费者组进行消费，同时指定的消费顺序从最新的消息开始
$ ./build/gokafka.mac  --cluster dev-kafka consumer --topic test-bgbiao-1 --group consumer-group --offset latest




# 6.查看集群的broker列表以及controller
$ ./build/gokafka.mac  --cluster dev-kafka describe broker
controller: 3
brokers num: 3
broker list: [2 1 3]
id:2		 broker:192.168.0.23:9092
id:1		 broker:192.168.0.22:9092
id:3		 broker:192.168.0.24:9092

# 7.查看topic的日志大小
$ ./build/gokafka.mac  --cluster dev-kafka describe loginfo --topic-list test-bgbiao-1
topic:test-bgbiao-1
172.16.32.23:9092
logdir:/soul/data/kafka/kafka-logs
topic-part		log-size(M)		offset-lag
----------		-----------		----------
test-bgbiao-1-0		0		0
test-bgbiao-1-1		0		0
test-bgbiao-1-2		0		0
172.16.32.22:9092
logdir:/soul/data/kafka/kafka-logs
topic-part		log-size(M)		offset-lag
----------		-----------		----------
test-bgbiao-1-0		0		0
test-bgbiao-1-1		0		0
test-bgbiao-1-2		0		0
172.16.32.24:9092
logdir:/soul/data/kafka/kafka-logs
topic-part		log-size(M)		offset-lag
----------		-----------		----------
test-bgbiao-1-0		0		0
test-bgbiao-1-1		0		0
test-bgbiao-1-2		0		0

# 8.列出kafka集群的消费者组
$ ./build/gokafka.mac  --cluster dev-kafka list  consumer-g | grep -i gokafka
GoKafka

# 9.查看消费者组详细信息
# 因为上面我们停止了消费，该消费者组当前是空的
$ ./build/gokafka.mac  --cluster dev-kafka describe consumer-g --consumer-group-list GoKafka
--------------------------------------------------------------------------------------------
consumer-group:GoKafka consumer-state:Empty

# 我们来看一个有实际消费状态的消费者组
# 可以看到一个消费者组下的全部消费者线程以及消费者实例，消费的对应的topic列表
$ ./build/gokafka.mac  --cluster dev-kafka describe consumer-g --consumer-group-list group-sync
--------------------------------------------------------------------------------------------
consumer-group:group-sync consumer-state:Stable
consumer-id					consumer-ip			topic-list
consumer-1-a9437739-e5cb-4b41-a9d3-2640b9878965	/172.16.64.207		[sync-dev]
consumer-1-92eb690b-d327-468b-8990-9c21e1ee405d	/172.16.32.235		[sync-dev]


# 10. 还可以查看某个消费者组消费某个topic的日志详情
$ ./build/gokafka.mac  --cluster dev-kafka describe consumer-group-offset --group group-sync --topic sync-dev
sync-dev
topic-part:sync-dev-0 log-offsize:98
topic-part:sync-dev-1 log-offsize:0
topic-part:sync-dev-2 log-offsize:340
topic-part:sync-dev-3 log-offsize:261






```
