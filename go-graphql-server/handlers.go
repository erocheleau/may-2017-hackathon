package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/graphql-go/graphql"
)

type QueryPost struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("POST/GET on /graphql"))
}

func GraphQLGETHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("totlo")
	query, err := url.ParseQuery(r.URL.RawQuery)
	fmt.Println("totlo2")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("Error parsing query: ", err.Error())
		return
	}
	fmt.Println("totlo3")
	w.Header().Set("Content-Type", "application/json; chartype: utf-8")
	w.WriteHeader(200)
	for _, value := range query {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: value[0],
		})

		if result.HasErrors() {
			b, err := json.Marshal(result.Errors)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprint(w, string(b))
		} else {
			b, err := json.Marshal(result.Data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprint(w, string(b))
		}
	}
}
func GraphQLPOSTHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var queryPost QueryPost
	err := decoder.Decode(&queryPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("Error parsing query: ", err.Error())
		return
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; chartype: utf-8")
	w.WriteHeader(200)
	fmt.Println(queryPost.Query)
	fmt.Println(queryPost.OperationName)
	fmt.Println(queryPost.Variables)

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  queryPost.Query,
		OperationName:  queryPost.OperationName,
		VariableValues: queryPost.Variables,
	})

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))

}
