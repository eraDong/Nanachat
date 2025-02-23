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
<<<<<<< HEAD
	cfg, err := bootstrap.LoadConfig("../bootstrap")
=======
	cfg, err := bootstrap.LoadConfig("../bootstrap/")
>>>>>>> 909841ba889776a9b31cf36d3ca72ee99920909b
	if err != nil {
		log.Fatal("cannot load env, err:", err)
	}

<<<<<<< HEAD
	connPool, err := pgxpool.New(context.Background(), cfg.DBSource.DSN())
	fmt.Printf("pool: %v\n", connPool)
=======
	connPool, err := pgxpool.New(context.Background(), cfg.DBSource)
>>>>>>> 909841ba889776a9b31cf36d3ca72ee99920909b
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
