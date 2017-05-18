package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/graphql-go/graphql"
)

type QueryPost struct {
	Query string `json:"query"`
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

	w.Header().Set("Content-Type", "application/json; chartype: utf-8")
	w.WriteHeader(200)
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

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: queryPost.Query,
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
