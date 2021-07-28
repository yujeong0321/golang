//setup mongodb default config
package dao

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "golang/model"
)

func init() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "local", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("DAO Error: ", err)
	}
}

func InsertTask(task model.Task) {
	newtask := model.NewTask(task.Title, task.Status, task.Deadline)

	err := mgm.Coll(newtask).Create(newtask)

	if err != nil {
		fmt.Println("InsertTaskErr", err)
	}

}
