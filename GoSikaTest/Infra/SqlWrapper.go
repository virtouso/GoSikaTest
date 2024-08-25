package Infra

import (
	"fmt"
	"github.com/virtouso/GoSikaTest/Domain"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// better in env variables
const dsn = "sqlserver://sa:moeen777@127.0.0.1:1433?database=sika"

var db *gorm.DB

// no time for abstract repo
func Init() {

	dsn := "sqlserver://:@127.0.0.1:1433?database=sika"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to SQL Server successfully!")

	db.AutoMigrate(&Domain.User{}, &Domain.Address{})
}

func CrateUser(user *Domain.User) {
	db.Create(&user)
}
