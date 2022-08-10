package model

import (
	"Backend/internal/Errmsg"
)

// DataInventory 定义数据资产数据结构，与aliyun sdk中定义一致
type DataInventory struct {
	UUID				string 	`gorm:"primaryKey" json:"UUID"`
	GroupName        	string  `gorm:"type:varchar(50)" json:"GroupName"`
	CreationTime        string  `gorm:"type:varchar(50)" json:"CreationTime"` //创建该数据资产实例的时间。使用时间戳表示，单位：毫秒。
	DepartName          string 	`gorm:"type:varchar(200)" json:"DepartName"`  //数据资产实例所属部门的名称。
	InstanceId          int  	`gorm:"type:int" json:"Id"`	//数据安全中心服务中记录的数据资产实例的唯一标识ID。
	InstanceDescription string 	`gorm:"type:varchar(200)" json:"InstanceDescription"` //数据资产实例的描述信息。
	Labelsec            bool   	`gorm:"type:boolean" json:"Labelsec"`  //数据资产实例的安全状态
	LastFinishTime      string  `gorm:"type:varchar(50)" json:"LastFinishTime"` //最近一次扫描数据资产实例的完成时间。使用时间戳表示，单位：毫秒。
	Name                string 	`gorm:"type:varchar(200)" json:"Name"`  //数据资产实例的名称。
	Owner               string 	`gorm:"type:varchar(200)" json:"Owner"` //拥有该数据资产实例的阿里云账号。
	ProductCode         string 	`gorm:"type:varchar(200)" json:"ProductCode"` //数据资产实例所属产品的名称，包括MaxCompute、OSS、RDS等。关于支持的具体产品名称
	ProductId           string 	`gorm:"type:varchar(200)" json:"ProductId"`  //数据资产实例所属产品的ID。
	Protection          bool   	`gorm:"type:boolean" json:"Protection"`  //数据资产实例的防护状态
	RiskLevelId         int  	`gorm:"type:int" json:"RiskLevelId"` //数据资产实例的风险等级ID。风险等级ID越高，表示识别出的数据越敏感
	RiskLevelName       string 	`gorm:"type:varchar(200)" json:"RiskLevelName"` // 数据资产实例的风险等级名称
	RuleName            string 	`gorm:"type:varchar(200)" json:"RuleName"` //数据资产实例命中的敏感数据识别规则的名称。
	Sensitive           bool   	`gorm:"type:boolean" json:"Sensitive"`  //数据资产实例中是否包含敏感数据。
	SensitiveCount      int  	`gorm:"type:int" json:"SensitiveCount"` //数据资产实例中包含的敏感数据总数。例如：当数据资产为RDS时，表示该实例中数据库的敏感总表数
	TenantName          string 	`gorm:"type:varchar(200)" json:"TenantName"`  //租户的名称。
	TotalCount          int  	`gorm:"type:int" json:"TotalCount"` //数据资产实例中的数据总数。例如：当数据资产为RDS时，表示该实例中数据库的总表数。
}

// InventoryQueryInfo 定义api接口查询参数
type InventoryQueryInfo struct {
	GroupName 		string
	InstanceName    string
	InstanceType	string
	PageNum  		int
	PageSize 		int
}

// AddInventory 新增aliyun数据实例资产
func AddInventory(data *DataInventory) int {
	err := db.Create(&data).Error
	if err!=nil {
		return Errmsg.ERROR
	}
	return Errmsg.SUCCESS
}

// 获取aliyun数据资产列表 并进行分页展示
func GetInventory(query InventoryQueryInfo) ([]DataInventory, int64, int64) {
	var result []DataInventory
	var resTotal, inventoryTotal int64
	// 获取数据库中Instance总数
	// InstanceName 模糊搜索
	// 分页处理
	db.Where(&DataInventory{GroupName: query.GroupName,ProductCode: query.InstanceType}).Where("name like ?", "%"+query.InstanceName+"%").Find(&result).Limit(-1).Count(&inventoryTotal)
	if query.PageNum == 0 || query.PageSize == 0 {
		db.Where(&DataInventory{GroupName: query.GroupName,ProductCode: query.InstanceType}).Where("name like ?", "%"+query.InstanceName+"%").Order("total_count desc, name").Limit(-1).Find(&result)
		resTotal = int64(len(result))
	} else {
		db.Where(&DataInventory{GroupName: query.GroupName,ProductCode: query.InstanceType}).Where("name like ?", "%"+query.InstanceName+"%").Order("total_count desc, name").Limit(query.PageSize).Offset((query.PageNum - 1) * query.PageSize).Find(&result)
		resTotal = int64(len(result))
	}
	return result, resTotal, inventoryTotal
}
