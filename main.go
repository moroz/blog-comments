package main

import (
	"database/sql"
	"fmt"
	"log"
	"uuid"

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

	var version string
	err = db.QueryRow(`select sqlite_version();`).Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)

	_, err = db.Exec(`create table test (id uuid not null primary key, inserted_at datetime not null default (unixepoch()))`)
	if err != nil {
		log.Fatal(err)
	}

	var id uuid.UUID
	err = db.QueryRow(`insert into test (id) values (?) returning id`, uuid.NewV7()).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	var count int
	err = db.QueryRow(`select count(*) from test`).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
