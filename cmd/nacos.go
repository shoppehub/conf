package main

import (
	"log"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Option struct {
	Endpoint    string   `json:"endpoint"`
	NamespaceId string   `json:"namespaceId"`
	AccessKey   string   `json:"accessKey"`
	SecretKey   string   `json:"secretKey"`
	GroupName   string   `json:"groupName"`
	DataIds     []string `json:"dataIds"`
}

type Config struct {
	DataId string `json:"dataId"`
}

func InitNacos(option *Option) {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	clientConfig := constant.ClientConfig{
		//
		Endpoint:       option.Endpoint,
		NamespaceId:    option.NamespaceId,
		AccessKey:      option.AccessKey,
		SecretKey:      option.SecretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		log.Println(err)
		return
	}

	if len(option.DataIds) == 0 {
		return
	}

	for _, dataId := range option.DataIds {

		// 监听配置
		configClient.ListenConfig(vo.ConfigParam{
			DataId: dataId,
			Group:  option.GroupName,
			OnChange: func(namespace, group, dataId, data string) {
				v.MergeConfig(strings.NewReader(data))
			},
		})

		// 获取配置
		content, err := configClient.GetConfig(vo.ConfigParam{
			DataId: dataId,
			Group:  option.GroupName})

		if err != nil {
			log.Println(err)
			continue
		}

		v.MergeConfig(strings.NewReader(content))

	}

}
