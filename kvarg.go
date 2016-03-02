package main

import (
	"fmt"
	"github.com/gorilla/mux"
	redis "gopkg.in/redis.v3"
	"net/http"
	"os"
)

var redisClient *redis.Client

func init() {
	addr := os.Getenv("DB_PORT_6379_TCP_ADDR")
	port := os.Getenv("DB_PORT_6379_TCP_PORT")
	if addr == "" {
		addr = "localhost"
	}
	if port == "" {
		port = "6379"
	}
	addrPort := fmt.Sprintf("%s:%s", addr, port)

	redisClient = redis.NewClient(&redis.Options{
		Addr:     addrPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisClient.Ping().Err()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", readHandler)
	r.HandleFunc("/{key}/{value}", writeHandler)
	fmt.Println("kvarg version 0.0.1")
	http.ListenAndServe(":8080", r)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := redisClient.Get(key).Result()
	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(value))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
