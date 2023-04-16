package datasource

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"lottery_llx/conf"
	"os"
	"sync"
	"time"
)

var dbLock sync.Mutex
var masterInstance *gorm.DB

func NewDbMaster() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbMaster.User, conf.DbMaster.Password, conf.DbMaster.Host, conf.DbMaster.Port, conf.DbMaster.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error

	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	masterInstance = instance
	return masterInstance
}

// InstanceDbMaster 得到唯一的主库实例
func InstanceDbMaster() *gorm.DB {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()

	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}
