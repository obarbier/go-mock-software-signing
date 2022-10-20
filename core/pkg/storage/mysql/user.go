package mysql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/obarbier/custom-app/core/pkg/storage"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	InsertUserStatement        = "INSERT INTO USER (USERNAME, PASSWORD_HASH, POLICY, CREATED_DATE, UPDATED_DATE) VALUES (?,?,?,CURRENT_TIMESTAMP(),?)"
	UpdatePolicyStatement      = "UPDATE USER SET POLICY = ? WHERE USER_ID = ?"
	DeleteUserStatement        = "DELETE FROM USER  WHERE USER_ID = ?"
	UpdateUserStatement        = "UPDATE USER SET USERNAME = ?, PASSWORD_HASH=? , UPDATED_DATE= CURRENT_TIMESTAMP() WHERE USER_ID = ?"
	SelectPolicyByUserID       = "SELECT POLICY FROM USER WHERE USER_ID=?"
	SelectSingleUserByUserName = "SELECT USER_ID, USERNAME, POLICY, PASSWORD_HASH FROM USER WHERE USERNAME=?"
	SelectAllUser              = "SELECT USER_ID, USERNAME, POLICY, PASSWORD_HASH FROM USER"
	SelectSingleUserByUserID   = "SELECT USER_ID, USERNAME, POLICY, PASSWORD_HASH FROM USER WHERE USER_ID=?"
)

type UserStorage struct {
	db *sql.DB
}

func NewMysqlUserStorage(cfg ...storage.Configure) (*UserStorage, error) {
	db, err := storage.NewSQLStorage(cfg...)
	if err != nil {
		log_utils.Fatal("error while initializing storage layer", err)
	}
	return &UserStorage{db: db}, nil
}

func (u *UserStorage) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	// Prepare statement for inserting data
	stmtIns, err := u.db.PrepareContext(ctx, InsertUserStatement) // ? = placeholder
	if err != nil {
		errMsg := "error while creating user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	defer storage.CloseSqlResource(stmtIns)
	b, err2 := policyToByte(err, &user.Policy)
	if err2 != nil {
		errMsg := "fail decoding json"
		log_utils.Error(errMsg, err2)
		return nil, errors.New(errMsg)
	}
	res, err := stmtIns.ExecContext(ctx, user.UserName, user.Password, b, nil)
	if err != nil {
		errMsg := "error while creating user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		errMsg := "error while creating user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)

	}

	user.ID = lastId
	return user, nil
}

func (u *UserStorage) ReadUser(ctx context.Context, i int64) (*models.User, error) {
	stmt, err := u.db.PrepareContext(ctx, SelectSingleUserByUserID)
	if err != nil {
		log_utils.Error(err)
	}
	defer storage.CloseSqlResource(stmt)
	p, err := storage.ReadUser(stmt.QueryRowContext(ctx, i))
	if err != nil {
		errMsg := "error while reading user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	log_utils.Debug(fmt.Sprintf("%+v", *p))
	return p, nil
}

func (u *UserStorage) FindAllUser(ctx context.Context) ([]*models.User, error) {
	stmt, err := u.db.PrepareContext(ctx, SelectAllUser)
	if err != nil {
		errMsg := "error while reading all user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	defer storage.CloseSqlResource(stmt)
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		errMsg := "error while reading all user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	defer storage.CloseSqlResource(rows)
	var allUsers []*models.User
	for rows.Next() {
		p, err := storage.ReadUser(rows)
		if err != nil {
			errMsg := "error while reading all user"
			log_utils.Error(errMsg, err)
			return nil, errors.New(errMsg)
		}
		allUsers = append(allUsers, p)
	}
	if err = rows.Err(); err != nil {
		errMsg := "error while reading all user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	return allUsers, nil
}

func (u *UserStorage) UpdateUser(ctx context.Context, user *models.User) error {
	// Prepare statement for inserting data
	stmtIns, err := u.db.PrepareContext(ctx, UpdateUserStatement) // ? = placeholder
	if err != nil {
		errMsg := "error while updating user"
		log_utils.Error(errMsg, err)
		return errors.New(errMsg)
	}
	defer storage.CloseSqlResource(stmtIns)
	_, err = stmtIns.ExecContext(ctx, user.UserName, user.Password, user.ID)
	if err != nil {
		errMsg := "error while updating user"
		log_utils.Error(errMsg, err)
		return errors.New(errMsg)
	}
	return err
}

func (u *UserStorage) DeleteUser(ctx context.Context, i int64) error {
	// Prepare statement for inserting data
	stmtIns, err := u.db.PrepareContext(ctx, DeleteUserStatement) // ? = placeholder
	if err != nil {
		errMsg := "error while deleting user"
		log_utils.Error(errMsg, err)
		return errors.New(errMsg)
	}
	defer storage.CloseSqlResource(stmtIns)
	_, err = stmtIns.ExecContext(ctx, i)
	if err != nil {
		errMsg := "error while deleting user"
		log_utils.Error(errMsg, err)
		return errors.New(errMsg)
	}
	return err
}

func (u *UserStorage) FindByUserName(ctx context.Context, username string) (*models.User, error) {
	stmt, err := u.db.PrepareContext(ctx, SelectSingleUserByUserName)
	if err != nil {
		errMsg := "error while finding user by name user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	defer storage.CloseSqlResource(stmt)
	p, err := storage.ReadUser(stmt.QueryRowContext(ctx, username))
	if err != nil {
		errMsg := "error while finding user by name user"
		log_utils.Error(errMsg, err)
		return nil, errors.New(errMsg)
	}
	log_utils.Debug(fmt.Sprintf("%+v", *p))
	return p, nil
}

func (u *UserStorage) CreatePolicy(_ context.Context, i int64, policy *models.Policy) error {
	// Prepare statement for inserting data
	stmtIns, err := u.db.Prepare(UpdatePolicyStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}
	defer storage.CloseSqlResource(stmtIns)
	b, err2 := policyToByte(err, policy)
	if err2 != nil {
		return err2
	}
	_, err = stmtIns.Exec(b, i)
	if err != nil {

		panic(err.Error()) // TODO(obarbier): proper error handling instead of panic in your app
	}

	return err

}

func policyToByte(err error, policy *models.Policy) ([]byte, error) {
	b, err := json.Marshal(policy)
	if err != nil {
		log_utils.Error("%s", err)
		return nil, err
	}
	return b, nil
}

func (u *UserStorage) GetPolicy(_ context.Context, i int64) (*models.Policy, error) {
	stmt, err := u.db.Prepare(SelectPolicyByUserID)
	if err != nil {
		log_utils.Error(err)
	}
	defer storage.CloseSqlResource(stmt)
	p := new(models.Policy)
	var policyjson []byte
	err = stmt.QueryRow(i).Scan(&policyjson)
	if err != nil {
		log_utils.Error(err)
	}
	err = json.Unmarshal(policyjson, p)
	if err != nil {
		log_utils.Error("while while decoding policy json")
		return nil, err
	}
	log_utils.Debug(fmt.Sprintf("%+v", *p))
	return p, nil
}

var _ storage.UserStorage = &UserStorage{}
