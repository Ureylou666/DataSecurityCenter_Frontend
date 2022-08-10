package router

import (
	"Backend/api/v1"
	"Backend/internal/application/setting"
	"Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(setting.AppMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	api := r.Group("api/v1")
	// 数据安全 > 存储阶段 > 数据资产 Inventory模块
	api.POST("/inventory", v1.GetInventory)
	// 数据安全 > 存储阶段 > 数据清单 Column模块
	api.POST("/column", v1.GetColumn)
	// 数据安全 > 存储阶段 > 分级规则 rules模块
	api.POST("/rules", v1.GetRules)
	r.Run(setting.HttpPort)
}
