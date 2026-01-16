package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	collection = client.Database("cinema").Collection("movies")
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	ht
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Movies service listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	var movies []map[string]interface{}
	cursor, _ := collection.Find(context.TODO(), map[string]interface{}{})
	cursor.All(context.TODO(), &movies)
	json.NewEncoder(w).Encode(movies)
}
