package database

import (
	"database/sql"
	"fmt"
	"log"

	. "go-jet-env/database/membership/table"
	_ "github.com/go-sql-driver/mysql"
)

// dbを接続するためのメソッド

func DBconect() (*sql.DB, error){
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
		return nil,err
	} else {
		fmt.Println("データベース接続成功")
	}

	rows, err :=
		User.INSERT(User.ID, User.Name, User.Password).
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial").Exec(db)

	if err != nil {
		return nil,err
	}

	updatedCount, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("ユーザーIDは重複していません")
	fmt.Println("更新された行数:", updatedCount)

    return db,nil

}
