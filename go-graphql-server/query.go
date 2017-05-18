package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type QueryBuilder Query

type Query struct {
	q               string
	fieldsToInclude []string
}

type QueryResponse struct {
	Results    []QueryResponseResult `json: "results"`
	TotalCount int                   `json: "totalCount"`
}

type QueryResponseResult struct {
	Title    string                 `json: "title"`
	UniqueId string                 `json: "UniqueId"`
	Raw      map[string]interface{} `json: "raw"`
}

type CoveoField struct {
	Name        string `json: "name"`
	Description string `json: "Description"`
}

type CoveoFields struct {
	Fields []CoveoField `json: "fields"`
}

func (qb *QueryBuilder) addField(field string) {
	for _, a := range qb.fieldsToInclude {
		if a == field {
			return
		}
	}

	qb.fieldsToInclude = append(qb.fieldsToInclude, field)
}

func DoQuery(q string) QueryResponse {
	uri := "https://cloudplatform.coveo.com/rest/search/v2/?access_token=7b9b9300-3901-437b-bafd-51ae596f1b16&q=" + q

	resp, err := http.Get(uri)
	fmt.Println("Uri: " + uri)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer resp.Body.Close()

	var queryResponse QueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&queryResponse); err != nil {
		log.Println(err)
	}

	return queryResponse
}

func DoGetFields() CoveoFields {
	uri := "https://cloudplatform.coveo.com/rest/search/v2/fields?access_token=7b9b9300-3901-437b-bafd-51ae596f1b16"
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer resp.Body.Close()

	var fields CoveoFields
	if err := json.NewDecoder(resp.Body).Decode(&fields); err != nil {
		log.Println(err)
	}

	return fields
}
