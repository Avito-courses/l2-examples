package main

import (
	"go-course-l2-db/connections"

	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	conn := connections.InitPool(ctx)
	defer conn.Close()

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		"INSERT INTO directors (last_name, first_name) VALUES ($1, $2);",
		"Иван", "Иванов",
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tx.Exec(ctx,
		"DELETE FROM directors WHERE first_name = $1 AND last_name = $2;",
		"Иван", "Иванов",
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println(err)
	}
}
