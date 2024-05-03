package db

import (
	"database/sql"
	"fmt"
	"github.com/b3liv3r/trainers-for-gym/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"time"
)

func NewSqlDB(logger *zap.Logger, conf config.DB) (*sqlx.DB, error) {
	var dsn string
	var err error
	var dbRaw *sql.DB

	dsn = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Name)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * time.Duration(conf.Timeout))

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", conf.Timeout, err)
		case <-ticker.C:
			dbRaw, err = sql.Open(conf.Driver, dsn)
			if err != nil {
				logger.Error("failed to open the database", zap.String("dsn", dsn), zap.Error(err))
				return nil, err
			}
			err = dbRaw.Ping()
			if err == nil {
				db := sqlx.NewDb(dbRaw, conf.Driver)
				db.SetMaxOpenConns(conf.MaxConn)
				db.SetMaxIdleConns(conf.MaxConn)
				return db, nil
			}
			logger.Error("failed to connect to the database", zap.String("dsn", dsn), zap.Error(err))
		}
	}
}
