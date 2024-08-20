package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mari-muthu-k/gin-template/globals"
)

//connect mysql db using gorm
func ConnectRelationalDB(dbDriver string,config *gorm.Config)error{
	var dsn string
	var err error

	// Keep only the driver you use & remove the rest
	switch dbDriver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		globals.RelationalDb, err = gorm.Open(mysql.Open(dsn), config)
		

	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		globals.RelationalDb, err = gorm.Open(postgres.Open(dsn), config)
		

	case "sqlite":
		dsn = fmt.Sprintf("file:%s?cache=shared&mode=rwc", os.Getenv("DB_NAME"))
		globals.RelationalDb, err = gorm.Open(sqlite.Open(dsn), config)
		

	default:
		break
	}

	return err
}

func DisconnectRelationalDB(){
	dbSQL, err := globals.RelationalDb.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}