package controller

import (
	"context"
	"database/sql"

	// "go-jet-env/database/membership/model"
	"net/http"

	"go-jet-env/database"

	"github.com/gin-gonic/gin"
)

func getSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func postSignup(c *gin.Context) {
	r := database.RDBaccount{DB: c.MustGet("db").(*sql.DB)} // RDBaccount インスタンスを取得

	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	user, err := r.Signup(context.Background(), id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}

// func postSignup(c *gin.Context, r database.RDBaccount, ctx context.Context) {
// 	id := c.PostForm("user_id")
// 	pw := c.PostForm("password")
// 	user, err := r.Signup(ctx,id,pw)
// 	// !上記が改変が必要な部分 Signup処理を渡す必要がある
// 	if err != nil {
// 		c.Redirect(301, "/signup")
// 		return
// 	}
// 	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
// }

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func postLogin(c *gin.Context) {
	r := database.RDBaccount{DB: c.MustGet("db").(*sql.DB)} // RDBaccount インスタンスを取得

	id := c.PostForm("user_id")
	pw := c.PostForm("password")

	user, err := r.Login(context.Background(), id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "top.html", gin.H{"user": user})
}

// func postLogin(c *gin.Context, r database.RDBaccount, ctx context.Context) {
// 	id := c.PostForm("user_id")
// 	pw := c.PostForm("password")

// 	user, err := r.Login(ctx, id, pw)
// 	// !上記が改変が必要な部分  login処理を渡す必要がある
// 	if err != nil {
// 		c.Redirect(301, "/login")
// 		return
// 	}
// 	c.HTML(http.StatusOK, "top.html", gin.H{"user": user})
// }
