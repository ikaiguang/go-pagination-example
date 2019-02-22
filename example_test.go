package example

import (
	"github.com/ikaiguang/go-pagination"
	"github.com/ikaiguang/go-pagination-example/controllers"
	pb "github.com/ikaiguang/go-pagination-example/protobuf"
	"testing"
)

// generate test data
func TestGenerateTestData(t *testing.T) {
	return // continue

	//err := models.GenerateTestData()
	//if err != nil {
	//	t.Errorf("testing : models.GenerateTestData error : %v", err)
	//	return
	//}
}

// controller list
func TestUserController_List(t *testing.T) {

	// paging mode : page number
	testPageNumberMode(t)

	// paging mode : cursor mode && asc(order by id asc)
	testCursorModeAsc(t)

	// paging mode : cursor mode && desc(order by id desc)
	testCursorModeDesc(t)
}

// paging mode : cursor mode && desc(order by id desc)
func testCursorModeDesc(t *testing.T) {

	var controller controllers.UserController // controller
	var req *pb.UserListReq                   // list req
	var resp *pb.UserListResp                 // list resp
	var err error                             // error
	var cursorOption *pagination.PagingOption // option

	// init cursor option
	cursorOption = pagination.DefaultPagingOption()

	// init req
	req = &pb.UserListReq{ActionUserID: 1, PagingOption: cursorOption}

	// ===== paging mode : cursor : order by id desc ===== //
	// ===== goto 3rd page ===== //
	// cursor option
	cursorOption.PagingMode = pagination.PagingModeCursor // cursor mode
	cursorOption.PageSize = 2                             // page size : 2
	cursorOption.GotoPageNumber = 3                       // goto 3th page
	cursorOption.CursorColumn = "id"                      // order by id desc
	cursorOption.CursorDirection = "desc"                 // order by id desc

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n cursor : order by id desc && goto 3th page \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}

	// ===== paging mode : cursor : order by id desc ===== //
	// ===== goto 1st page(preceding page) ===== //
	// ===== Skip from page 3 to page 1 ===== //
	// cursor option
	cursorOption.CurrentPageNumber = cursorOption.GotoPageNumber // current page number
	cursorOption.GotoPageNumber = 1                              // goto 1st page
	cursorOption.CursorValue = resp.PagingResult.CursorValue     // cursor value

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n cursor : order by id desc && goto 1st page(Skip from page 3 to page 1) \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}

	// ===== paging mode : cursor : order by id desc ===== //
	// ===== goto 5th page(next page) ===== //
	// ===== Skip from page 1 to page 5 ===== //
	// cursor option
	cursorOption.CurrentPageNumber = cursorOption.GotoPageNumber // current page number
	cursorOption.GotoPageNumber = 5                              // goto 5th page
	cursorOption.CursorValue = resp.PagingResult.CursorValue     // cursor value

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n cursor : order by id desc && goto 5th page(Skip from page 1 to page 5) \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}
}

// paging mode : cursor mode && asc(order by id asc)
func testCursorModeAsc(t *testing.T) {

	var controller controllers.UserController // controller
	var req *pb.UserListReq                   // list req
	var resp *pb.UserListResp                 // list resp
	var err error                             // error
	var cursorOption *pagination.PagingOption // option

	// init cursor option
	cursorOption = pagination.DefaultPagingOption()

	// init req
	req = &pb.UserListReq{ActionUserID: 1, PagingOption: cursorOption}

	// ===== paging mode : cursor : order by id asc ===== //
	// ===== goto 3rd page ===== //
	// cursor option
	cursorOption = pagination.DefaultPagingOption()
	cursorOption.PagingMode = pagination.PagingModeCursor // cursor mode
	cursorOption.PageSize = 2                             // page size : 2
	cursorOption.GotoPageNumber = 3                       // goto 3th page
	cursorOption.CursorColumn = "id"                      // order by id asc
	cursorOption.CursorDirection = "asc"                  // order by id asc

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n cursor : order by id asc && goto 3th page \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}

	// ===== paging mode : cursor : order by id asc ===== //
	// ===== goto 1st page(preceding page) ===== //
	// ===== Skip from page 3 to page 1 ===== //
	// cursor option
	cursorOption.CurrentPageNumber = cursorOption.GotoPageNumber // current page number
	cursorOption.GotoPageNumber = 1                              // goto 1st page
	cursorOption.CursorValue = resp.PagingResult.CursorValue     // cursor value

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n cursor : order by id asc && goto 1st page(Skip from page 3 to page 1) \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}

	// ===== paging mode : cursor : order by id asc ===== //
	// ===== goto 5th page(next page) ===== //
	// ===== Skip from page 1 to page 5 ===== //
	// cursor option
	cursorOption.CurrentPageNumber = cursorOption.GotoPageNumber // current page number
	cursorOption.GotoPageNumber = 5                              // goto 5th page
	cursorOption.CursorValue = resp.PagingResult.CursorValue     // cursor value

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n cursor : order by id asc && goto 5th page(Skip from page 1 to page 5) \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}
}

// paging mode : page number
func testPageNumberMode(t *testing.T) {

	var controller controllers.UserController     // controller
	var req *pb.UserListReq                       // list req
	var resp *pb.UserListResp                     // list resp
	var err error                                 // error
	var pageNumberOption *pagination.PagingOption // option

	// init page number option
	pageNumberOption = pagination.DefaultPagingOption()

	// init req
	req = &pb.UserListReq{ActionUserID: 1, PagingOption: pageNumberOption}

	// ===== paging mode : page number ===== //
	// ===== goto 3rd page ===== //
	// page number option
	pageNumberOption = pagination.DefaultPagingOption()
	pageNumberOption.PageSize = 2                                        // page size : 2
	orderBy := &pagination.PagingOrder{Column: "age", Direction: "desc"} // order by age desc
	pageNumberOption.OrderBy = append(pageNumberOption.OrderBy, orderBy) // order by age desc
	pageNumberOption.GotoPageNumber = 3                                  // goto 3rd page

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n page number : order age desc && goto 3rd page \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}

	// ===== paging mode : page number ===== //
	// ===== goto 1st page(preceding page) ===== //
	pageNumberOption.GotoPageNumber = 1 // goto 1st page(next page)

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n page number : order age desc && goto 1st page \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}

	// ===== paging mode : page number ===== //
	// ===== goto 5th page(next page) ===== //
	pageNumberOption.GotoPageNumber = 5 // goto 5th page(next page)

	resp, err = controller.List(req)
	if err != nil {
		t.Errorf("testing : controller.List error : %v", err)
		return
	} else {
		format := "\n page number : order age desc && goto 5th page \n"
		format += "\n paging result : %v \n"
		format += "\n list : %v \n"

		t.Logf(format, resp.PagingResult, resp.Data)
	}
}
