/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:k-broker.go
*Author:Xuebiao Xu
*Date:2020年05月08日
*Description:
*
================================================================*/
package api

type BrokerAddr struct {
	BrokerId int32  `json:"brokerId"`
	BrokerIp string `json:"brokerIp"`
}
