package drivers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type gormDriver struct {
}

func NewGormDriver() gormDriver {
	return gormDriver{}
}

func (g gormDriver) Boot(config internalContracts.IConfiguration, models ...interface{}) interface{} {
	database, err := config.Get("Database")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", database.(string))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models...)
	return db
}
