package main

import (
	"context"
	_ "github.com/jackc/pgx/v5"
	"github.com/logn-soft/logn-back/internal/ent"
	"log"
	"net/http"
)

func main() {
	dsn := "postgresql://root:secret@postgres:5432/logn2?sslmode=disable"
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("cant connect to db %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("failed opening connection to db: %v", err)
		}
	}(client)

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
