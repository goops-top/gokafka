/*
Copyright 2020 BGBiao Ltd. All rights reserved.
@File   : adminApi.go
@Time   : 2021/03/29 15:50:52
@Update : 2021/03/29 15:50:52
@Author : BGBiao
@Version: 1.0
@Contact: weichaungxxb@qq.com
@Desc   : None
*/
package controller

import (
	"gokafka/api"

	"context"

	"github.com/goops-top/utils/kafka"
)

// 构造 kafka 集群的元数据信息
func NewClusterContext(clusterInfo api.ClusterInfo) context.Context {
	var ctx context.Context
	// 集群的必须填参数
	ctx = context.WithValue(context.Background(), "name", clusterInfo.Name)
	ctx = context.WithValue(ctx, "version", clusterInfo.Version)
	ctx = context.WithValue(ctx, "brokers", clusterInfo.Brokers)
	ctx = context.WithValue(ctx, "sasl", clusterInfo.Sasl)

	if clusterInfo.Sasl {
		ctx = context.WithValue(ctx, "saslType", clusterInfo.SaslType)
		ctx = context.WithValue(ctx, "saslUser", clusterInfo.SaslUser)
		ctx = context.WithValue(ctx, "saslPassword", clusterInfo.SaslPassword)
	}

	return ctx
}

type ClusterApi struct {
	AdminApi    *kafka.AdminApi
	ProducerApi *kafka.ProducerApi
	ConsumerApi *kafka.Api
}

// 创建一个多功能的api 接口
// apiType: admin,producer,consumer
func NewClusterApi(ctx context.Context, apiType string) (ClusterApi, error) {
	// clusterName := ctx.Value("name").(string)
	// clusterVersion := ctx.Value("name").(string)
	isSasl := ctx.Value("sasl").(bool)
	brokers := ctx.Value("brokers").([]string)

	var kfkAdmin *kafka.AdminApi
	var kfkProducer *kafka.ProducerApi
	var kfkConsumer *kafka.Api

	switch apiType {
	case "admin":
		if isSasl {
			// 没有返回值的函数，要么生要么直接死
			kfkAdmin = kafka.NewClusterAdminWithSASLPlainText(brokers, ctx.Value("saslUser").(string), ctx.Value("saslPassword").(string))
		} else {
			kfkAdmin = kafka.NewClusterAdmin(brokers)
		}
		return ClusterApi{AdminApi: kfkAdmin}, nil

	case "producer":
		if isSasl {
			kfkProducer = kafka.NewProducerApiWithSASLPlainText(brokers, ctx.Value("saslUser").(string), ctx.Value("saslPassword").(string))
		} else {
			kfkProducer = kafka.NewProducerApi(brokers)
		}
	// consumer 需要使用consumer group 和 offset信息才可以创建
	case "consumer":
		if isSasl {
			kfkConsumer = kafka.NewConsumerApiWithSASLPlainText(brokers, "GoKafka", "", ctx.Value("saslUser").(string), ctx.Value("saslPassword").(string))
		} else {
			kfkConsumer = kafka.NewConsumerApi(brokers, "GoKafka", "")
		}
	}

	return ClusterApi{AdminApi: kfkAdmin, ProducerApi: kfkProducer, ConsumerApi: kfkConsumer}, nil
}
