package ut

import (
	"testing"

	"github.com/sgostarter/libconfig"
)

const (
	defaultUTConfig = "ut.yaml"

	CfgItemRedis = 0x01
	CfgItemMySQL = 0x02
)

func SetupUTConfig() *Config {
	return SetupUTConfigEx(defaultUTConfig, nil)
}

func SetupUTConfigEx(fileName string, configPaths []string) *Config {
	cfg := &Config{}

	var err error

	if len(configPaths) == 0 {
		_, err = libconfig.Load(fileName, cfg)
	} else {
		_, err = libconfig.LoadOnConfigPath(fileName, configPaths, cfg)
	}

	if err != nil {
		return nil
	}

	return cfg
}

func SetupUTConfig4Redis(t *testing.T) *Config {
	return SetupUTConfig4RedisEx(defaultUTConfig, nil, t)
}

func SetupUTConfig4RedisEx(fileName string, configPaths []string, t *testing.T) *Config {
	cfg := SetupUTConfigEx(fileName, configPaths)
	if cfg == nil || cfg.RedisDSN == "" {
		t.SkipNow()

		return nil
	}

	return cfg
}

func SetupAndCheckUTConfig(checkItems int, t *testing.T) *Config {
	return SetupAndCheckUTConfigGetEx(defaultUTConfig, nil, checkItems, t)
}

func SetupAndCheckUTConfigGetEx(fileName string, configPaths []string, checkItems int, t *testing.T) *Config {
	cfg := SetupUTConfigEx(fileName, configPaths)
	if cfg == nil {
		t.SkipNow()

		return nil
	}

	if checkItems&CfgItemRedis == CfgItemRedis {
		if cfg.RedisDSN == "" {
			t.SkipNow()

			return nil
		}
	}

	if checkItems&CfgItemMySQL == CfgItemMySQL {
		if cfg.MysqlDSN == "" {
			t.SkipNow()

			return nil
		}
	}

	return cfg
}
