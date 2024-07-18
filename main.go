package main

import (
	"database/sql"
	db2 "github.com/Igorcand/hexagonal/adapters/db"
	"github.com/Igorcand/hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	ProductService := application.NewuProductService(productDbAdapter)
	product, _ := ProductService.Create("Product example", 30)

	ProductService.Enable(product)
}