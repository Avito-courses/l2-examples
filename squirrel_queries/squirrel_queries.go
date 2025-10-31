package main

import (
	"context"
	"fmt"
	"log"

	"go-course-l2-db/connections"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	ctx := context.Background()

	pool := connections.InitPool(ctx)
	defer pool.Close()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	queryBuilder := psql.
		Select(
			"d.id as director_id",
			"d.first_name || ' ' || d.last_name as director_full_name",
			"COUNT(f.id) as films_count",
		).
		From("directors d").
		Join("films f ON d.id = f.director_id").
		Where(sq.GtOrEq{"f.release_date": "2000-01-01"}).
		GroupBy("d.id", "d.first_name", "d.last_name").
		Having("COUNT(f.id) >= ?", 1).
		OrderBy("films_count DESC", "director_full_name ASC")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(query)
	fmt.Println(args)

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var stat []DirectorStats
	for rows.Next() {
		var d DirectorStats
		if err = rows.Scan(&d.DirectorID, &d.DirectorName, &d.FilmsCount); err != nil {
			fmt.Println(err)
		}
		stat = append(stat, d)
	}

	for _, d := range stat {
		fmt.Printf("Director: %s, Films: %d\n", d.DirectorName, d.FilmsCount)
	}
}

type DirectorStats struct {
	DirectorID   int
	DirectorName string
	FilmsCount   int
}
