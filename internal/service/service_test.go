package service

import (
	"github.com/go-redis/redis"
	"github.com/hatlonely/go-kit/cli"
	"github.com/jinzhu/gorm"
)

var emailCli *cli.EmailCli
var mysqlCli *gorm.DB
var redisCli *redis.Client
var service *AccountService

func init() {
	var err error
	emailCli = cli.NewEmail("hatlonely@foxmail.com", "xuckndegounrbfhf")
	mysqlCli, err = cli.NewMysql(
		cli.WithMysqlAddr("localhost", 3306),
		cli.WithMysqlAuth("root", ""),
		cli.WithMysqlDatabase("account"),
	)
	if err != nil {
		panic(err)
	}
	redisCli, err = cli.NewRedis()
	if err != nil {
		panic(err)
	}
	service, err = NewAccountService(mysqlCli, redisCli, emailCli)
	if err != nil {
		panic(err)
	}
}
