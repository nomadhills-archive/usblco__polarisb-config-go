package polarisb_config_go

type ConfigProvider struct {
	PolarisbConfig *PolarisbConfig `json:"polarisbConfig"`
}

type PolarisbConfig struct {
	Interstellar string            `json:"interstellar"`
	Values       map[string]string `json:"values"`
}

func AddPolarisbConfiguration(configFilePath string) *ConfigProvider {
	newConfigProvider := &ConfigProvider{}
	Load(newConfigProvider, configFilePath, "polarisbConfig")
	return newConfigProvider
}
