package conf

import (
	"database/sql"

	"tapera/util/env"
)

const (
	envDbConnStr     = "DB_Conn_STR"
	envDbMaxIdleConn = "DB_MAX_IDLE_CONN"
	envDbMaxOpenConn = "DB_MAX_OPEN_CONN"
	envDbDialect     = "DB_DIALECT"
)

// DbConfig func
type DbConfig struct {
	ConnStr     string
	MaxIdleConn int
	MaxOpenConn int
	Dialect     string
}

// Db func
func (f *Factory) Db(cfg *DbConfig) *sql.DB {
	db, err := sql.Open(cfg.Dialect, cfg.ConnStr)
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

// Db func
func (f *EnvFactory) Db() *sql.DB {
	return f.factory.Db(&DbConfig{
		ConnStr:     env.Str(envDbConnStr, true, nil),
		MaxIdleConn: env.Int(envDbMaxIdleConn, false, 0),
		MaxOpenConn: env.Int(envDbMaxOpenConn, false, 0),
		Dialect:     env.Str(envDbDialect, false, "postgres"),
	})
}
