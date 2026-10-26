package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/moroz/comments/config"
	_ "turso.tech/database/tursogo"
	_ "turso.tech/database/tursogo-serverless"
)

func main() {
	db, err := sql.Open(config.DatabaseSettings())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var result string
	if err := db.QueryRow("select unixepoch()").Scan(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
