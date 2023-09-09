package database

import (
	"context"
	"database/sql"
	"go-jet-env/database/membership/model"
	. "go-jet-env/database/membership/table"

	// "os/user"

	. "github.com/go-jet/jet/v2/mysql"
)

type RDBaccount struct {
	DB *sql.DB
}

// サインアップ

func (p RDBaccount) Signup(ctx context.Context, userid string, password string) (*model.User, error) {
	var result model.User

	db, err := DBconect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = User.INSERT(
		User.MutableColumns,
	).MODEL(model.User{
		UserId:   userid,
		Password: password,
	}).ExecContext(ctx, p.DB)
	if err != nil {
		panic(err)
	}
	return &result, err
}

// ログイン

func (p RDBaccount) Login(ctx context.Context, username string, password string) (*model.User, error) {
	var result model.User

	db, err := DBconect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = User.SELECT(
		User.AllColumns,
	).FROM(User).
		WHERE(User.UserId.EQ(String(username)).AND(User.Password.EQ(String(password)))).ExecContext(ctx, p.DB)
	if err != nil {
		panic(err)
	}

	return &result, err
}
