package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func WatchConfig() {

	Instance.WatchConfig()
	Instance.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

}
