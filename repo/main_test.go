package repo

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/eraDong/NanaChat/bootstrap"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore *Store

func TestMain(m *testing.M) {
	cfg, err := bootstrap.LoadConfig("../bootstrap")
	if err != nil {
		log.Fatal("cannot load env, err:", err)
	}

	connPool, err := pgxpool.New(context.Background(), cfg.DBSource.DSN())
	fmt.Printf("pool: %v\n", connPool)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
