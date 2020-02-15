package main

import (
	"flag"
	_ "gorm_generate/initialize"
	kinit "gorm_generate/initialize"
	kroute "gorm_generate/route"
	kcontrol "gorm_generate/work/control"

	kruntime "github.com/heyuanlong/go-tools/runtime"
)

func main() {
	types := flag.String("t", "", "启动类型，空：正常，race：刷比赛，bonus：刷分红，price：刷兑换比例")
	flag.Parse()

	kinit.InitConf("conf/config.json")

	if *types == "api" {
		kruntime.MainGetPanicAndLoop(func() {
			kruntime.Pid("api.pid")
			kinit.InitLog("api")
			kinit.InitMysql()

			port, _ := kinit.Conf.GetInt("server.port")
			r := kroute.NewRouteStruct(port)
			r.SetMiddleware(kroute.SetCommonHeader)

			r.Load(kcontrol.Newgorm())

			r.Run()
		})
	}
}
