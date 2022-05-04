package repositories

import (
	"alura-go-web/infra"
	"alura-go-web/models"
)

func FindProduto(id string) models.Produto {
	db := infra.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select *from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := models.Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()

	return produtoParaAtualizar
}

func BuscaProdutos() []models.Produto {
	db := infra.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select *from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}
	produtos := []models.Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
		defer db.Close()
	}

	return produtos
}

func InsertProduto(produto models.Produto) {
	db := infra.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)

	defer db.Close()
}

func DeletaProduto(idDoProduto string) {
	db := infra.ConectaComBancoDeDados()

	deletaProdutoDoBanco, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deletaProdutoDoBanco.Exec(idDoProduto)
	defer db.Close()
}

func UpdateProduto(produto models.Produto) {
	db := infra.ConectaComBancoDeDados()

	prepare, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	prepare.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)

	defer db.Close()
}
