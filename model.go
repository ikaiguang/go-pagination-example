package example

import (
	"fmt"
	page "github.com/ikaiguang/go-pagination"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// new db connection
func newDbConnection() (*gorm.DB, error) {

	dbDiver := "mysql"
	dbDsn := "root:Mysql.123456@tcp(127.0.0.1:3306)/test?"
	dbDsn += "charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(dbDiver, dbDsn)

	if err != nil {
		err = fmt.Errorf("gorm.Open dsn error : %v", err)
		return nil, err
	}
	//defer db.Close()

	db.LogMode(true)

	return db, nil
}

// where condition
type WhereCondition struct {
	page.PagingWhere
}

// where conditions
func WhereConditions(dbConn *gorm.DB, whereConditions []*WhereCondition) *gorm.DB {

	// where
	for _, where := range whereConditions {
		// db.Where("id = ?", id)
		whereStr := fmt.Sprintf("%s %s %s", where.Column, where.Symbol, where.Placeholder)
		dbConn = dbConn.Where(whereStr, where.Data)
	}
	return dbConn
}

// pagination
func Pagination(dbConn *gorm.DB, pagingOptionCollection *page.PagingOptionCollection) *gorm.DB {
	// limit offset
	dbConn = dbConn.Limit(pagingOptionCollection.Limit).Offset(pagingOptionCollection.Offset)

	// where
	for _, where := range pagingOptionCollection.Where {
		// db.Where("id = ?", id)
		whereStr := fmt.Sprintf("%s %s %s", where.Column, where.Symbol, where.Placeholder)
		dbConn = dbConn.Where(whereStr, where.Data)
	}

	// order
	for _, order := range pagingOptionCollection.Order {
		dbConn = dbConn.Order(fmt.Sprintf("%s %s", order.Column, order.Direction))
	}
	return dbConn
}

// user model
type UserModel struct {
	Id   int64  `gorm:"PRIMARY_KEY;COLUMN:id"` // id
	Name string `gorm:"COLUMN:name"`           // name
	Age  int64  `gorm:"COLUMN:age"`            // age
}

// user table name
func (m *UserModel) TableName() string {
	return "pagination_users"
}

// new model
func (m *UserModel) NewModel() *UserModel {
	return new(UserModel)
}

// user model list
func (m *UserModel) List(whereConditions []*WhereCondition, pagingOptionCollection *page.PagingOptionCollection) (*[]UserModel, int64, error) {
	var count int64
	var list []UserModel

	// db conn
	db, err := newDbConnection()
	if err != nil {
		err = fmt.Errorf("UserModel.List newDbConnection error : %v", err)
		return &list, count, err
	}
	defer db.Close()

	// query where
	userDb := WhereConditions(db, whereConditions)
	defer userDb.Close()

	if err := userDb.Table(m.TableName()).Count(&count).Error; err != nil {
		err = fmt.Errorf("UserModel.List Count error : %v", err)
		return &list, count, err
	} else if count == 0 {
		return &list, count, err // empty
	}

	// pagination
	if err := Pagination(userDb, pagingOptionCollection).Find(&list).Error; err != nil {
		err = fmt.Errorf("UserModel.List Find error : %v", err)
		return &list, count, err
	}
	return &list, count, err
}
