package db

import (
	"github.com/jinzhu/gorm"
_ 	"github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

const db_user string = "root"
const db_pass string = "root"
const db_name string = "golang"

func Database() *gorm.DB {
	//open a db connection
	fmt.Print(fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local\n", db_user, db_pass, db_name))
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_pass, db_name)	)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}