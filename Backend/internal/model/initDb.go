package model

// 用于配置连接数据库
import (
	"Backend/internal/application/setting"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDb() {
	//dsn := Utils.DbUser + ":" + Utils.DbPassword + "@tcp(" + Utils.DbHost + ":" + Utils.DbPort + ")/" + Utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=" + setting.DbHost + " user=" + setting.DbUser + " password=" + setting.DbPassword + " dbname=" + setting.DbName + " port=" + setting.DbPort + " sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}
	db.AutoMigrate(&DataInventory{},&DataColumn{}, &DataRules{})
	// 连接池配置
	//sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	//sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。不要设置大于框架timeout时间
	//sqlDB.SetConnMaxLifetime(time.Second * 10)
	// 关闭数据库
	//sqlDB.Close()
}

func RestoreData() {
	initInventrory()
	initColumn()
}

// 初始化清空数据实例资产
func initInventrory() {
	db.Exec("DELETE FROM data_inventory")
}

func initColumn() {
	db.Exec("DELETE FROM data_column")
}
