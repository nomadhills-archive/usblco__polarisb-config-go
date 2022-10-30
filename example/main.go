package main

import polarisb_config_go "polarisb-config-go"

func main() {
	config := polarisb_config_go.AddPolarisbConfiguration("example/config.json")
	println(config.PolarisbConfig.Interstellar)
	println(config.GetValue("example_key1"))
}
