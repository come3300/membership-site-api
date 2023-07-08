package main

import (
	"database/sql"
	"fmt"
	"log"

	. "go-jet-env/database/go_mysql8_development/table"
	_ "github.com/go-sql-driver/mysql"
)




func main() {
	db, err := sql.Open("mysql", "webuser:webpass@tcp(db:3306)/go_mysql8_development")
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
		return
	} else {
		fmt.Println("データベース接続成功")
	}

	rows, err := 
	User.INSERT(User.ID,User.Name,User.Mail).
    VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial").Exec(db)
	

	if err != nil {
		return
	}

	updatedCount, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("ユーザーIDは重複していません")
	fmt.Println("更新された行数:", updatedCount)

}
