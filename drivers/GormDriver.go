package drivers

import (
	"github.com/jinzhu/gorm"
	contracts2 "nikan.dev/pronto/internals/contracts"
)

type gormDriver struct {
}

func NewGormDriver() gormDriver {
	return gormDriver{}
}

func (g gormDriver) Boot(config contracts2.IConfiguration, models ...interface{}) interface{} {
	database, err := config.Get("Database")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", database.(string))
	if err != nil {
		panic("failed to connect models")
	}
	db.AutoMigrate(models...)
	return db

}
