package gql

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/play-graphql/common"
	"github.com/play-graphql/model"
	"log"
)

var schema graphql.Schema

func InitSchema() error {
	// Schema
	bookFields := common.ObjToFields(&model.Book{})
	var bookType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "book",
		Description: "book represents book model",
		Fields:      bookFields,
	})

	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "query",
		Description: "query defines query functions",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the book",
						Type:        graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if _, ok := p.Args["id"]; ok {
						id := uint(p.Args["id"].(int))
						book, err := model.BookDAO.Find(context.Background(), id)
						return book, err
					}
					return model.Book{}, nil
				},
			},
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Args: graphql.FieldConfigArgument{
					"ids": &graphql.ArgumentConfig{
						Description: "list id of books",
						Type:        graphql.NewList(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if _, ok := p.Args["ids"]; ok {
						ids := make([]uint, 0)
						for _, id := range p.Args["ids"].([]interface{}) {
							ids = append(ids, uint(id.(int)))
						}
						books, err := model.BookDAO.FilterByIDs(context.Background(), ids)
						log.Println(books)
						return books, err
					}
					return model.BookDAO.FindAll(context.Background())
				},
			},
		},
	})
	schemaConfig := graphql.SchemaConfig{Query: queryType}
	sc, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return err
	}
	schema = sc
	return nil
}
