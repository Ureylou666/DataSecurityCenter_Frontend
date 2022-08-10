package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	//Server相关
	AppMode 	string
	HttpPort 	string
	//Database相关
	Db         	string
	DbHost 		string
	DbPort		string
	DbUser		string
	DbPassword	string
	DbName		string
	// 阿里云AKSK
	AKSK		Aliyun
)

type Aliyun struct {
	ISDP AliyunKeys
}

type AliyunKeys struct {
	AccessKey string
	AccessSecret string
}

func init() {
	viper.SetConfigFile("/Users/ureylou/Downloads/golang/complianceCenter/backend/config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil { fmt.Println("配置文件读取错误：", err) }
	fmt.Println(viper.Get("Server.HttpPort"))
	LoadServer()
	LoadDatabase()
}

// 读取服务器配置
func LoadServer() {
	AppMode = viper.GetString("Server.Appmode")
	HttpPort =viper.GetString("Server.HttpPort")
}

// 读取数据库配置
func LoadDatabase()  {
	Db = viper.GetString("Database.Db")
	DbHost = viper.GetString("Database.DbHost")
	DbPort = viper.GetString("Database.DbPort")
	DbUser = viper.GetString("Database.DbUser")
	DbPassword = viper.GetString("Database.DbPassword")
	DbName = viper.GetString("Database.DbName")
}
