package conf

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Instance *viper.Viper

const CONFIG_NAME = "app.yaml"

var HOME = filepath.Join(UserHomeDir(), ".shoppe")

// 初始化应用配置
func Init(appName string) {

	Instance = viper.New()

	Instance.SetConfigName(CONFIG_NAME)
	Instance.SetConfigType("yaml")

	if appName != "" {
		config := filepath.Join(HOME, appName)
		if Exists(config) {
			log.Println("load conf: " + config + string(os.PathSeparator) + CONFIG_NAME)
			ct, err := os.ReadFile(filepath.Join(config, CONFIG_NAME))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(ct))

			Instance.MergeConfig(bytes.NewReader(ct))
		}
	}
	envConfig := os.Getenv("config")
	if envConfig != "" {
		Instance.AddConfigPath(envConfig)
	} else {
		Instance.AddConfigPath("./")
	}
	Instance.MergeInConfig()
}
