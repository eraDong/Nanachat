package repo

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DBSource = "postgres://root:root@localhost:5432/nanachat?sslmode=disable"
)

var testStore *store

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:%w", err)
	}
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
