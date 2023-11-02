package datastore

import (
	"errors"
	"fmt"
	"log"
	"user/internal/constants"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLConn *gorm.DB

func ConnectMySQL() error {
	if MySQLConn != nil {
		return nil
	}
	host := viper.GetString(constants.MYSQL_HOST)
	port := viper.GetString(constants.MYSQL_PORT)
	dbName := viper.GetString(constants.MYSQL_DATABASE_NAME)
	user := viper.GetString(constants.MYSQL_USERNAME)
	password := viper.GetString(constants.MYSQL_PASSWORD)
	sourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%%2FKolkata", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(sourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
		return errors.Join(err, errors.New("unable to connect to mysql db"))
	}
	MySQLConn = db
	return nil
}
