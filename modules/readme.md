## goops-kafka相关接口

**clusteradmin接口**

```
    // clusterAdmin是一个接口,拥有如下常用方法

    // 只读操作
    // ListTopics() (map[string]TopicDetail, error)
    topics,_ := clusterAdmin.ListTopics()
    fmt.Println(topics)
    // DescribeTopics(topics []string) (metadata []*TopicMetadata, err error)
    // topicDetail,_ := clusterAdmin.DescribeTopics([]string{"appeventjsonlog_unuseful"})
    // fmt.Println(topicDetail)
	  // ListPartitionReassignments(topics string, partitions []int32) (topicStatus map[string]map[int32]*PartitionReplicaReassignmentsStatus, err error)
    // DescribeConfig(resource ConfigResource) ([]ConfigEntry, error)
    // AlterConfig(resourceType ConfigResourceType, name string, entries map[string]*string, validateOnly bool) error
    // ListConsumerGroups() (map[string]string, error)
    // DescribeConsumerGroups(groups []string) ([]*GroupDescription, error)
    // ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (*OffsetFetchResponse, error)
    // DescribeCluster() (brokers []*Broker, controllerID int32, err error)
    // DescribeLogDirs(brokers []int32) (map[int32][]DescribeLogDirsResponseDirMetadata, error)

    // 写操作
    // CreateTopic(topic string, detail *TopicDetail, validateOnly bool) error
    // DeleteTopic(topic string) error
    // CreatePartitions(topic string, count int32, assignment [][]int32, validateOnly bool) error
    // AlterPartitionReassignments(topic string, assignment [][]int32) error
    // DeleteRecords(topic string, partitionOffsets map[int32]int64) error
    // AlterConfig(resourceType ConfigResourceType, name string, entries map[string]*string, validateOnly bool) error
    // DeleteConsumerGroup(group string) error
```

**Producer接口**

- [X] 同步生产(syncProducer): 从字符串直接生产
- [ ] 同步生产(syncProducer): 从标准输入进行直接生产

**Consumer接口**

- [X] 默认的消息消费，采用`gokafka`消费组
