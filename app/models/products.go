package models

import (
	"fmt"
	"time"
)

// モデル
// 構造体定義、DBにアクセスする関数、レシーバメソッドを記載する。

// 商品構造体
type Product struct {
	Id	string	`json:"id"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

// 商品一覧取得レスポンス構造体
type ProductListRes struct {
	Products	[]Product	`json:"products"`
	Pagination	Pagination `json:"pagination"`
}

// 商品の検索を行う
func SearchProducts(page int, perPage int, name string) (products []Product, err error) {
	// nameは部分一致
	cmd := `
		select id, name, description, created_at, updated_at
		from products
		where name like concat('%', ?, '%')
		limit ?
		offset ?
	`

	offset := (page * perPage) - perPage

	rows, err := Db.Query(cmd, name, perPage, offset)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
		}
		products = append(products, product)
	}
	rows.Close()

	return products, err
}

// 商品のトータル件数を取得する
func FetchProductsTotal() (total int, err error) {
	cmd := `
		select count(*)
		from products
	`

	err = Db.QueryRow(cmd).Scan(&total)

	return total, err
}

// IDから商品1件を取得する
func FetchProductById(id string) (product Product, err error) {
	cmd := `
		select id, name, description, created_at, updated_at
		from products
		where id = ?
	`

	err = Db.QueryRow(cmd, id).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return product, err
}

// 商品を作成する
func (s *Product) CreateProduct() (id string, err error) {
	cmd := `
		insert into products (name, description)
		values (?, ?)
	`

	_, err = Db.Exec(cmd, s.Name, s.Description)
	if err != nil {
		fmt.Println(err)
	}

	err = Db.QueryRow(`select last_insert_id()`).Scan(
		&id,
	)

	return id, err
}

// 指定IDの商品を更新する
func (s *Product) UpdateProductById() (err error) {
	cmd := `
		update products
		set name = ?, description = ?
		where id = ? 
	`

	_, err = Db.Exec(cmd, s.Name, s.Description, s.Id)

	return err
}

// 指定IDの商品を削除する
func DeleteProductById(id string) (err error) {
	cmd := `
		delete
		from products
		where id = ? 
	`

	_, err = Db.Exec(cmd, id)

	return err
}
