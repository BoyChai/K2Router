package control

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func init() {
	var err error
	viper.SetConfigType("yaml")
	cave, _ := os.Open("conf.yaml")
	err = viper.ReadConfig(cave)
	if err != nil {
		fmt.Println("Log:viper读取配置出现错误:", err)
		return
	}
	var conf config
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Println("Log:viper加载配置:", err)
		return
	}
	TargetCIDR = conf.TargetCIDR
	IpPool = append(IpPool, conf.IP)
}

type config struct {
	IP         string `yaml:"IP"`
	TargetCIDR string `yaml:"TargetCIDR"`
}
