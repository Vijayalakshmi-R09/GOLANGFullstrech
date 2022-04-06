package main

import (
	"API_DB/app"
)

func main() {
	app.Start()

	// usersCollection := client.Database("Employee").Collection("Employee")
	// //dropps all the existing documents in the collections and makes empty
	// if err = usersCollection.Drop(context.TODO()); err != nil {
	// 	panic(err)
	// }
}
