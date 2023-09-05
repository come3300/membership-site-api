package database

import (
	"context"
	"database/sql"
	"go-jet-env/database/membership/model"
	. "go-jet-env/database/membership/table"

	. "github.com/go-jet/jet/v2/mysql"
)

type RDBaccount struct {
	DB *sql.DB
}

// サインアップ

func (p RDBaccount) Signup(ctx context.Context, username, password string) error {
	db, err := DBconect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = User.INSERT(
		User.MutableColumns,
	).MODEL(model.User{
		Name:     username,
		Password: password,
	}).ExecContext(ctx, p.DB)
	if err != nil {
		panic(err)
	}

	return err
}

// ログイン

func (p RDBaccount) Login(ctx context.Context, username string, password string) error {
	db, err := DBconect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = User.SELECT(
		User.AllColumns,
	).FROM(User).
	WHERE(User.Name.EQ(String(username)).AND(User.Password.EQ(String(password)))).ExecContext(ctx, p.DB)
	if err != nil {
		panic(err)
	}

	return err
}
