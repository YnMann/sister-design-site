package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/YnMann/go-sister-design-site/internal/config"
	"github.com/YnMann/go-sister-design-site/pkg/logger"
	"github.com/tinrab/retry"
)

type MariaDBConnections struct {
	Default *sql.DB
}

func NewMariaDBConnection(cfg config.MariaDBConnections) (*MariaDBConnections, error) {
	var defaultDB *sql.DB
	var err error

	retry.ForeverSleep(time.Second*2, func(i int) error {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
			cfg.Default.User, cfg.Default.Pass, cfg.Default.Host, cfg.Default.Port, cfg.Default.DB)

		defaultDB, err = sql.Open("mysql", dsn)
		if err != nil {
			logger.Errorf("MariaDB connection (defaultDB), error: %v", err)
			return err
		}

		defaultDB.SetMaxOpenConns(100)
		defaultDB.SetMaxIdleConns(5)

		if err = defaultDB.Ping(); err != nil {
			logger.Errorf("MariaDB ping (defaultDB), error: %v", err)
			return err
		}

		return nil
	})

	return &MariaDBConnections{defaultDB}, nil
}
