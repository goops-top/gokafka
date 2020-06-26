/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:main.go
*Author:Xuebiao Xu
*Date:2020年05月07日
*Description:
*
================================================================*/
package modules

import (
	"fmt"

	"gokafka/api"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type AdminApi struct {
	Admin sarama.ClusterAdmin
}

// init the sdk client config
func newConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = Clientid
	config.ChannelBufferSize = 256
	//config.Version = sarama.V2_5_0_0
	// kafka管理接口向后兼容，因此可以适当使用老接口进行管理，否则对于不同版本的集群可能造成兼容性问题
	// 其实也可以将版本开放出去进行兼容

	version, versionErr := sarama.ParseKafkaVersion("1.0.0")
	if versionErr != nil {
		log.Fatalf("Error parsing Kafka version: %v", versionErr)
	}
	config.Version = version

	if confErr := config.Validate(); confErr != nil {
		fmt.Println("kafka config 解析错误，check please.")
	}
	return config
}

// init a clusteradmin api
func NewClusterAdmin(brokerList []string) *AdminApi {
	config := newConfig()
	// https://pkg.go.dev/github.com/Shopify/sarama?tab=doc#ClusterAdmin
	admin, adminErr := sarama.NewClusterAdmin(brokerList, config)
	if adminErr != nil {
		fmt.Printf("%v\n", adminErr)
		panic("kafka connection failed ")
	}

	return &AdminApi{Admin: admin}
}

// Close the adminApi
func (adminApi *AdminApi) Close() {
	adminApi.Admin.Close()
}

/*
type TopicDetail struct {
	NumPartitions     int32
	ReplicationFactor int16
	ReplicaAssignment map[int32][]int32
	ConfigEntries     map[string]*string
}
*/

// Create a topic with default (partition:3,replicationFactor:3)
func (adminApi *AdminApi) CreateTopic(name string) (bool, error) {
	var topicDetail sarama.TopicDetail
	topicDetail.NumPartitions = 3
	topicDetail.ReplicationFactor = 3
	// CreateTopic(topic string, detail *TopicDetail, validateOnly bool) error
	err := adminApi.Admin.CreateTopic(name, &topicDetail, false)
	if err != nil {
		return false, err
	}
	return true, err
}

// Create a custom topic with the partNum,replicaFactor and one config
func (adminApi *AdminApi) CreateCustomTopic(name string, partNum int32, replicaFactor int16, config map[string]string) (bool, error) {
	var topicDetail sarama.TopicDetail
	topicDetail.NumPartitions = partNum
	topicDetail.ReplicationFactor = replicaFactor
	for key, value := range config {
		topicDetail.ConfigEntries = map[string]*string{key: &value}
	}
	err := adminApi.Admin.CreateTopic(name, &topicDetail, false)
	if err != nil {
		return false, err
	}
	return true, err

}

// Delete a topic
func (adminApi *AdminApi) DeleteTopic(name string) (bool, error) {
	err := adminApi.Admin.DeleteTopic(name)
	if err != nil {
		return false, err
	}
	return true, err
}

// ListTopics() (map[string]TopicDetail, error)
func (adminApi *AdminApi) ListTopic() (map[string]sarama.TopicDetail, error) {
	return adminApi.Admin.ListTopics()
}

// topic info struct (map[string]sarama.TopicDetail)
type TopicInfo struct {
	// topic base info
	Name              string             `json:"name"`
	PartitionNum      int32              `json:"partitionNum"`
	Replication       int16              `json:"replication"`
	ReplicaAssignment map[int32][]int32  `json:"replicaAssignment"`
	ConfigEntries     map[string]*string `json:"configEntries"`
	// topic metadata state info
	PartId              int32   `json:"partId"`
	PartLeader          int32   `json:"partLeader"`
	PartReplicas        []int32 `json:"partReplicas"`
	PartIsr             []int32 `json:"partIsr"`
	PartOfflineReplicas []int32 `json:"partOfflineReplicas"`
}

// ListTopicInfo by topic list
// while the topics is empty,will list all topic
func (adminApi *AdminApi) ListTopicsInfo(topiclist []string) ([]TopicInfo, error) {
	var topicsInfos []TopicInfo
	topics, topicErr := adminApi.ListTopic()
	if topicErr != nil {
		return topicsInfos, topicErr
	}
	if len(topiclist) != 0 {
		for _, topicName := range topiclist {
			for topic, topicInfos := range topics {
				if topic == topicName {
					topicInfo := &TopicInfo{
						Name:              topic,
						PartitionNum:      topicInfos.NumPartitions,
						Replication:       topicInfos.ReplicationFactor,
						ReplicaAssignment: topicInfos.ReplicaAssignment,
						ConfigEntries:     topicInfos.ConfigEntries,
					}
					topicsInfos = append(topicsInfos, *topicInfo)
				}

			}
		}
	} else {
		for topic, topicInfos := range topics {
			topicInfo := &TopicInfo{
				Name:              topic,
				PartitionNum:      topicInfos.NumPartitions,
				Replication:       topicInfos.ReplicationFactor,
				ReplicaAssignment: topicInfos.ReplicaAssignment,
				ConfigEntries:     topicInfos.ConfigEntries,
			}
			topicsInfos = append(topicsInfos, *topicInfo)
		}
	}
	return topicsInfos, topicErr
}

// Describe some topics
// DescribeTopics(topics []string) (metadata []*TopicMetadata, err error)
func (adminApi *AdminApi) describeTopics(topics []string) ([]*sarama.TopicMetadata, error) {
	return adminApi.Admin.DescribeTopics(topics)
}

// describe topic metainfo
func (adminApi *AdminApi) DescribeTopics(topiclist []string) ([]TopicInfo, error) {
	var topicsInfos []TopicInfo
	topicMetadataList, topicMetaErr := adminApi.Admin.DescribeTopics(topiclist)
	if topicMetaErr != nil {
		return topicsInfos, topicMetaErr
	}
	for _, topicMetaData := range topicMetadataList {
		for _, partinfo := range topicMetaData.Partitions {
			topicInfo := &TopicInfo{
				Name:                topicMetaData.Name,
				PartId:              partinfo.ID,
				PartLeader:          partinfo.Leader,
				PartReplicas:        partinfo.Replicas,
				PartIsr:             partinfo.Isr,
				PartOfflineReplicas: partinfo.OfflineReplicas,
			}
			topicsInfos = append(topicsInfos, *topicInfo)
		}
	}
	return topicsInfos, topicMetaErr
}

// List the consumer group avaliable
// map[string]string 中key为消费者组，value为消费者状态，consumer表示正在消费
func (adminApi *AdminApi) ListConsumerGroups() (map[string]string, error) {
	return adminApi.Admin.ListConsumerGroups()
}

// Describe the given consumer groups
func (adminApi *AdminApi) DescribeConsumerGroups(groups []string) ([]*sarama.GroupDescription, error) {
	return adminApi.Admin.DescribeConsumerGroups(groups)
}

// List the consumer group offset available
func (adminApi *AdminApi) ListConsumerGroupOffsets(group string, topicPart map[string][]int32) (*sarama.OffsetFetchResponse, error) {
	return adminApi.Admin.ListConsumerGroupOffsets(group, topicPart)
}

// Describe cluster
// return : broker_list,controllerId,error
func (adminApi *AdminApi) DescribeCluster() ([]*sarama.Broker, int32, error) {
	return adminApi.Admin.DescribeCluster()
}

// get broker id list
// 返回controllerid
// brokerid 列表
// brokerid和broker地址对应关系
// 注意:如果结构体指针方法中使用var定义了变量之后，在函数返回中就不需要写名称了(返回数据中的名称其实就是相当于var name type).
// 如下方法可以将返回中的名称去掉，把方法中的var注释去掉也可以
func (adminApi *AdminApi) GetBrokerIdList() (controllerId int32, brokerIds []int32, brokerInfo []api.BrokerAddr) {
	brokers, controllerId, clusterErr := adminApi.DescribeCluster()
	if clusterErr != nil {
		fmt.Printf("failed to get the kafka cluster broker id list with err:%v\n", clusterErr)
		panic(clusterErr)
	}
	//var brokerInfo []api.BrokerAddr
	//var brokerIds []int32
	for _, broker := range brokers {
		brokerinfo := &api.BrokerAddr{
			BrokerId: broker.ID(),
			BrokerIp: broker.Addr(),
		}

		brokerInfo = append(brokerInfo, *brokerinfo)
		brokerIds = append(brokerIds, broker.ID())
	}

	return controllerId, brokerIds, brokerInfo

}

// Describe logdir
func (adminApi *AdminApi) DescribeLogDirs(brokers []int32) (map[int32][]sarama.DescribeLogDirsResponseDirMetadata, error) {
	return adminApi.Admin.DescribeLogDirs(brokers)
}

// modify the sarama.DescribeLogDirsResponseDirMetadata
type LogDirResponseData struct {
	sarama.DescribeLogDirsResponseDirMetadata
}

type LogdirTopicLoginfo struct {
	Path    string         `json:"path"`
	LogInfo []TopicLogInfo `json:"logInfo"`
}

//the topic partition loginfo on someone broker with a logdir.
type TopicLogInfo struct {
	TopicPart string `json:"topicPart"`
	LogSize   int64  `json:"logSize"`
	OffsetLag int64  `json:"offsetLag"`
}

// get the topic-partation size from a logdir
// logdir.Path
// https://pkg.go.dev/github.com/Shopify/sarama?tab=doc#DescribeLogDirsResponseDirMetadata
func (logdir LogDirResponseData) GetLogSize(topic string) LogdirTopicLoginfo {
	logInfos := &LogdirTopicLoginfo{}
	logInfos.Path = logdir.Path
	var topicLogInfos []TopicLogInfo
	if topic == "" {
		for _, topicsPartSize := range logdir.Topics {
			for _, logSize := range topicsPartSize.Partitions {
				topicLog := &TopicLogInfo{
					TopicPart: fmt.Sprintf("%v-%v", topicsPartSize.Topic, logSize.PartitionID),
					LogSize:   logSize.Size / 1024 / 1024,
					OffsetLag: logSize.OffsetLag,
				}
				topicLogInfos = append(topicLogInfos, *topicLog)
			}
		}
	} else {
		for _, topicsPartSize := range logdir.Topics {
			if topic == topicsPartSize.Topic {
				for _, logSize := range topicsPartSize.Partitions {
					topicLog := &TopicLogInfo{
						TopicPart: fmt.Sprintf("%v-%v", topicsPartSize.Topic, logSize.PartitionID),
						LogSize:   logSize.Size / 1024 / 1024,
						OffsetLag: logSize.OffsetLag,
					}
					topicLogInfos = append(topicLogInfos, *topicLog)
				}
			}
		}
	}
	logInfos.LogInfo = topicLogInfos
	return *logInfos
}

// Getloginfo from the broker id
func (adminApi *AdminApi) GetLogFromBrokers(ids []int32) {
	// 需要校验用户输入的id的合法性吗
	//_,brokerIds,_ := adminApi.GetBrokerIdList()
	brokerLogs, brokerLogErr := adminApi.DescribeLogDirs(ids)
	if brokerLogErr != nil {
		fmt.Printf("err:%v\n", brokerLogErr)
	}
	for id, logdirs := range brokerLogs {
		// 其实需要id和broker动态转换的接口
		fmt.Printf("broker_id:%v\n", id)
		// 开始遍历节点下每个日志目录下的topic分区的数据量
		// logdirs是多个挂载目录下的数据
		for _, logdir := range logdirs {
			topicLog := &LogDirResponseData{logdir}
			fmt.Println(topicLog.GetLogSize(""))
		}
	}
}

type BrokerLogInfo struct {
	BrokerId int32                `json:"brokerId"`
	BrokerIp string               `json:"brokerIp"`
	LogDatas []LogdirTopicLoginfo `json:"logDatas"`
}

// Getloginfo from the a topic
func (adminApi *AdminApi) GetLogFromTopic(topic string) []BrokerLogInfo {
	var brokersLogInfo []BrokerLogInfo

	_, brokerIds, brokerInfos := adminApi.GetBrokerIdList()

	brokerLogs, brokerLogErr := adminApi.DescribeLogDirs(brokerIds)
	if brokerLogErr != nil {
		fmt.Printf("err:%v\n", brokerLogErr)
	}

	for id, logdirs := range brokerLogs {
		brokerLogInfo := &BrokerLogInfo{}
		// 其实需要id和broker动态转换的接口
		var brokerIp string
		for _, broker := range brokerInfos {
			if id == broker.BrokerId {
				brokerIp = broker.BrokerIp
			}
		}

		// 开始遍历节点下每个日志目录下的topic分区的数据量
		// logdirs是多个挂载目录下的数据
		var logdirsData []LogdirTopicLoginfo
		for _, logdir := range logdirs {
			topicLog := &LogDirResponseData{logdir}
			logdata := topicLog.GetLogSize(topic)
			// 去除logdir下没有的数据的目录
			if len(logdata.LogInfo) != 0 {
				logdirsData = append(logdirsData, topicLog.GetLogSize(topic))
			}
		}
		// 去除broker上没有日志的数据
		if len(logdirsData) != 0 {
			brokerLogInfo.LogDatas = logdirsData
			brokerLogInfo.BrokerId = id
			brokerLogInfo.BrokerIp = brokerIp

			brokersLogInfo = append(brokersLogInfo, *brokerLogInfo)
		}

	}

	return brokersLogInfo

}

type TopicsBrokerLogInfo struct {
	Name    string `json:"name"`
	LogData []BrokerLogInfo
}

// Getloginfo from topic list
func (adminApi *AdminApi) GetLogFromTopics(topics []string) []TopicsBrokerLogInfo {
	var topicBrokersLogInfo []TopicsBrokerLogInfo
	for _, topic := range topics {
		brokerLogInfo := &TopicsBrokerLogInfo{
			Name:    topic,
			LogData: adminApi.GetLogFromTopic(topic),
		}
		topicBrokersLogInfo = append(topicBrokersLogInfo, *brokerLogInfo)
	}
	return topicBrokersLogInfo
}

// List ConsumerGroups
// 查看存活的在消费的消费者列表
func (adminApi *AdminApi) ListConsumerGroup() ([]string, error) {
	var consumerList []string

	consumerGroupList, consumerErr := adminApi.ListConsumerGroups()
	if consumerErr != nil {
		return consumerList, consumerErr
	}

	for consumer, consumerState := range consumerGroupList {
		if consumerState == "consumer" {
			consumerList = append(consumerList, consumer)
		}
	}

	return consumerList, consumerErr
}

// ConsumerGroup and members info
type ConsumerGroupMember struct {
	GroupID    string `json:"groupID"`
	State      string `json:"state"`
	ClientInfo []ConsumerMemberInfo
}
type ConsumerMemberInfo struct {
	ClientID  string   `json:"clientID"`
	ClientIP  string   `json:"clientIP"`
	TopicList []string `json:"topicList"`
}

// Describe the ConsumerGroups
func (adminApi *AdminApi) DescribeConsumerGroup(groups []string) (consumergroupmembers []*ConsumerGroupMember, err error) {

	groupInfo, groupErr := adminApi.DescribeConsumerGroups(groups)

	if groupErr != nil {
		return consumergroupmembers, groupErr
	}

	for _, v := range groupInfo {
		consumerGroupMember := &ConsumerGroupMember{}

		consumerGroupMember.GroupID = (*v).GroupId
		consumerGroupMember.State = (*v).State

		var consumerMemberInfos []ConsumerMemberInfo
		for name, members := range (*v).Members {
			var consumerMemberInfo ConsumerMemberInfo
			consumerMemberInfo.ClientID = name
			consumerMemberInfo.ClientIP = members.ClientHost

			// fmt.Printf("%v-%v\n",name,members.ClientHost)

			consumerMemberMetadata, _ := members.GetMemberMetadata()

			consumerMemberInfo.TopicList = consumerMemberMetadata.Topics

			consumerMemberInfos = append(consumerMemberInfos, consumerMemberInfo)
		}

		consumerGroupMember.ClientInfo = consumerMemberInfos

		// fmt.Println(consumerGroupMember)
		consumergroupmembers = append(consumergroupmembers, consumerGroupMember)
	}

	return consumergroupmembers, nil
}

// List the ConsumerGroupOffSets
//

func (adminApi *AdminApi) ListConsumerGroupOffSet(group, topic string) ([]TopicPartOffSet, error) {
	topics, topicErr := adminApi.ListTopicsInfo([]string{topic})
	if topicErr != nil {
		panic(topicErr)
	}

	partNum := topics[0].PartitionNum
	parts := []int32{}
	var i int32
	for i = 0; i < partNum; i++ {
		parts = append(parts, i)
	}

	topicPart := map[string][]int32{topic: parts}

	return adminApi.ListConsumerGroupOffSets(group, topicPart)
}

// topic-part 的offset信息
type TopicPartOffSet struct {
	TopicPart string
	OffSet    int64
}

// 获取topic的分区列表
func (adminApi *AdminApi) ListConsumerGroupOffSets(group string, topicPartitions map[string][]int32) ([]TopicPartOffSet, error) {

	var topicPartsOff []TopicPartOffSet
	offsetResp, offsetErr := adminApi.Admin.ListConsumerGroupOffsets(group, topicPartitions)
	if offsetErr != nil {
		return topicPartsOff, offsetErr
	}

	for topic, partOffSet := range offsetResp.Blocks {
		fmt.Println(topic)
		for part, offset := range partOffSet {
			// 将offset和logsize放在一起即可查看到堆积量
			topicPartsOff = append(topicPartsOff, TopicPartOffSet{TopicPart: fmt.Sprintf("%v-%v", topic, part), OffSet: offset.Offset})
			// fmt.Printf("topic-part:%v\t offset:%v\n", fmt.Sprintf("%v-%v", topic, part), offset.Offset)
		}
	}
	return topicPartsOff, offsetErr
}
