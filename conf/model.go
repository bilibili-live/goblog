package conf

type IDbConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type IConf struct {
	Port int `yaml:"port"`
	Db   IDbConf
}
