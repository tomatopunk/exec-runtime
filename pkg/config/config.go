package config

import "github.com/kelseyhightower/envconfig"

type Runtime struct {
	Listen string `default:"unix:///var/run/exec-runtime.sock"`
	Image  ImageConfig
}

type ImageConfig struct {
	Mode          string `default:"nexus"`
	Endpoint      string `default:"http://nexus:8080"`
	ImageCacheDir string `default:"/var/exec-runtime"`
}

func LoadConfig() (*Runtime, error) {
	var config Runtime
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
