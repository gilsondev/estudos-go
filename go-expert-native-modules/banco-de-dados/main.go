package main

import (
	"fmt"

	"github.com/gilsondev/estudos-go/go-expert-native-modules/banco-de-dados/database"
	"github.com/gilsondev/estudos-go/go-expert-native-modules/banco-de-dados/models"
)

func main() {
	db, err := database.OpenConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	products, err := models.GetProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}

	// product := models.NewProduct("Product 1", 12.34)
	// err = product.Insert(db)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao inserir produto: %s\n", err)
	// } else {
	// 	fmt.Println("Produtos inseridos com sucesso!")
	// }

	// product.Name = "Updated Product"
	// err = product.Update(db)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao atualizar produto: %s\n", err)
	// } else {
	// 	fmt.Println("Produtos atualizados com sucesso!")
	// }

	// product, err := models.GetProduct(db, "f99fe647-781d-4782-8275-6f31f1928c32")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao buscar produtos: %s\n", err)
	// }

	// fmt.Println(product)

	// err = product.Delete(db)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao deletar produtos: %s\n", err)
	// } else {
	// 	fmt.Println("Produtos deletados com sucesso!")
	// }

}
