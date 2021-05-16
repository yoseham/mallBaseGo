package repositories

import (
	"app/common"
	"app/datamodels"
	"database/sql"
	"strconv"
)

//第一步开发接口
//第二步实现接口
type IProduct interface {
	//连接数据库
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) (bool, error)
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

//通过创建初始化函数验证接口
func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table, db}
}

//数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := sql.Open("mysql", "root:07597321@tcp(106.54.91.157:3306)/imooc?charset=utf8")
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
		if p.table == "" {
			p.table = "product"
		}
	}
	return
}

//添加
func (p *ProductManager) Insert(product *datamodels.Product) (id int64, err error) {
	//准备sql
	sql := "INSERT product SET productName=?, productNum=?, productImage=?, productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	//传入参数
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

//删除
func (p *ProductManager) Delete(productID int64) (result bool, err error) {
	sql := "DELETE FROM product WHERE ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(productID)
	if err != nil {
		return false, err
	}
	return true, nil
}

//更新
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	sql := "UPDATE product set productName=? productNum=? productImage=? productUrl=? where ID=?" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductUrl, product.ProductImage)
	if err != nil {
		return err
	}
	return
}

//根据商品ID查询商品
func (p *ProductManager) SelectByKey(productID int64) (product *datamodels.Product, err error) {
	sql := "SELECT * FROM product WHERE ID=" + strconv.FormatInt(product.ID, 10)
	row, errRow := p.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return &datamodels.Product{}, errRow
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	common.DataToStructByTagSql(result, product)
	return
}

//获取所有商品
func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	sql := "SELECT * FROM product"
	rows, errRow := p.mysqlConn.Query(sql)
	defer rows.Close()
	if errRow != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		products = append(products, product)
	}
	return
}
