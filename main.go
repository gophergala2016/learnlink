package main

import (
	"encoding/json"
	"fmt"
	"github.com/jebrial/learnlink/models"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
)

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

type config struct {
	dbUrl string
}

func main() {
	//load the config
	file, err := os.Open("config.json")
	if err != nil {
		log.Panic(err)
	}
	decoder := json.NewDecoder(file)
	config := config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Panic(err)
	}
	// connect to the database
	db, err2 := models.NewDB(config.dbUrl)
	if err2 != nil {
		log.Panic(err2)
	}

	ctx := context.WithValue(context.Background(), "db", db)

	http.Handle("/users", &ContextAdapter{ctx, ContextHandlerFunc(usersIndex)})
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		panic(err)
	}

}

func usersIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	users, err := models.AllUsers(ctx)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "%d, %s, %s\n", user.Id, user.Name, user.Email)
	}
}
