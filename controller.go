package example

import (
	"fmt"
	page "github.com/ikaiguang/go-pagination"
)

type UserController struct {
	Model *UserModel
}

// get user list
func (c *UserController) List(option *page.PagingOption) ([]*UserModel, *page.PagingResult, error) {

	// get paging query option collection
	pagingOptionCollection, err := page.GetOptionCollection(option, c.Model.NewModel())
	if err != nil {
		err := fmt.Errorf("UserController.List GetPagingOptionCollection error : %v", err)
		return nil, nil, err
	}

	// list
	list, count, err := c.Model.List([]*WhereCondition{}, pagingOptionCollection)
	if err != nil {
		err := fmt.Errorf("UserController.List model.List error : %v", err)
		return nil, nil, err
	}

	// set paging result
	pagingResultCollection := &page.PagingResultCollection{
		TotalRecords: count,
		ResultSlice:  list,
	}
	pagingResult, err := page.SetPagingResult(pagingOptionCollection, pagingResultCollection)
	if err != nil {
		err := fmt.Errorf("UserController.List SetPagingResult error : %v", err)
		return nil, nil, err
	}
	return list, pagingResult, nil
}
