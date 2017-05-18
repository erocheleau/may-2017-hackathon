package main

import (
	"fmt"
	"log"
	"net/http"

	"net/url"

	"github.com/graphql-go/graphql"
	"github.com/husobee/vestigo"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user

/*
   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func main() {

	// Router
	router := vestigo.NewRouter()
	vestigo.AllowTrace = true //Trace is on

	// Handlers
	router.HandleFunc("/", Index)

	router.Get("/graphql", GraphQLGETHandler)
	router.Post("/graphql", GraphQLPOSTHandler)

	// Serve
	log.Fatal(http.ListenAndServe(":8080", router))
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("POST/GET on /graphql"))
}

func GraphQLGETHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("Error parsing query: ", err.Error())
		return
	}

	w.WriteHeader(200)
	for key, value := range query {
		fmt.Fprintf(w, "Query: %s => %s ", key, value)
	}
}
func GraphQLPOSTHandler(w http.ResponseWriter, r *http.Request) {

}
