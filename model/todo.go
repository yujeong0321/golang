//define todo model
package model

import (
	"github.com/kamva/mgm/v3"
)

type status int

//상수 정의
const (
	UNKNOWN status = iota
	TODO
	DONE
)

func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}

//task 구조체
type Task struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Status           status `json:"status" bson:"status"`
	Deadline         string `json:"deadline" bson:"deadline"`
}

func NewTask(title string, status status, deadline string) *Task {

	return &Task{
		Title:    title,
		Status:   status,
		Deadline: deadline,
	}

}
