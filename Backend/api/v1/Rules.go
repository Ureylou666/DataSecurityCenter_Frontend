package v1

import (
	"Backend/internal/Errmsg"
	"Backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRules(c *gin.Context) {
	var queryinfo model.RulesQueryInfo
	err := c.ShouldBindJSON(&queryinfo)
	// 判断用户输入
	if (err != nil) || (len(queryinfo.RuleName) > 50) || (queryinfo.PageSize > 50){
		c.JSON(http.StatusOK, gin.H{
			"Code": Errmsg.ERROR,
			"Message": Errmsg.GetErrMsg(Errmsg.Error_QueryInfo),
		})
		return
	}
	if queryinfo.RiskLevelName == "All" { queryinfo.RiskLevelName = "" }
	resData, resTotal, rulesTotal := model.GetRules(queryinfo)
	// 未获取到对应数据
	if resData == nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": Errmsg.Error_NotFound,
			"Message": Errmsg.GetErrMsg(Errmsg.Error_NotFound),
		})
		return
	}
	// 正常恢复
	c.JSON(http.StatusOK, gin.H{
		"Code":         Errmsg.SUCCESS,
		"Message":      Errmsg.GetErrMsg(Errmsg.SUCCESS),
		"Rules_Total": 	rulesTotal,
		"Res_Total":    resTotal,
		"Res_Data":     resData,
	})
}
