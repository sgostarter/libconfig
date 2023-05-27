package libconfig

import (
	"os"
	"path"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func fileExist(path string) bool {
	_, err := os.Lstat(path)

	return !os.IsNotExist(err)
}

func searchConfigFile(fileName string, configPaths []string) string {
	for _, dir := range configPaths {
		if f := path.Join(dir, fileName); fileExist(f) {
			return f
		}
	}

	return ""
}

func LoadOnConfigPath(configName string, configPaths []string, cfg interface{}) (configFileUsed string, err error) {
	_ = envconfig.Process("", cfg)

	defaultConfigPaths, _ := GetDefaultAppConfigRoot()

	allConfigPaths := make([]string, 0, len(configPaths)+len(defaultConfigPaths))
	allConfigPaths = append(allConfigPaths, configPaths...)
	allConfigPaths = append(allConfigPaths, defaultConfigPaths...)

	configFileUsed = searchConfigFile(configName, allConfigPaths)
	if configFileUsed == "" {
		return
	}

	f, err := os.Open(configFileUsed)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()

	err = yaml.NewDecoder(f).Decode(cfg)

	return
}

func Load(configName string, cfg interface{}) (configFileUsed string, err error) {
	return LoadOnConfigPath(configName, []string{
		"./", "./config/", "../", "../config/", "../../", "../../config/",
	}, cfg)
}
