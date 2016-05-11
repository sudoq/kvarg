package main

import (
	"fmt"
	"github.com/gorilla/mux"
	redis "gopkg.in/redis.v3"
	"math"
	"net/http"
	"os"
	"time"
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

	var err error
	connOk := false
	tries := 0
	maxTries := 5
	for !connOk && tries < maxTries {
		err = redisClient.Ping().Err()
		if err != nil {
			fmt.Println(err)
			tries += 1
			time.Sleep(time.Second * math.Pow10(tries))
			continue
		}
		connOk = true
	}
	os.Exit(1)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", readHandler).Methods("GET")
	r.HandleFunc("/{key}", writeHandler).Methods("PUT")
	fmt.Println("kvarg version 0.0.1")
	http.ListenAndServe(":8080", r)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := redisClient.Get(key).Result()
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(value))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["key"]
	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(http.StatusText(400)))
		return
	}
	if _, ok := r.PostForm["value"]; !ok {
		w.WriteHeader(400)
		w.Write([]byte(http.StatusText(400)))
		return
	}
	value := r.PostForm["value"][0]
	err = redisClient.Set(key, value, 0).Err()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(http.StatusText(200)))
}
