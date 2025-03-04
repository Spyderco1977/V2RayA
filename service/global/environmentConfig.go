package global

import (
	"github.com/stevenroose/gonfig"
	"log"
	"sync"
)

type Params struct {
	Address string `id:"address" short:"a" default:"0.0.0.0:2017" desc:"监听地址"`
	Config  string `id:"config" short:"c" default:"/etc/v2ray/v2raya.json" desc:"V2RayA配置文件路径"`
	Mode    string `id:"mode" short:"m" desc:"可选systemctl, service, docker, common. 不设置则自动检测"`
}

var params Params

func initFunc() {
	err := gonfig.Load(&params, gonfig.Conf{
		FileDisable:       true,
		FlagIgnoreUnknown: false,
		EnvPrefix:         "V2RAYA_",
	})
	if err != nil {
		if err.Error() != "unexpected word while parsing flags: '-test.v'" {
			log.Fatal(err)
		}
	}
}

func GetEnvironmentConfig() *Params {
	var once sync.Once
	once.Do(initFunc)
	return &params
}
