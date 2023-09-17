package controller

import (
	"context"
	"database/sql"
	"log"

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

func wrap() {
	r := gin.Default()

	db, err := sql.Open("mysql", "webuser:webpass@tcp(db:3306)/membership")
	if err != nil {
		log.Fatal(err)
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/signup", postSignup)

	r.Run()
}


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
