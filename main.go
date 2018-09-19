package main

import (
	"context"
	"github.com/globalsign/mgo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/play-graphql/gql"
	"github.com/play-graphql/model"
	"time"
)

// init prepare all resource
func init() {
	log.SetFlags(log.Lshortfile)
	s, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("could not dial mongodb: %v", err)
	}
	if err := s.Login(&mgo.Credential{Username: "root", Password: "123456", Source: "admin"}); err != nil {
		log.Fatalf("could not login mongodb: %v", err)
	}
	db := s.DB("graphql")
	model.InitBookDAO(db)
	log.Printf("init mongodb resource ==> done!")
	if err := gql.InitSchema(); err != nil {
		log.Fatalf("could not init graphql schema: %v", err)
	}
	log.Printf("init graphql schema ==> done!")
}

func main() {
	defer model.BookDAO.Close()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", gql.GraphQLHandler)
	srv := &http.Server{
		Addr:    "0.0.0.0:9099",
		Handler: mux,
	}
	go func() {
		log.Printf("listen on port :9099\n")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("could not listen and serve on 8080: %v\n", err)
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("server down")
	return
}

