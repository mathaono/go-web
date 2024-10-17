package controllers

import (
	"go-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Variável que armazena o path de onde estão os templates HTML
// Variável do tipo template com a func Must() usada para trabalhar com template HTML no código Go
// Func ParseGlob() é utilizada para passar o path de onde estão os templates HTML
var temp = template.Must(template.ParseGlob("templates/*.html"))

// func index recebe uma convensão padrão para todas as requisições http (w http.ResponseWriter, r *http.Request)
func Index(w http.ResponseWriter, r *http.Request) {
	//Variável que armazena o retorno da func SearchProducts()
	allProducts := models.SearchProducts()

	// Utilizamos a variável temp do tipo template para acessarmos nosso template HTML
	// Parâmetro 1 = "w", quem consegue acessar a resposta da requisição
	// Parâmetro 2 = "Index", o template que será exibido na requisição
	// Parâmetro 3 = "products", caso haja alguma alteração no template
	temp.ExecuteTemplate(w, "Index", allProducts)

}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		convQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.CreateNewProduct(name, description, convPrice, convQuantity)
	}

	http.Redirect(w, r, "/", http.StatusCreated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProductId(productId)

	http.Redirect(w, r, "/", 200)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		quantityStr := r.FormValue("quantity")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Erro na conversão do ID: ", err)
		}

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		models.UpdateProduct(id, name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}
