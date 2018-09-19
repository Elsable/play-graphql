package model

import (
	"context"
	"github.com/globalsign/mgo"
	"log"
	"testing"
)

var (
	ctx     = context.Background()
	mgoURL = "localhost:27017"
)

func TestBookDAO_Find(t *testing.T) {
	s, err := mgo.Dial(mgoURL)
	if err != nil {
		t.Fatalf("could not dial mongodb: %v", err)
	}
	if err := s.Login(&mgo.Credential{Username:"root", Password:"123456", Source:"admin"}); err != nil {
		t.Fatalf("could not login to mongodb: %v", err)
	}
	db := s.DB("graphql")
	BookDAO = NewBookDAO(db)
	book, err := BookDAO.Find(ctx, 1)
	if err != nil {
		t.Fatalf("could not find: %v", err)
	}
	log.Println(book.Title)
}

func TestBookDAO_FilterByStatus(t *testing.T) {
	s, err := mgo.Dial(mgoURL)
	if err != nil {
		t.Fatalf("could not dial mongodb: %v", err)
	}
	if err := s.Login(&mgo.Credential{Username:"root", Password:"123456", Source:"admin"}); err != nil {
		t.Fatalf("could not login to mongodb: %v", err)
	}
	db := s.DB("graphql")
	bookDAO := NewBookDAO(db)
	books, err := bookDAO.FilterByStatus(ctx, "PUBLISH")
	if err != nil {
		t.Fatalf("could not filter by status: %v", err)
	}
	log.Println(len(books))
}

func TestBookDAO_FindAll(t *testing.T) {
	s, err := mgo.Dial(mgoURL)
	if err != nil {
		t.Fatalf("could not dial mongodb: %v", err)
	}
	if err := s.Login(&mgo.Credential{Username:"root", Password:"123456", Source:"admin"}); err != nil {
		t.Fatalf("could not login to mongodb: %v", err)
	}
	db := s.DB("graphql")
	BookDAO = NewBookDAO(db)
	books, err := BookDAO.FindAll(ctx)
	if err != nil {
		t.Fatalf("could not filter by status: %v", err)
	}
	log.Println(len(books))
	//for _, book := range books {
	//	if book.ID == 0 {
	//		log.Println(book.Title)
	//		log.Println(book.ISBN)
	//		if err := book.Destroy(); err != nil {
	//			t.Fatalf("could not destroy book: %v", err)
	//		}
	//	}
	//}
}

func TestBookDAO_FilterByIDs(t *testing.T) {
	s, err := mgo.Dial(mgoURL)
	if err != nil {
		t.Fatalf("could not dial mongodb: %v", err)
	}
	if err := s.Login(&mgo.Credential{Username:"root", Password:"123456", Source:"admin"}); err != nil {
		t.Fatalf("could not login to mongodb: %v", err)
	}
	db := s.DB("graphql")
	bookDAO := NewBookDAO(db)
	books, err := bookDAO.FilterByIDs(ctx, []uint{0,})
	if err != nil {
		t.Fatalf("could not filter by status: %v", err)
	}
	log.Println(len(books))
}
