package models

import (
	"database/sql"

	"alura.com/app_web/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAll() []Produto {
	db := db.DatabaseConnect()

	allProducts, err := db.Query("select * from produtos order by id asc")

	handleErr(err)

	produtos := []Produto{}

	for allProducts.Next() {
		p := fetchProduct(allProducts)
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func fetchProduct(allProducts *sql.Rows) Produto {
	var id, quantidade int
	var nome, descricao string
	var preco float64

	err := allProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

	handleErr(err)

	p := Produto{}
	p.Id = id
	p.Descricao = descricao
	p.Nome = nome
	p.Preco = preco
	p.Quantidade = quantidade
	return p
}

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.DatabaseConnect()

	insertFromDb, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	handleErr(err)

	insertFromDb.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProductByID(id string) {

	db := db.DatabaseConnect()

	deleteProductByID, err := db.Prepare("delete from produtos where id = $1")

	handleErr(err)

	deleteProductByID.Exec(id)

	defer db.Close()
}

func FindProductByID(id string) Produto {
	db := db.DatabaseConnect()

	productRow, err := db.Query("select * from produtos where id=$1", id)

	handleErr(err)

	for productRow.Next() {
		p := fetchProduct(productRow)
		defer db.Close()
		return p
	}

	defer db.Close()
	panic("Produto não encontrado")
}

func UpdateProduct(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.DatabaseConnect()

	updateProduct, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id = $5")

	handleErr(err)

	updateProduct.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
