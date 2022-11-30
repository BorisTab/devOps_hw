package main

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("mydb")
	collection, err := getCollection(db, "test")
	if err != nil {
		panic(err)
	}

	runServer(collection)
}

func getCollection(db *mongo.Database, name string) (*mongo.Collection, error) {
	names, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	for _, n := range names {
		if n == name {
			return db.Collection(name), nil
		}
	}

	db.CreateCollection(context.TODO(), name)
	return db.Collection(name), nil
}

func runServer(collection *mongo.Collection) {
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")

		if key == "" || value == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := collection.InsertOne(context.TODO(), bson.M{"key": key, "value": value})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var result bson.M
		err := collection.FindOne(context.TODO(), bson.M{"key": key}).Decode(&result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result["value"].(string)))
	})

	http.HandleFunc("/delete_all", func(w http.ResponseWriter, r *http.Request) {
		_, err := collection.DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
