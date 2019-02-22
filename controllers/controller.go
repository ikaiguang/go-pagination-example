package controllers

import (
	"fmt"
	"github.com/ikaiguang/go-pagination"
	"github.com/ikaiguang/go-pagination-example/models"
	pb "github.com/ikaiguang/go-pagination-example/protobuf"
)

// UserController user controller
type UserController struct {
	model models.User
}

// List get user list
func (c *UserController) List(in *pb.UserListReq) (*pb.UserListResp, error) {

	// validate something
	if in.ActionUserID <= 0 {
		panic("ActionUserID <= 0")
	}

	// paging option
	pagingOptionCollection, err := pagination.GetOptionCollection(in.PagingOption, c.model.NewModel())
	if err != nil {
		return nil, fmt.Errorf("User.List pagination.GetOptionCollection error : %v", err)
	}

	// 获取
	dataModels, recordCount, err := c.model.List(in, pagingOptionCollection)
	if err != nil {
		return nil, fmt.Errorf("*UserController.List error : %v", err)
	}

	// paging result
	pagingResultCollection := &pagination.PagingResultCollection{
		TotalRecords: recordCount,
		ResultSlice:  &dataModels,
	}
	pagingResult, err := pagination.SetPagingResult(pagingOptionCollection, pagingResultCollection)
	if err != nil {
		return nil, fmt.Errorf("User.List pagination.SetPagingResult error : %v", err)
	}

	// proto
	dataSlice := c.setProtocols(dataModels)

	return &pb.UserListResp{Data: dataSlice, PagingResult: pagingResult}, nil
}

// setProtocols : set protocols
func (c *UserController) setProtocols(dataModels []*models.User) []*pb.User {

	var protocols = make([]*pb.User, len(dataModels))

	for index := range dataModels {
		protocols[index] = c.setProtocol(dataModels[index])
	}
	return protocols
}

// setProtocol : set protocol
func (c *UserController) setProtocol(dataModel *models.User) *pb.User {
	return &pb.User{
		Id:   dataModel.Id,
		Name: dataModel.Name,
		Age:  dataModel.Age,
	}
}
