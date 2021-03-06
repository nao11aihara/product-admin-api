package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

// 初期化関数
// main.goでimportするとmain関数より前に呼ばれる
func init() {
	// 第二引数: userName:password@(host:port)/dbName
	Db, err = sql.Open("mysql", "root:root_password@tcp(db:3306)/product_admin_api?parseTime=true")

	if err != nil {
		fmt.Println(err)
	}
}

// ページネーション構造体
type Pagination struct {
	Page	int	`json:"page"`
	PerPage	int	`json:"perPage"`
	Total	int	`json:"total"`
}
