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
}
