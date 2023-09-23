package db

import (
	"fmt"
	"url-shortener-service/models"

	"url-shortener-service/utils"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func NewGormDb() *gorm.DB {
	dbUrl := viper.GetString("db.connectionString")
	opts := &gorm.Config{}
	fmt.Println("Trying to connect DB........")
	db, err = gorm.Open(postgres.Open(dbUrl), opts)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	fmt.Println("Database connected!")
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(viper.GetInt("db.maxConnection"))
	sqlDb.SetMaxOpenConns(viper.GetInt("db.maxConnection"))
	// sqlDb.LogMode(viper.GetBool("DB_LOG_MODE"))

	return db
}

func Insert(model interface{}) (tx *gorm.DB) {
	if db == nil {
		return nil
	}
	return db.Create(model)
}
func Update(model interface{}) (tx *gorm.DB) {
	if db == nil {
		return nil
	}
	return db.Save(model)
}
func Delete(model interface{}) (tx *gorm.DB) {
	if db == nil {
		return nil
	}
	return db.Delete(model)
}
func GetItemByValue(field string, value string) (tx *gorm.DB) {
	return db.Where(field+" = ?", value)
}
func GetRecordByValue(table string, field string, value string) (tx *gorm.DB) {
	return db.Table(table).Where(field+" = ?", value)
}
func GetRecordByQuery(params *models.QueryParams) (tx *gorm.DB) {
	queryString := utils.GenerateQuery(params)
	return db.Table(params.Table).Where(queryString)
}
