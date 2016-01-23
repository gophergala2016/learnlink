package main

import (
	//"encoding/json"
	"fmt"
	"github.com/jebrial/learnlink/models"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	//"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	dbUrl string
}

func dbWare() func(goji.Handler) goji.Handler {
	// // //load the config
	// file, err := ioutil.ReadFile("config.json")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// var config Config
	// err = json.Unmarshal(file, &config)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// connect to the database

	db, err2 := models.NewDB("DATABASE URL HERE") //models.NewDB(config.dbUrl)
	if err2 != nil {
		log.Panic(err2)
	}

	return func(h goji.Handler) goji.Handler {
		fn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			ctx = context.WithValue(ctx, "db", db)
			h.ServeHTTPC(ctx, w, r)
		}
		return goji.HandlerFunc(fn)
	}

}

func main() {

	mux := goji.NewMux()
	mux.UseC(dbWare())
	mux.HandleFuncC(pat.Get("/users"), usersIndex)
	err := http.ListenAndServe(":3001", mux)
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
