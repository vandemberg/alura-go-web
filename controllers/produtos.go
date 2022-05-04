package controllers

import (
	"alura-go-web/models"
	"alura-go-web/repositories"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var REDIRECT_STATUS = 301

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := repositories.BuscaProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preço", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao do preço", err)
		}

		produto := models.Produto{}
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = precoConvertido
		produto.Quantidade = quantidadeConvertida

		repositories.InsertProduto(produto)
	}

	http.Redirect(w, r, "/", REDIRECT_STATUS)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	repositories.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/", REDIRECT_STATUS)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := repositories.FindProduto(idDoProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, _ := strconv.ParseFloat(preco, 64)
		quantidadeConvertida, _ := strconv.Atoi(quantidade)
		idConvertido, _ := strconv.Atoi(id)

		produto := models.Produto{}
		produto.Id = idConvertido
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = precoConvertido
		produto.Quantidade = quantidadeConvertida

		repositories.UpdateProduto(produto)
	}

	http.Redirect(w, r, "/", REDIRECT_STATUS)
}
