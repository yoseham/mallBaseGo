package repositories

import (
	"app/common"
	"app/datamodels"
	"database/sql"
	"fmt"
	"strconv"
)

type IOrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(int64) (bool, error)
	Update(order *datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func NewOrderManagerRepository(table string, db *sql.DB) IOrderRepository {
	return &OrderManagerRepository{table, db}
}

func (o *OrderManagerRepository) Conn() (err error) {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
		if o.table == "" {
			o.table = "order"
		}
	}
	return
}

func (o *OrderManagerRepository) Insert(order *datamodels.Order) (orderID int64, err error) {
	if err := o.Conn(); err != nil {
		return 0, err
	}
	//准备sql
	sql := "INSERT " + o.table + " SET UserID=?, ProductID=?, OrderStatus"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return 0, err
	}
	//传入参数
	result, errRes := stmt.Exec(order.UserID, order.ProductID, order.OrderStatus)
	if errRes != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (o *OrderManagerRepository) Delete(orderID int64) (result bool, err error) {
	if err := o.Conn(); err != nil {
		return false, err
	}
	sql := "DELETE FROM " + o.table + " WHERE ID=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(orderID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *OrderManagerRepository) Update(order *datamodels.Order) (err error) {
	if err := o.Conn(); err != nil {
		return err
	}
	sql := "UPDATE " + o.table + " SET UserID=?, ProductID=?, OrderStatus=? here ID=" + strconv.FormatInt(order.ID, 10)
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return err
	}
	_, errRes := stmt.Exec(order.UserID, order.ProductID, order.OrderStatus)
	if errRes != nil {
		return err
	}
	return
}

func (o *OrderManagerRepository) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT * FROM " + o.table + " WHERE ID=" + strconv.FormatInt(orderID, 10)
	row, errRow := o.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return &datamodels.Order{}, errRow
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Order{}, nil
	}
	order = &datamodels.Order{}
	common.DataToStructByTagSql(result, order)
	return
}

func (o *OrderManagerRepository) SelectAll() (orders []*datamodels.Order, err error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT * FROM " + o.table
	rows, errRow := o.mysqlConn.Query(sql)
	defer rows.Close()
	if errRow != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orders = append(orders, order)
	}
	return
}

func (o *OrderManagerRepository) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT o.ID, p.ProductName, o.OrderStatus FROM imooc.order as o left join imooc.product as p on o.ProductID = p.ID"
	rows, errRow := o.mysqlConn.Query(sql)
	defer rows.Close()
	if errRow != nil {
		fmt.Println(errRow)
		return nil, err
	}
	orderMap = common.GetResultRows(rows)
	return
}
