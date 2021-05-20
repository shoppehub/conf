package conf

// 根据 key 获取配置
func GetString(key string) string {
	return v.GetString(key)
}

// 根据 key 获取数组
func GetStrings(key string) []string {
	return v.GetStringSlice(key)
}
