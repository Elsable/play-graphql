package gql

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
	"log"
)


func executeQuery(query string) (*graphql.Result, error) {
	log.Printf("incoming: %v", query)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		return nil, result.Errors[0]
	}
	return result, nil
}

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	rlt, err := executeQuery(r.URL.Query().Get("query"))
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	rJSON, _ := json.Marshal(rlt)
	log.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
	json.NewEncoder(w).Encode(rlt)
}

