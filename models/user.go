package models

import (
	"fmt"
	"github.com/ikaiguang/go-pagination"
	pb "github.com/ikaiguang/go-pagination-example/protobuf"
	"github.com/jinzhu/gorm"
)

// User user model
type User struct {
	Id   int64  `gorm:"PRIMARY_KEY;COLUMN:id"` // id
	Name string `gorm:"COLUMN:name"`           // name
	Age  int64  `gorm:"COLUMN:age"`            // age
}

// TableName user table name
func (m *User) TableName() string {
	return "pagination_users"
}

// NewModel new model
func (m *User) NewModel() *User {
	return new(User)
}

// List user model list
func (m *User) List(in *pb.UserListReq, pagingOptionCollection *pagination.PagingOptionCollection) ([]*User, int64, error) {
	var recordCount int64
	var resultSlice []*User

	// where
	dbConn := m.WhereConditions(dbPool, in)
	//defer dbConn.Close()

	// count
	if err := dbConn.Table(m.TableName()).Count(&recordCount).Error; err != nil {
		return resultSlice, recordCount, fmt.Errorf("User.List Count error : %v", err)
	} else if recordCount == 0 {
		return resultSlice, recordCount, err // empty
	}

	// pagination
	if err := Pagination(dbConn, pagingOptionCollection).Find(&resultSlice).Error; err != nil {
		return resultSlice, recordCount, fmt.Errorf("User.List Find error : %v", err)
	}
	return resultSlice, recordCount, nil
}

// WhereConditions orm where
func (m *User) WhereConditions(dbConn *gorm.DB, in *pb.UserListReq) *gorm.DB {

	// where
	if in.UserID > 0 {
		dbConn = dbConn.Where("id = ?", in.UserID)
	}
	return dbConn
}
