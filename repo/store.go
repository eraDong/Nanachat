package repo

import "github.com/jackc/pgx/v5/pgxpool"

type store struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) *store {
	return &store{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
