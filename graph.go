package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/antonio-nirina/monitore-sp-api/graphql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main(){
	schemaConfig := graphql.SchemaConfig{
	  Query: graphql.NewObject(graphql.ObjectConfig{
	    Name:   "RootQuery",
	    Fields: queries.GetRootFields(),
	  }),
	  Mutation: graphql.NewObject(graphql.ObjectConfig{
	    Name:   "RootMutation",
	    Fields: mutations.GetRootFields(),
	  }),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
	  log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
	  	Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/", httpHandler)
	fmt.Println("ready: listening...\n")
	http.ListenAndServe(":8080", nil)
}