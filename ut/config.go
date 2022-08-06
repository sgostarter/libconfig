package ut

type Config struct {
	RedisDSN string `yaml:"redis_dsn" envconfig:"REDIS_DSN"`
	MysqlDSN string `yaml:"mysql_dsn" envconfig:"MYSQL_DSN"`
}
