package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// struct for storing data
type Employee struct {
	EmpID   int    `json:"EMP_ID" bson:"EMP_ID"`
	EmpName string `json:"EMP_NAME" bson:"EMP_NAME"`
	Dept    string `json:"EMP_DEPT" bson:"EMP_DEPT"`
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	var body Employee
	params := mux.Vars(r)["EMP_ID"]
	id, _ := strconv.Atoi(params)
	//err := Collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "EMP_ID", Value: intVar}}).Decode(&body)-->one way
	err := Collection.FindOne(context.TODO(), bson.M{"EMP_ID": id}).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(body)
}

func getEmployeeByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var result []Employee
	params := mux.Vars(r)["EMP_NAME"]
	fmt.Print(params)
	cur, err := Collection.Find(context.TODO(), bson.M{"EMP_NAME": params}) //returns a *mongo.Cursor
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor
		var elem Employee
		err := cur.Decode(&elem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		result = append(result, elem) //appending document pointed by Next()
	}
	cur.Close(context.TODO()) //closes the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(result)

}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var results []Employee
	//collection := Db().Database("Employee").Collection("Employee")
	cur, err := Collection.Find(context.TODO(), bson.D{}) //returns a *mongo.Cursor
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor
		var elem Employee
		err := cur.Decode(&elem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		results = append(results, elem) //appending document pointed by Next()
	}
	cur.Close(context.TODO()) //closes the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

// var client *mongo.Client

func createEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	//collection := Db().Database("Employee").Collection("Employee")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := Collection.InsertOne(ctx, employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Println("One record inserted successfully...")
	json.NewEncoder(w).Encode("Record Inserted Status Code:")
	json.NewEncoder(w).Encode(http.StatusOK)

	//json.NewEncoder(w).Encode(result.InsertedID)
	//json.NewEncoder(w).SetEscapeHTML(true)
	//log.Print("inserted")
	//return employee

}
