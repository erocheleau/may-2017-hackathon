package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func getQueryType() *graphql.Object {
	resultInterface := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Result",
		Description: "A search result",
		Fields: graphql.Fields{
			"UniqueId": &graphql.Field{
				Type:        graphql.ID,
				Description: "The result UniqueId",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "The result title",
			},
		},
	})

	queryResultObject := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Result",
		Description: "A search result",
		Fields: graphql.Fields{
			"totalCount": &graphql.Field{
				Type:        graphql.Int,
				Description: "Totalcount",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(QueryResponse).TotalCount, nil
				},
			},

			"results": &graphql.Field{
				Type:        graphql.NewList(resultInterface),
				Description: "The list of results",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					b, err := json.Marshal(p.Source.(QueryResponse))
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(string(b))

					return p.Source.(QueryResponse).Results, nil
				},
			},
		},
	})

	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "QueryResponse",
		Description: "The query response",
		Fields: graphql.Fields{
			"queryResults": &graphql.Field{
				Type:        queryResultObject,
				Description: "The query result",
				Args: graphql.FieldConfigArgument{
					"q": &graphql.ArgumentConfig{
						Description: "Query Expression",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return DoQuery(p.Args["q"].(string)), nil
				},
			},
		},
	})
}
