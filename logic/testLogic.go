package logic

import (    
	"../api"
	"fmt"
)

type TestModel struct {
	Hello string
}

func Test() *TestModel {
	dbSetting := api.LoadDbSetting()
	fmt.Println(dbSetting)
	return &TestModel{
		Hello: "world",
	}
}