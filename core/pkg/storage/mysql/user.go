package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/obarbier/custom-app/core/pkg/storage"
)

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	InsertUserStatement        = "INSERT INTO USER (USERNAME, PASSWORD_HASH, POLICY, CREATED_DATE, UPDATED_DATE) VALUES (?,?,?,CURRENT_TIMESTAMP(),?)"
	UpdatePolicyStatement      = "UPDATE USER SET POLICY = ? WHERE USER_ID = ?"
	DeleteUserStatement        = "DELETE FROM USER  WHERE USER_ID = ?"
	UpdateUserStatement        = "UPDATE USER SET USERNAME = ?, PASSWORD_HASH=? , UPDATED_DATE= CURRENT_TIMESTAMP() WHERE USER_ID = ?"
	SelectPolicyByUserID       = "SELECT POLICY FROM USER WHERE USER_ID=?"
	SelectSingleUserByUserName = "SELECT USER_ID, USERNAME, PASSWORD_HASH FROM USER WHERE USERNAME=?"
	SelectAllUser              = "SELECT USER_ID, USERNAME, PASSWORD_HASH FROM USER"
	SelectSingleUserByUserID   = "SELECT USER_ID, USERNAME, PASSWORD_HASH FROM USER WHERE USER_ID=?"
)

type UserStorage struct {
	db *sql.DB
}

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

func NewMysqlStorage(cfg ...Configure) (*UserStorage, error) {
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

	return &UserStorage{db: db}, nil
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

func (u *UserStorage) CreateUser(user *models.User) (*models.User, error) {
	// Prepare statement for inserting data
	stmtIns, err := u.db.Prepare(InsertUserStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // TODO(obarbier): Close the statement when we leave main() / the program terminates

	res, err := stmtIns.Exec(user.UserName, user.Password, nil, nil)
	if err != nil {

		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log_utils.Error(err) // TODO(obarbier): error handling
	}

	user.ID = lastId
	return user, nil
}

func (u *UserStorage) ReadUser(i int64) (*models.User, error) {
	stmt, err := u.db.Prepare(SelectSingleUserByUserID)
	if err != nil {
		log_utils.Error(err)
	}
	defer stmt.Close()
	p := new(models.User)
	err = stmt.QueryRow(i).Scan(&p.ID, &p.UserName, &p.Password)
	if err != nil {
		log_utils.Error(err)
	}
	log_utils.Debug(fmt.Sprintf("%+v", *p))
	return p, nil
}

func (u *UserStorage) FindAllUser() ([]*models.User, error) {
	stmt, err := u.db.Prepare(SelectAllUser)
	if err != nil {
		log_utils.Error(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log_utils.Error(err)
		return nil, err
	}
	defer rows.Close()
	var allUsers []*models.User
	for rows.Next() {
		p := new(models.User)
		err = rows.Scan(&p.ID, &p.UserName, &p.Password)
		if err != nil {
			log_utils.Error(err)
			return nil, err
		}
		allUsers = append(allUsers, p)
	}
	if err = rows.Err(); err != nil {
		log_utils.Error(err)
		return nil, err
	}
	return allUsers, nil
}

func (u *UserStorage) UpdateUser(user *models.User) error {
	// Prepare statement for inserting data
	stmtIns, err := u.db.Prepare(UpdateUserStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // TODO(obarbier): Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(user.UserName, user.Password, user.ID)
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	return err
}

func (u *UserStorage) DeleteUser(i int64) error {
	// Prepare statement for inserting data
	stmtIns, err := u.db.Prepare(DeleteUserStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // TODO(obarbier): Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(i)
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	return err
}

func (u *UserStorage) FindByUserName(username string) (*models.User, error) {
	stmt, err := u.db.Prepare(SelectSingleUserByUserName)
	if err != nil {
		log_utils.Error(err)
	}
	defer stmt.Close()
	p := new(models.User)
	err = stmt.QueryRow(username).Scan(&p.ID, &p.UserName, &p.Password)
	if err != nil {
		log_utils.Error(err)
	}
	log_utils.Debug(fmt.Sprintf("%+v", *p))
	return p, nil
}

func (u *UserStorage) CreatePolicy(i int64, policy *models.Policy) error {
	// Prepare statement for inserting data
	stmtIns, err := u.db.Prepare(UpdatePolicyStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // TODO(obarbier): Close the statement when we leave main() / the program terminates
	b, err := json.Marshal(policy)
	if err != nil {
		log_utils.Error("%s", err)
		return err
	}
	_, err = stmtIns.Exec(b, i)
	if err != nil {

		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}

	return err

}

func (u *UserStorage) GetPolicy(i int64) (*models.Policy, error) {
	stmt, err := u.db.Prepare(SelectPolicyByUserID)
	if err != nil {
		log_utils.Error(err)
	}
	defer stmt.Close()
	p := new(models.Policy)
	var policyjson []byte
	err = stmt.QueryRow(i).Scan(&policyjson)
	if err != nil {
		log_utils.Error(err)
	}
	json.Unmarshal(policyjson, p)
	log_utils.Debug(fmt.Sprintf("%+v", *p))
	return p, nil
}

var _ storage.UserStorage = &UserStorage{}
