package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"version": &graphql.Field{
			Type:        graphql.String,
			Description: "Homebrew version",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				out, err := exec.Command("brew", "-v").CombinedOutput()
				// fmt.Printf("version result: %s\n", string(out))
				if err != nil {
					fmt.Println(err.Error())
					return "cannot get version", nil
				}
				stdout := string(out)
				version := strings.Split(stdout, "\n")

				return version[0], nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Error! Unexpected errors: %v", result.Errors)
	}

	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":3001", nil)
}
