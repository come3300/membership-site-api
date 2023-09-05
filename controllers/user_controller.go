package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	//ファイル分割したモジュール
	//DB関係
	"wordapp/db"
)
// それぞれのパスを割り振るメソッド
func handle() {
	//ユーザー登録
	http.HandleFunc("/user-form", HandlerUserForm)
	//ユーザー登録確認
	http.HandleFunc("/user-confirm", HandlerUserConfirm)
	//ログイン
	http.HandleFunc("/user-login", HandlerUserFormLogin)
	//ログイン確認
	http.HandleFunc("/user-confirm-login", HandlerUserConfirmLogin)
	//単語と意味の表示
	http.HandleFunc("/word-list-display", HandlerWordDisplay)
	//単語と意味の登録
	http.HandleFunc("/word-registered", HandlerWordInput)
	//単語と意味の登録の確認
	http.HandleFunc("/word-registered-confirm", HandlerWordInputConfirm)
	// サーバーを起動
	http.ListenAndServe(":8080", nil)
}


//各処理をviewに渡すための処理

