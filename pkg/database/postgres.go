package database

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/AgazadeAV/my-first-go-project/ent"
	_ "github.com/lib/pq"
	"log"
)

func NewEntClient() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=user dbname=usersdb password=password sslmode=disable search_path=myschema")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
