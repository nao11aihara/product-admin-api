package models

import (
	"time"
)

// モデル
// 構造体定義、DBにアクセスする関数、レシーバメソッドを記載する。

// ショップ構造体
type Shop struct {
	Id	string	`json:"id"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

// IDからショップ1件を取得する
func FetchShopById(id string) (shop Shop, err error){
	cmd := `
		select id, name, description, created_at, updated_at
		from shops
		where id = ?
	`

	shop = Shop{}

	err = Db.QueryRow(cmd, id).Scan(
		&shop.Id,
		&shop.Name,
		&shop.Description,
		&shop.CreatedAt,
		&shop.UpdatedAt,
	)

	return shop, err
}