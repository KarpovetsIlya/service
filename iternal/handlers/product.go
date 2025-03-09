package handlers

import (
	"database/sql"
	"net/http"
	"service/iternal/domain/model"

	"github.com/gin-gonic/gin"
)

func CreateProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct model.Product

		if err := c.BindJSON(&newProduct); err != nil {
			return
		}

		insertStmt, err := db.Prepare(`INSERT INTO "product" (name, volume, alcohol, description, price, image, category) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
		if err != nil {
			panic(err)
		}

		defer insertStmt.Close()

		_, err = insertStmt.Exec(newProduct.Name, newProduct.Volume, newProduct.Alcohol, newProduct.Description, newProduct.Price, newProduct.Image, newProduct.Category)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, newProduct)
	}
}

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT * FROM "product"`)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var products []model.Product
		for rows.Next() {
			var p model.Product

			err = rows.Scan(&p.ID, &p.Name, &p.Volume, &p.Alcohol, &p.Description, &p.Price, &p.Image, &p.Category)
			if err != nil {
				panic(err)
			}
			products = append(products, p)
		}
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, products)
	}
}

func UpdateProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateProduct model.Product

		if err := c.BindJSON(&updateProduct); err != nil {
			return
		}

		updateStmt, err := db.Prepare(`UPDATE "product" set name = $1, volume = $2, alcohol = $3, description = $4, price = $5, image = $6, category = $7 where id = $8`)
		if err != nil {
			panic(err)
		}

		defer updateStmt.Close()

		_, err = updateStmt.Exec(updateProduct.Name, updateProduct.Volume, updateProduct.Alcohol, updateProduct.Description, updateProduct.Price, updateProduct.Image, updateProduct.Category, updateProduct.ID)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, updateProduct)
	}
}

func DeleteProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		deleteStmt, err := db.Prepare(`DELETE FROM "product" where id = $1`)
		if err != nil {
			panic(err)
		}

		defer deleteStmt.Close()

		_, err = deleteStmt.Exec(id)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusNoContent, nil)
	}
}
