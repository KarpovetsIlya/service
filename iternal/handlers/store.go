package handlers

import (
	"database/sql"
	"net/http"
	"service/iternal/domain/model"

	"github.com/gin-gonic/gin"
)

func CreateStores(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newStore model.Store

		if err := c.BindJSON(&newStore); err != nil {
			return
		}

		insertStmt, err := db.Prepare(`INSERT INTO "store" (address, coordinates) VALUES ($1, $2)`)
		if err != nil {
			panic(err)
		}

		defer insertStmt.Close()

		_, err = insertStmt.Exec(newStore.Address, newStore.Coordinates)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, newStore)
	}
}

func GetStores(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT * FROM "store"`)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var stores []model.Store
		for rows.Next() {
			var s model.Store

			err = rows.Scan(&s.ID, &s.Address, s.Coordinates)
			if err != nil {
				panic(err)
			}
			stores = append(stores, s)
		}
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, stores)
	}
}

func UpdateStores(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateStore model.Store

		if err := c.BindJSON(&updateStore); err != nil {
			return
		}

		updateStmt, err := db.Prepare(`UPDATE "store" set address = $1, coordinates = $2 where id = $3`)
		if err != nil {
			panic(err)
		}

		defer updateStmt.Close()

		_, err = updateStmt.Exec(updateStore.Address, updateStore.Coordinates, updateStore.ID)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, updateStore)
	}
}

func DeleteStores(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		deleteStmt, err := db.Prepare(`DELETE FROM "store" where id = $1`)
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
