package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"qlnv/pkg/config"
	"strconv"
	"sync"
)

var mysqlInstance *gorm.DB
var mysqlOnce sync.Once

func GetMySQLInstance() *gorm.DB {
	mysqlOnce.Do(func() {
		cfg := config.LoadConfig("pkg/config/config.yaml")
		dsn := cfg.Mysql.User + ":" + cfg.Mysql.Password + "@tcp(" +
			cfg.Mysql.Host + ":" + strconv.Itoa(cfg.Mysql.Port) + ")/" +
			cfg.Mysql.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to MySQL: ", err)
			return
		}
		log.Info("Connected to MySQL: ", dsn)
		mysqlInstance = db
	})

	return mysqlInstance
}
