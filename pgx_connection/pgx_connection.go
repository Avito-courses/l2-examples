package main

import (
	"context"
	"fmt"

	"go-course-l2-db/connections"
)

func main() {
	ctx := context.Background()

	conn := connections.InitConnection(ctx)
	defer conn.Close(ctx)

	var title string
	err := conn.QueryRow(ctx,
		"SELECT title FROM films WHERE id = $1",
		2).
		Scan(&title)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(title)
}
