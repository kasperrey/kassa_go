package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	id      int
	naam    string
	prijs   float64
	barcode int64
}

type Bankkaart struct {
	kaartId string
	code    string
	geld    float64
}

var db *sql.DB
var err error

func startDatabank() {
	db, err = sql.Open("mysql", "kasper:kasper@tcp(0.0.0.0:3306)/kassa")

	// handle error, if any.
	if err != nil {
		panic(err)
	}
}

func getProductByName(name string) Product {
	result, err := db.Query(fmt.Sprintf("SELECT * FROM producten WHERE naam = \"%s\"", name))
	if err != nil {
		panic(err)
	}
	for result.Next() {

		var id int
		var naam string
		var prijs float64
		var barcode int64

		err = result.Scan(&id, &naam, &prijs, &barcode)

		// handle error
		if err != nil {
			panic(err)
		}
		return Product{id: id, naam: naam, prijs: prijs, barcode: barcode}
	}
	return Product{0, "", 0, 0}
}

func getAllProducts() []Product {
	result, err := db.Query(fmt.Sprintf("SELECT * FROM producten"))
	if err != nil {
		panic(err)
	}
	var productenLijst []Product
	for result.Next() {

		var id int
		var naam string
		var prijs float64
		var barcode int64

		err = result.Scan(&id, &naam, &prijs, &barcode)

		// handle error
		if err != nil {
			panic(err)
		}
		productenLijst = append(productenLijst, Product{id: id, naam: naam, prijs: prijs, barcode: barcode})
	}
	return productenLijst
}

func getAllBankkaarts() []Bankkaart {
	result, err := db.Query(fmt.Sprintf("SELECT * FROM bankkaarten"))
	if err != nil {
		panic(err)
	}
	var bankkaartLijst []Bankkaart
	for result.Next() {

		var kaartId string
		var code string
		var geld float64

		err = result.Scan(&kaartId, &code, &geld)

		// handle error
		if err != nil {
			panic(err)
		}
		bankkaartLijst = append(bankkaartLijst, Bankkaart{kaartId: kaartId, code: code, geld: geld})
	}
	return bankkaartLijst
}

func createProduct(product Product) {
	query := "INSERT INTO producten (naam, prijs, barcode) VALUES (?, ?, ?)"
	db.ExecContext(context.Background(), query, product.naam, product.prijs, product.barcode)
}

func closeDatabase() {
	defer db.Close()
}
