package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"service/iternal/database"
	"service/iternal/routes"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"service/iternal/config"
)

func main() {
	db := config.DB()

	if !database.TableExists(db, "stock") {
		fmt.Println("Таблица не существует")

		m, err := migrate.New("file://iternal/database/migrations", config.GetDBConnString())
		if err != nil {
			panic(err)
		}

		err = m.Up()
		if err != nil {
			panic(err)
		}

		fmt.Println("Таблица создана!")
	}

	r := routes.SetupRouter(config.DB())

	r.Run(":8080")
}
