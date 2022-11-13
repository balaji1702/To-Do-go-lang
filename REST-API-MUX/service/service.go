package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "rest-api-mux/db_config"
	"rest-api-mux/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection = db.Collection

// Get_task godoc
// @Summary Get_all task
// @Description get all task
// @Tags task
// @Accept  json
// @Produce  json
// @router /get-task [get]
func Get_task(w http.ResponseWriter, r *http.Request) {
	log.Print("get_task is  being performed")

	ctx := context.Background()
	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}})
	result, _ := collection.Find(ctx, bson.M{}, opts)
	var res []bson.M
	if err := result.All(ctx, &res); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(res)

}

// Create_task godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags task
// @Accept  json
// @Produce  json
// @Param task body models.To_do true "Create task"
// @router /create-task [POST]
func Create_task(w http.ResponseWriter, r *http.Request) {
	log.Print("creating a task")

	var input models.To_do
	_ = json.NewDecoder(r.Body).Decode(&input)
	ctx := context.Background()
	_, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"msg": "task done sucessfully"})

}

// Update_task godoc
// @Summary update a task status
// @Description update a task status based on task name
// @Tags task
// @Accept  json
// @Produce  json
// @Param  task_name path string true "task_name"
// @Param q query string true "update task" Enums("COMPLETED","PENDING")
// @router /update-task [PUT]
func Update_task(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tname := vars["task_name"]
	query := r.URL.Query().Get("status")
	collection := db.Collection
	filter := bson.M{"task_name": tname}

	update := bson.M{"$set": bson.M{"status": query}}

	log.Print("updating  a task")
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"msg": "updated sucedfully"})

}

// Delete_task godoc
// @Summary delete a task
// @Description delete a task with the taskname
// @Tags task
// @Accept  json
// @Produce  json
// @Param  task_name path string true "task_name"
// @router /update-task [DELETE]
func Delete_task(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	collection := db.Collection
	filter := bson.M{"task_name": vars["task_name"]}
	log.Print("task deletion")
	collection.DeleteOne(context.Background(), filter)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"msg": "deleted sucedfully"})

}
