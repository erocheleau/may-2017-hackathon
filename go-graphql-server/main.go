package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"net/url"

	"github.com/graphql-go/graphql"
	"github.com/husobee/vestigo"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: getQueryType(),
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
	w.Header().Set("Content-type", "application/json")
	for _, value := range query {
		//fmt.Fprintf(w, "Query: %s => %s ", key, value)
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: value[0],
		})

		if result.Errors != nil {
			fmt.Fprintf(w, "%v", result.Errors)
		}

		b, err := json.Marshal(result.Data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(b))
	}
}
func GraphQLPOSTHandler(w http.ResponseWriter, r *http.Request) {

}
