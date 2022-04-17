package config

type Config struct {
	MasterDB Mysql `yaml:"MasterDB"`
}

type Mysql struct {
	DSNTemplate  string `yaml:"DSNTemplate"`
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	DBName       string `yaml:"DBName"`
	Timeout      string `yaml:"Timeout"`
	ReadTimeout  string `yaml:"ReadTimeout"`
	WriteTimeout string `yaml:"WriteTimeout"`
}

func LoadConfig() (*Config, error) {
	return &Config{}, nil
}
