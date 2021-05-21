package repositories

import (
	"app/common"
	"app/datamodels"
	"database/sql"
	"errors"
	"fmt"
)

type IUserRepository interface {
	Conn() error
	Select(string) (*datamodels.User, error)
	Insert(*datamodels.User) (int64, error)
}

type UserRepository struct {
	table string
	db    *sql.DB
}

func (u *UserRepository) Conn() error {
	if u.db == nil {
		mysql, errMysql := common.NewMysqlConn()
		if errMysql != nil {
			return errMysql
		}
		u.db = mysql
		if u.table == "" {
			u.table = "user"
		}
	}
	return nil
}

func (u *UserRepository) Select(userName string) (*datamodels.User, error) {
	if err := u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	sql := "SELECT * FROM " + u.table + " WHERE UserName=?"
	fmt.Println(sql)
	row, errRows := u.db.Query(sql, userName)
	fmt.Println(row)
	defer row.Close()
	if errRows != nil {
		return &datamodels.User{}, errRows
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("用户不存在")
	}
	user := &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return user, nil
}

func (u *UserRepository) Insert(user *datamodels.User) (int64, error) {
	if err := u.Conn(); err != nil {
		return 0, err
	}
	sql := "INSERT " + u.table + " SET NickName=?, UserName=?, Password=?"
	stmt, errStmt := u.db.Prepare(sql)
	if errStmt != nil {
		return 0, errStmt
	}
	result, errRes := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if errRes != nil {
		return 0, errRes
	}
	return result.LastInsertId()
}

func NewUserRepository(table string, db *sql.DB) IUserRepository {
	return &UserRepository{table, db}
}
