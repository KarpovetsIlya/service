package handlers

import (
	"database/sql"
	"net/http"
	"service/iternal/domain/model"

	"github.com/gin-gonic/gin"
)

func CreateStocks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newStock model.Stock

		if err := c.BindJSON(&newStock); err != nil {
			return
		}

		insertStmt, err := db.Prepare(`INSERT INTO "stock" (imageStock) VALUES ($1)`)
		if err != nil {
			panic(err)
		}

		defer insertStmt.Close()

		_, err = insertStmt.Exec(newStock.ImageStock)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, newStock)
	}
}

func GetStocks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT * FROM "stock"`)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var stocks []model.Stock
		for rows.Next() {
			var s model.Stock

			err = rows.Scan(&s.ID, &s.ImageStock)
			if err != nil {
				panic(err)
			}
			stocks = append(stocks, s)
			if err != nil {
				panic(err)
			}
			c.IndentedJSON(http.StatusOK, stocks)
		}
	}
}

func UpdateStocks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateStock model.Stock

		if err := c.BindJSON(&updateStock); err != nil {
			return
		}

		updateStmt, err := db.Prepare(`UPDATE "stock" set imageStock = $1 where id = $2`)
		if err != nil {
			panic(err)
		}

		defer updateStmt.Close()

		_, err = updateStmt.Exec(updateStock.ImageStock, updateStock.ID)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, updateStock)
	}
}

func DeleteStocks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		deleteStmt, err := db.Prepare(`DELETE FROM "stock" where id = $1`)
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
