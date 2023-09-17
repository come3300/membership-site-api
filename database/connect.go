package database

import (
	"database/sql"
	"fmt"
	"log"

	// . "go-jet-env/database/membership/table"

	_ "github.com/go-sql-driver/mysql"
)

// TODO テーブルを再編成

// dbを接続するためのメソッド

func DBconect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "webuser:webpass@tcp(db:3306)/membership")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()
	fmt.Println("設定完了！")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return nil, err
	} else {
		fmt.Println("データベース接続成功")
	}

	return db, nil

}
