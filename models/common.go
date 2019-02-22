package models

import (
	"fmt"
	"github.com/ikaiguang/go-pagination"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbPool *gorm.DB

func init()  {
	var err error
	// dbPool conn
	dbPool, err = newDbConnection()
	if err != nil {
		panic(fmt.Errorf("init newDbConnection error : %v", err))
	}
	//defer dbPool.Close()
}

// newDbConnection new dbPool connection
func newDbConnection() (*gorm.DB, error) {

	dbDiver := "mysql"
	dbDsn := "root:Mysql.123456@tcp(127.0.0.1:3306)/test?"
	dbDsn += "charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(dbDiver, dbDsn)

	if err != nil {
		err = fmt.Errorf("gorm.Open dsn error : %v", err)
		return nil, err
	}
	//defer dbPool.Close()

	db.LogMode(true)

	return db, nil
}

// WhereCondition where condition
type WhereCondition struct {
	pagination.PagingWhere
}

// WhereConditions where conditions
func WhereConditions(dbConn *gorm.DB, whereConditions []*WhereCondition) *gorm.DB {

	// where
	for _, where := range whereConditions {
		// dbPool.Where("id = ?", id)
		// dbPool.Where("id in (?)", ids)
		whereStr := fmt.Sprintf("%s %s %s", where.Column, where.Symbol, where.Placeholder)
		dbConn = dbConn.Where(whereStr, where.Data)
	}
	return dbConn
}

// Pagination pagination
func Pagination(dbConn *gorm.DB, pagingOptionCollection *pagination.PagingOptionCollection) *gorm.DB {
	// limit offset
	dbConn = dbConn.Limit(pagingOptionCollection.Limit).Offset(pagingOptionCollection.Offset)

	// where
	for _, where := range pagingOptionCollection.Where {
		// dbPool.Where("id = ?", id)
		whereStr := fmt.Sprintf("%s %s %s", where.Column, where.Symbol, where.Placeholder)
		dbConn = dbConn.Where(whereStr, where.Data)
	}

	// order
	for _, order := range pagingOptionCollection.Order {
		dbConn = dbConn.Order(fmt.Sprintf("%s %s", order.Column, order.Direction))
	}
	return dbConn
}
