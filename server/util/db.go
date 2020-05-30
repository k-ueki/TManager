package util

import "github.com/jinzhu/gorm"

const table = "tmanager"

func NewDB() (*gorm.DB, error) {
	//CONNECT := "root:@tcp(localhost:3306)/" + table
	CONNECT := "root:@/" + table
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		return nil, err
	}
	return db, nil
}
