package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

const (
	DIALECT = "mysql"
)

// InitDbsByConfig 初始化数据库配置
func InitDbsByConfig(cfs []*DatasourceConfig) {
	for _, cf := range cfs {
		connectDb(cf)
	}
}

func connectDb(datasource *DatasourceConfig) {
	if datasource.Port == "" {
		datasource.Port = "3306"
	}
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", datasource.User,
		datasource.Password, datasource.Host, datasource.Port, datasource.Database)
	log.Println("db args: ", args)
	_, err := gorm.Open(DIALECT, args)
	if err != nil {
		log.Panic(err)
	}
}
