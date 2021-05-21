package conf

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var v *viper.Viper

// 初始化应用配置
func Init(appName string) {

	v = viper.New()

	v.SetConfigType("yaml")
	if appName != "" {
		config := "$HOME/.shoppe/" + appName + ".yaml"
		if Exists(AbsPathify(config)) {
			log.Println("load conf: " + AbsPathify(config))
			v.AddConfigPath(config)
		}
	}

	if envConfig := os.Getenv("config"); envConfig != "" {
		v.AddConfigPath(envConfig)
	}

	v.MergeInConfig()

	// initRemote()
}

func initRemote() {
	if v.GetString("remote.provider") == "nacos" {
		port := v.GetUint64("remote.port")
		if port == 0 {
			port = 8080
		}
		// nacos.NewNacosConfigManager(&nacos.Option{
		// 	Url:         v.GetString("remote.url"),
		// 	Port:        port,
		// 	NamespaceId: "public",
		// 	GroupName:   "DEFAULT_GROUP",
		// 	Config:      nacos.Config{DataId: "config_dev"},
		// 	Auth:        nil,
		// })

		v.AddRemoteProvider("", "", "")

	}

}
