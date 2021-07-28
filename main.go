package main

import (
	"golang/dao"
	model "golang/model"
)

func main() {

	task1 := model.NewTask("task1", model.TODO, "20210730")
	dao.InsertTask(*task1)

}
