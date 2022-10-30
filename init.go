package polarisb_config_go

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func Load(config *ConfigProvider, configFilePath string, configName string) {
	if configFilePath == "" {
		// load the config file path from the environment variable
		configFilePath = get("POLARISB_CONFIG_FILE_PATH", "config.json")
	}

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}
}
