package libconfig

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

func LoadOnConfigPath(configName string, configPaths []string, cfg interface{}) (configFileUsed string, err error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.SetDefault("chart::values", map[string]interface{}{
		"ingress": map[string]interface{}{
			"annotations": map[string]interface{}{
				"traefik.frontend.rule.type":                 "PathPrefix",
				"traefik.ingress.kubernetes.io/ssl-redirect": "true",
			},
		},
	})

	sp := strings.LastIndex(configName, ".")
	if sp == -1 {
		v.SetConfigName(configName)
	} else {
		v.SetConfigName(configName[:sp])
		v.SetConfigType(configName[sp+1:])
	}

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	configFileUsed = v.ConfigFileUsed()

	_ = envconfig.Process("", cfg)
	err = v.Unmarshal(cfg)
	if err != nil {
		return
	}

	return
}

func Load(configName string, cfg interface{}) (configFileUsed string, err error) {
	return LoadOnConfigPath(configName, []string{
		"./", "./config/", "../", "../config/", "../../", "../../config/",
	}, cfg)
}
