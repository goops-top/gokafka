## gokafka 

gokafka是一个非JVM的kafka客户端程序，用于多kafka集群管理的命令行客户端，旨在帮助开发者、运维快速进行维护和管理多个kafka集群。

gokafka采用`sarama`库来实现日常kafka集群的基础demo功能。

**kafka-sdk相关**

[confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)

[sarama](https://github.com/Shopify/sarama)

前者是confluent公司开源的kafka-go的sdk，后者是Shopify公司开源的sdk。


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

```

**简单使用**

```
# 查看工具版本
```
