package orm

import (
	"../config"
	"fmt"
	"github.com/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)


func InitializeDB() *gorm.DB {
	// 根据驱动配置进行初始化
	config.DB.GetConf()
	switch config.DB.GetDatabaseType(){
	case "mysql":
		return initMySqlGorm()

	default:
		return initMySqlGorm()
	}
}

func initMySqlGorm() *gorm.DB {


	dbConfig := config.DB.GetConf()

	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset +"&parseTime=True&loc=Local"

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束

	}); err != nil {
		fmt.Println("mysql connect failed")
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		fmt.Println("mysql connect success")
		return db
	}
}

