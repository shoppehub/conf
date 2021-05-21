package conf

func Get(key string) interface{} {
	return Instance.Get(key)
}

// 根据 key 获取配置
func GetString(key string) string {
	return Instance.GetString(key)
}

// 根据 key 获取数组
func GetStrings(key string) []string {
	return Instance.GetStringSlice(key)
}

// 根据 key 获取map
func GetMap(key string) map[string]interface{} {
	return Instance.GetStringMap(key)
}

// 根据 key 获取 bool
func GetBool(key string) bool {
	return Instance.GetBool(key)
}

// 根据 key 获取 int64
func GetInt(key string) int64 {
	return Instance.GetInt64(key)
}
