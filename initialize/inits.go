package initialize

import (
	"fmt"
	"log"
	"runtime"

	conf "github.com/heyuanlong/go-tools/conf"
	kgorm "github.com/heyuanlong/go-tools/db/gorm"
	klog "github.com/heyuanlong/go-tools/log"
	kruntime "github.com/heyuanlong/go-tools/runtime"
	"github.com/jinzhu/gorm"
)

const CONFIG_PATH = "conf/config.json"

var Conf *conf.Kconf

var LogError *klog.LlogFile
var LogWarn *klog.LlogFile
var LogInfo *klog.LlogFile
var LogDebug *klog.LlogFile

var Gorm *gorm.DB

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核设置

	conft, err := conf.NewKconf(CONFIG_PATH)
	if err != nil {
		log.Panic("conf parse fail:", err)
	}
	Conf = conft

	LogError = klog.Error
	LogWarn = klog.Warn
	LogInfo = klog.Info
	LogDebug = klog.Debug

}

func InitConf(confile string) {
	conft, err := conf.NewKconf(confile)
	if err != nil {
		log.Panic("conf parse fail:", err)
	}
	Conf = conft
}

func InitLog(logfile string) {
	var err error
	kruntime.CreateDir("log")
	LogError, err = klog.NewLlogFile("log/"+logfile+"_err.log", "[Error]", klog.LstdFlags|klog.Lshortfile, klog.LOG_LEVEL_ERROR, klog.LOG_LEVEL_ERROR, 50)
	if err != nil {
		log.Println(err)
	}
	LogWarn = LogError
	LogInfo = LogError
	LogDebug = LogError
}

func InitMysql() {
	mysql_user, _ := Conf.GetString("mysql.user")
	mysql_password, _ := Conf.GetString("mysql.password")
	mysql_ip, _ := Conf.GetString("mysql.ip")
	mysql_port, _ := Conf.GetInt("mysql.port")
	mysql_port_str := fmt.Sprintf("%d", mysql_port)
	mysql_mysqldb, _ := Conf.GetString("mysql.mysqldb")
	tmpGorm, err := kgorm.NewGorm(mysql_user, mysql_password, mysql_ip, mysql_port_str, mysql_mysqldb)
	if err != nil {
		log.Panic(err)
	}
	Gorm = tmpGorm
}
