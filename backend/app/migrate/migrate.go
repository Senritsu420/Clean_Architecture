package main

import (
	"app/db"
	"app/model"
	"fmt"
)

func main() {
	// DBのインスタンスアドレスを取得
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
