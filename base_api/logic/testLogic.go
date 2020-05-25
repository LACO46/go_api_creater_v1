package logic

import (
	"../api"
)

type TestItemModel struct {
	Id string
	Name string
}

func Test() []TestItemModel {
	dbSetting := api.LoadDbSetting()
	db, err := api.ConnectDb(dbSetting)
	if err != nil {
		return nil
	}

	var testListModel []TestItemModel
	for _, item := range api.Test(db) {
		testItem := &TestItemModel {
			Id: item.Id,
			Name: item.Name,
		}

		testListModel = append(testListModel, *testItem)
	}
	return testListModel
}