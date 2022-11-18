package config
type Config struct {
	DB `yaml:"db"`
}

type DB struct {
	USERNAME string `yaml:"username"`
    PASSWORD string `yaml:"password"`
    HOST string `yaml:"host"`
	PORT int `yaml:"port"`
	DBNAME int `yaml:"dbname"`
}