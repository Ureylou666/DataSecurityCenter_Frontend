package v1

import (
	"Backend/internal/Errmsg"
	"Backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetColumn(c *gin.Context) {
	var queryinfo model.ColumnQueryInfo
	err := c.ShouldBindJSON(&queryinfo)
	if (queryinfo.GroupName == "All") { queryinfo.GroupName = "" }
	if (queryinfo.RiskLevelName == "All") { queryinfo.RiskLevelName = "" }
	// 判断用户输入
	if (err != nil) || (len(queryinfo.RuleName) > 50) || (queryinfo.PageSize > 50){
		c.JSON(http.StatusOK, gin.H{
			"Code": Errmsg.ERROR,
			"Message": Errmsg.GetErrMsg(Errmsg.Error_QueryInfo),
		})
		return
	}
	res_data, res_total, column_total := model.GetColumn(queryinfo)
	// 未获取到对应数据
	if res_data == nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": Errmsg.Error_NotFound,
			"Message": Errmsg.GetErrMsg(Errmsg.Error_NotFound),
		})
		return
	}
	// 正常恢复
	c.JSON(http.StatusOK, gin.H{
		"Code": Errmsg.SUCCESS,
		"Message": Errmsg.GetErrMsg(Errmsg.SUCCESS),
		"Column_Total": column_total,
		"Res_Total": res_total,
		"Res_Data": res_data,
	})
}
