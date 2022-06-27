package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"alura.com/app_web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.FindAll()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		strPreco := r.FormValue("preco")
		strQuantidade := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(strPreco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidade, err := strconv.Atoi(strQuantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CreateNewProduct(nome, descricao, preco, quantidade)
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	models.DeleteProductByID(id)

	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	product := models.FindProductByID(id)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		strID := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		strPreco := r.FormValue("preco")
		strQuantidade := r.FormValue("quantidade")

		id, err := strconv.Atoi(strID)

		handleErr(err)

		preco, err := strconv.ParseFloat(strPreco, 64)

		handleErr(err)

		quantidade, err := strconv.Atoi(strQuantidade)

		handleErr(err)

		models.UpdateProduct(id, nome, descricao, preco, quantidade)

		http.Redirect(w, r, "/", 301)

	}
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
