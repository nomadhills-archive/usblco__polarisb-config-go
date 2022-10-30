package polarisb_config_go

func (configProvider *ConfigProvider) GetValue(key string) (string, bool) {
	value, ok := configProvider.PolarisbConfig.Values[key]
	return value, ok
}
