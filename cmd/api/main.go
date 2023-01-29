package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/logn-soft/logn-back/internal/ent"
)

func main() {
	dsn := "postgresql://root:secret@postgres:5432/logn2?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("cant connect to db %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed closing connection to db: %v", err)
		}
	}()
	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// log the queries executed against to DB by Ent
	client = client.Debug()

	server := http.NewServeMux()
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
