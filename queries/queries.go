package main

import (
	"context"
	"fmt"
	"time"

	"go-course-l2-db/connections"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	//conn := connections.InitConnection(ctx)
	pool := connections.InitPool(ctx)
	defer pool.Close()

	exec(ctx, pool)
	queryRow(ctx, pool)
	query(ctx, pool)
}

func exec(ctx context.Context, pool *pgxpool.Pool) {
	commandTag, err := pool.Exec(ctx,
		"UPDATE films SET rating = $1 WHERE title = $2;",
		8.9, "Интерстеллар",
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandTag.String())
}

func queryRow(ctx context.Context, pool *pgxpool.Pool) {
	row := pool.QueryRow(ctx,
		"SELECT id, title, release_date, director_id, uuid, rating FROM films ORDER BY id DESC;",
	)
	var r film
	err := row.Scan(&r.id, &r.title, &r.releaseDate, &r.directorID, &r.uuid, &r.rating)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.String())
}

func query(ctx context.Context, pool *pgxpool.Pool) {
	rows, err := pool.Query(ctx, "SELECT id FROM films;")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r int

		err = rows.Scan(&r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r)
	}

	fmt.Println(rows.Err())
}

type film struct {
	id          int
	title       string
	releaseDate time.Time
	directorID  int
	uuid        pgtype.UUID
	rating      float64
}

func (f film) String() string {
	return fmt.Sprintf("film{id: %d, title: %s, releaseDate: %s, directorID: %d, uuid: %s, rating: %f}",
		f.id, f.title, f.releaseDate, f.directorID, f.uuid.String(), f.rating)
}
