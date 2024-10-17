package models

import (
	"go-web/db"
	"log"
)

// Estrutura de variáveis para Produtos
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

// func SearchProducts() recebe uma convensão padrão para todas as requisições http (w http.ResponseWriter, r *http.Request)
func SearchProducts() []Product {
	// Chamando a função de conexão com a base de dados connectDataBase()
	db := db.ConnectDataBase()

	// Setando e executando a query na tabela dentro da base de dados
	rows, err := db.Query("select * from products order by id asc;")
	if err != nil {
		panic(err)
	}

	// Criando um SLICE do tipo struct Product E uma variável do tipo struct Product, essa variável será armazenada dentro do slice
	var product Product
	var products []Product

	// Esse loop percorre cada linha da tabela products; a função db.Query() roda uma query na tabela products e armazena o resultado na variável "rows"
	for rows.Next() {

		// Criando uma VARIÁVEL do tipo struct Product
		var name, description string
		var price float64
		var id, quantity int

		//A função rows.Scan() atribui os valores de cada coluna da tabela aos respectivos valores da struct Product, armazenada na VARIÁVEL product
		if err := rows.Scan(&id, &name, &description, &price, &quantity); err != nil {
			panic(err)
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		// A VARIÁVEL do tipo struct Product é adicionada ao SLICE do tipo Product, não é possível adicionar um campo avulso (name, desc e etc), pois o item do slice é uma struct completa, os campos são atributos de cada um desses itens
		products = append(products, product)
	}
	defer db.Close()

	return products
	// É necessário usar código GO Template no template HTML = {{define "Index"}} ...código HTML ... {{end}}
	// É usado {{range}}...{{end}} dentro do template HTML onde desejamos aplicar os itens da nossa lista products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDataBase()

	insert, err := db.Prepare("insert into products(name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		log.Println("insert error on database: ", err)
	}

	insert.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProductId(id string) {
	db := db.ConnectDataBase()

	delete, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		log.Println("Erro ao deletar item da base de dados: ", err)
	}

	delete.Exec(id)

	defer db.Close()
}

// Essa função é responsável por abrir o formulário para que possamos editar os valores do produto
// Ela trás as informações do banco de dados do produto selecionado e mapeia cada valor para seus respectivos campos no formulário
func EditProduct(id string) Product {
	db := db.ConnectDataBase()
	query := "select * from products where id=$1"

	productDataBase, err := db.Query(query, id)
	if err != nil {
		log.Println("Erro ao atualizar item da base de dados: ", err)
	}

	updateProduct := Product{}

	//A função Next() do pacote database trás um resultado de uma linha.
	//Essa linha de resultado é trazida através da função Scan() também do pacote database, que trás as informações do banco de dados
	//Assim podemos armazenar essas informações em variáveis e consequentemente jogá-las dentro de um objeto ou slice, para mostrá-las na aplicação (seja API ou Web)
	for productDataBase.Next() {
		var id, quantity int
		var name, description string
		var price float64

		//Usamos o "E" comercial (&) para trazer o valor da variável do banco e armazenar na variável criada acima
		err := productDataBase.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			log.Println("Erro ao trazer informações de itens da base de dados: ", err)
		}

		updateProduct.ID = id
		updateProduct.Name = name
		updateProduct.Description = description
		updateProduct.Price = price
		updateProduct.Quantity = quantity
	}

	defer db.Close()

	return updateProduct
}

// Essa função é responsável por de fato ATUALIZAR na base de dados, os valores do produto selecionado
func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDataBase()
	query := "update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5"

	update, err := db.Prepare(query)
	if err != nil {
		log.Println("Erro ao atualizar base de dados: ", err)
	}

	update.Exec(name, description, price, quantity, id)

	defer db.Close()
}
