package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-jet-env/database/go_mysql8_development/model"
	. "go-jet-env/database/go_mysql8_development/table"

	. "github.com/go-jet/jet/v2/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Yamada struct {
	DB *sql.DB
}

func UpdateUserProfile(
	db *sql.DB,
	ID int32,
	Name string,
	Mail string,
) error {


	stmt := SELECT(
		User.ID, User.Name, User.Mail, 
		User.AllColumns,
	).FROM(
		User.
			INNER_JOIN(FilmActor, Actor.ActorID.EQ(FilmActor.ActorID)).
			INNER_JOIN(Film, Film.FilmID.EQ(FilmActor.FilmID)).
			INNER_JOIN(Language, Language.LanguageID.EQ(Film.LanguageID)).
			INNER_JOIN(FilmCategory, FilmCategory.FilmID.EQ(Film.FilmID)).
			INNER_JOIN(Category, Category.CategoryID.EQ(FilmCategory.CategoryID)),
	).WHERE(
		Language.Name.EQ(String("English")).
			AND(Category.Name.NOT_EQ(String("Action"))).
			AND(Film.Length.GT(Int(180))),
	).ORDER_BY(
		Actor.ActorID.ASC(),
		Film.FilmID.ASC(),
	)
	

	rows, err := User.
		SELECT(
			User.AllColumns,
		).FROM(
			User,
		).Exec(db)
	

	if err != nil {
		return err
	}

	updatedCount, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("ユーザーIDは重複していません")
	fmt.Println("更新された行数:", updatedCount)
	return nil
}

func ptrInt32(v int32) *int32 {
	return &v
}

// ポインタ型のstringを生成するヘルパー関数
func ptrString(v string) *string {
	return &v
}

func insertUser(db *sql.DB) error {

	user := model.User{
		ID:   1,
		Name: ptrString("John Doe"),
		Mail: ptrString("johndoe@example.com"),
	}

	insertQuery := User.INSERT(User.ID, User.Name, User.Mail).
		MODELS([]model.User{user})

	_, err := insertQuery.Exec(db)
	if err != nil {
		return err
	}

	return nil

}

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

}
