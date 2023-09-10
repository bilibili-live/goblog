package conf

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

var confPath = flag.String("conf", "conf.yaml", "config file")

var Instance IConf

func Init() error {
	flag.Parse()
	buf, err := os.ReadFile(*confPath)
	if err != nil {
		panic("parse config file error")
	}
	if err := yaml.Unmarshal(buf, &Instance); err != nil {
		panic("parse config file error")
	}
	return nil
}
