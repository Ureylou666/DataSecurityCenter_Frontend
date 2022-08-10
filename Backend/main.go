package main

import (
	"Backend/internal/model"
	"Backend/internal/router"
)

func main() {
	// 引用连接数据库
	model.InitDb()
	// 重置数据库中数据
	//model.RestoreData()
	// 获取初始化数据
	//aliyun.Entry("DAP")
	router.InitRouter()
}
