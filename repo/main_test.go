package repo

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/eraDong/NanaChat/bootstrap"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore *Store

func TestMain(m *testing.M) {
	cfg, err := bootstrap.LoadConfig("../bootstrap")
	if err != nil {
		log.Fatal("cannot load env, err:", err)
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.DBSource.DSN())
	if err != nil {
		log.Fatal("parse config failed:", err)
	}

	poolConfig.MaxConns = 20
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 30 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	connPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	// 设置连接池参数

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
