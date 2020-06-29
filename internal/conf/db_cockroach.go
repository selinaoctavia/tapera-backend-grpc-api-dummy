package conf

import (
	"database/sql"

	"tapera/util/env"

	_ "github.com/lib/pq" //https://golang.org/doc/effective_go.html#blank_import
)

const (
	envCcrDbConnStr     = "COCKCROACH_DB_CONN_STR"
	envCcrDbMaxIdleConn = "COCKCROACH_DB_MAX_IDLE_CONN"
	envCcrDbMaxOpenConn = "COCKCROACH_DB_MAX_OPEN_CONN"
)

// CockroachConfig struct
type CockroachConfig struct {
	ConnStr     string
	MaxIdleConn int
	MaxOpenConn int
}

// CockcroachDb func
func (f *Factory) CockcroachDb(cfg *CockroachConfig) *sql.DB {
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	return db
}

// CockcroachDb func
func (f *EnvFactory) CockcroachDb() *sql.DB {
	return f.factory.CockcroachDb(&CockroachConfig{
		ConnStr:     env.Str(envCcrDbConnStr, true, nil),
		MaxIdleConn: env.Int(envCcrDbMaxIdleConn, false, 0),
		MaxOpenConn: env.Int(envCcrDbMaxOpenConn, false, 0),
	})
}
