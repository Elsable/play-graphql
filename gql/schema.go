package gql

import "github.com/graphql-go/graphql"

var schema graphql.Schema

func InitSchema() error {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	sc, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return err
	}
	schema = sc
	return nil
}