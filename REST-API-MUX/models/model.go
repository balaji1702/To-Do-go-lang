package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type To_do struct {
	Task_name string `json:"task_name" bson:"task_name,omitempty"`
	Desc      string `json:"description" bson:"description,omitempty"`
	Priority  string `json:"priority" bson:"priority,omitempty"`
	Status    string `json:"status" bson:"status,omitempty"`
}
