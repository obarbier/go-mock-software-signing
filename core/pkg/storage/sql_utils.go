package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
	"time"
)

type Config struct {
	dbTcpAddress    string
	dbName          string
	username        string
	password        string
	connMaxLifetime time.Duration
	maxOpenConn     int
	maxIdleConn     int
}

func SetMaxIdleConn(i int) Configure {
	return func(config *Config) {
		config.maxIdleConn = i
	}
}

func SetMaxOpenConn(i int) Configure {
	return func(config *Config) {
		config.maxOpenConn = i
	}
}

func SetConnMaxLifetime(t int) Configure {
	return func(config *Config) {
		config.connMaxLifetime = time.Minute * time.Duration(t)
	}
}
func SetCredentials(username, password string) Configure {
	return func(config *Config) {
		config.username = username
		config.password = password
	}
}

func SetDBName(name string) Configure {
	return func(config *Config) {
		config.dbName = name
	}
}

func SetDBAddress(addr string) Configure {
	return func(config *Config) {
		config.dbTcpAddress = addr
	}
}

type Configure func(config *Config)

func NewSQLStorage(cfg ...Configure) (*sql.DB, error) {
	c := &Config{}
	for _, f := range cfg {
		f(c)
	}
	// VALIDATION
	err2 := validateConfig(c)
	if err2 != nil {
		return nil, err2
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.username, c.password, c.dbTcpAddress, c.dbName)
	log_utils.Debug(dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err) // TODO(obarbier): proper error handling instead of panic in your app
	}

	db.SetConnMaxLifetime(c.connMaxLifetime)
	db.SetMaxOpenConns(c.maxOpenConn)
	db.SetMaxIdleConns(c.maxIdleConn)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}

	return db, nil
}

func validateConfig(c *Config) error {
	if c.dbName == "" {
		return fmt.Errorf("set dbName")
	}
	if c.username == "" {
		return fmt.Errorf("set database username")
	}
	if c.password == "" {
		return fmt.Errorf("set database password")
	}
	if c.connMaxLifetime == time.Duration(0) {
		c.connMaxLifetime = time.Minute * 3
	}
	if c.maxIdleConn == 0 {
		c.maxIdleConn = 2
	}
	if c.maxOpenConn == 0 {
		c.maxIdleConn = 3
	}
	return nil
}

type SQLCloser interface {
	Close() error
}

type Row interface {
	Scan(dest ...any) error
}

func ReadUser(r Row) (*models.User, error) {
	p := new(models.User)
	var policyjson []byte
	err := r.Scan(&p.ID, &p.UserName, &policyjson, &p.Password)
	if err != nil {
		errMsg := "sql database scanning error"
		log_utils.Error(errMsg, err)
		return nil, err
	}
	err = json.Unmarshal(policyjson, &p.Policy)
	if err != nil {
		errMsg := "json parsing error"
		log_utils.Error(errMsg, err)
		return nil, err
	}
	return p, nil
}

func CloseSqlResource(stmt SQLCloser) {
	err := stmt.Close()
	if err != nil {
		log_utils.Error("aql closing failed with error", err)
	}
}
