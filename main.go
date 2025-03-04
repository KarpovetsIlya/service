package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host	 = "localhost"
	port	 = 5432
	user	 = "postgres"
	password = "1337"
	dbname	 = "postgres"
)

var db *sql.DB 

type product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Volume      string `json:"volume"`
	Alcohol     string `json:"alcohol"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Category    string `json:"category"`
}

type stock struct {
	ID 		   int `json:"id"`
	ImageStock string `json:"imageStock"`
}

type store struct {
	ID 		   int `json:"id"`
	Address    string `json:"address"`
	Coordinates string `json:"coordinates"`
}

func createProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	insertStmt, err := db.Prepare(`INSERT INTO "product" (name, volume, alcohol, description, price, image, category) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	CheckError(err)

	defer insertStmt.Close()

	_, err = insertStmt.Exec(newProduct.Name, newProduct.Volume, newProduct.Alcohol, newProduct.Description, newProduct.Price, newProduct.Image, newProduct.Category)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProducts(c *gin.Context) {
	rows, err := db.Query(`SELECT * FROM "product"`)
	CheckError(err)

	defer rows.Close()

	var products []product
	for rows.Next() {
		var p product
		
		err = rows.Scan(&p.ID, &p.Name, &p.Volume, &p.Alcohol, &p.Description, &p.Price, &p.Image, &p.Category)
		CheckError(err) 
		products = append(products, p)
	}
	CheckError(err)
	c.IndentedJSON(http.StatusOK, products)
}

func updateProducts(c *gin.Context) {
	var updateProduct product

	if err := c.BindJSON(&updateProduct); err != nil {
		return
	}

	updateStmt, err := db.Prepare(`UPDATE "product" set name = $1, volume = $2, alcohol = $3, description = $4, price = $5, image = $6, category = $7 where id = $8`)
	CheckError(err)

	defer updateStmt.Close()

	_, err = updateStmt.Exec(updateProduct.Name, updateProduct.Volume, updateProduct.Alcohol, updateProduct.Description, updateProduct.Price, updateProduct.Image, updateProduct.Category, updateProduct.ID)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, updateProduct)
}

func deleteProducts(c *gin.Context) {
	id := c.Param("id")

	deleteStmt, err := db.Prepare(`DELETE FROM "product" where id = $1`)
	CheckError(err)

	defer deleteStmt.Close()

	_, err = deleteStmt.Exec(id)
	CheckError(err)

	c.IndentedJSON(http.StatusNoContent, nil)
}

func createStocks(c *gin.Context) {
	var newStock stock

	if err := c.BindJSON(&newStock); err != nil {
		return
	}

	insertStmt, err := db.Prepare(`INSERT INTO "stock" (imageStock) VALUES ($1)`)
	CheckError(err)

	defer insertStmt.Close()

	_, err = insertStmt.Exec(newStock.ImageStock)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, newStock)
}

func getStocks(c *gin.Context) {
	rows, err := db.Query(`SELECT * FROM "stock"`)
	CheckError(err)

	defer rows.Close()

	var stocks []stock
	for rows.Next() {
		var s stock
		
		err = rows.Scan(&s.ID, &s.ImageStock)
		CheckError(err)
		stocks = append(stocks, s)
	}
	CheckError(err)
	c.IndentedJSON(http.StatusOK, stocks)
}

func updateStocks(c *gin.Context) {
	var updateStock stock

	if err := c.BindJSON(&updateStock); err != nil {
		return
	}

	updateStmt, err := db.Prepare(`UPDATE "stock" set imageStock = $1 where id = $2`)
	CheckError(err)

	defer updateStmt.Close()

	_, err = updateStmt.Exec(updateStock.ImageStock, updateStock.ID)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, updateStock)
}

func deleteStocks(c *gin.Context) {
	id := c.Param("id")

	deleteStmt, err := db.Prepare(`DELETE FROM "stock" where id = $1`)
	CheckError(err)

	defer deleteStmt.Close()

	_, err = deleteStmt.Exec(id)
	CheckError(err)

	c.IndentedJSON(http.StatusNoContent, nil)
}

func createStores(c *gin.Context) {
	var newStore store

	if err := c.BindJSON(&newStore); err != nil {
		return
	}

	insertStmt, err := db.Prepare(`INSERT INTO "store" (address, coordinates) VALUES ($1, $2)`)
	CheckError(err)

	defer insertStmt.Close()

	_, err = insertStmt.Exec(newStore.Address, newStore.Coordinates)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, newStore)
}

func getStores(c *gin.Context) {
	rows, err := db.Query(`SELECT * FROM "store"`)
	CheckError(err)

	defer rows.Close()

	var stores []store
	for rows.Next() {
		var s store
		
		err = rows.Scan(&s.ID, &s.Address, s.Coordinates)
		CheckError(err)
		stores = append(stores, s)
	}
	CheckError(err)
	c.IndentedJSON(http.StatusOK, stores)
}

func updateStores(c *gin.Context) {
	var updateStore store

	if err := c.BindJSON(&updateStore); err != nil {
		return
	}

	updateStmt, err := db.Prepare(`UPDATE "store" set address = $1, coordinates = $2 where id = $3`)
	CheckError(err)

	defer updateStmt.Close()

	_, err = updateStmt.Exec(updateStore.Address, updateStore.Coordinates, updateStore.ID)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, updateStore)
}

func deleteStores(c *gin.Context) {
	id := c.Param("id")

	deleteStmt, err := db.Prepare(`DELETE FROM "store" where id = $1`)
	CheckError(err)

	defer deleteStmt.Close()

	_, err = deleteStmt.Exec(id)
	CheckError(err)

	c.IndentedJSON(http.StatusNoContent, nil)
}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	router := gin.Default()
	router.Static("/img", "./img")
	router.GET("/products", getProducts)
	router.POST("/products", createProducts)
	router.PUT("/products", updateProducts)
	router.DELETE("/products/:id", deleteProducts)
	router.GET("/stocks", getStocks)
	router.POST("/stocks", createStocks)
	router.PUT("/stocks", updateStocks)
	router.DELETE("/stocks/:id", deleteStocks)
	router.GET("/stores", getStores)
	router.POST("/stores", createStores)
	router.PUT("/stores", updateStores)
	router.DELETE("/stores/:id", deleteStores)

	router.Run(":8000")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}