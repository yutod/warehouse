package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/graphql-go/graphql"
)

type FormulaInfo struct {
	Name    string
	Version Version
	// Version string
	// Latest  string
}

type Version struct {
	Current string
	Latest  string
}

var formulaType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Formula",
	Fields: graphql.Fields{
		"name":    &graphql.Field{Type: graphql.String},
		"version": &graphql.Field{Type: versionType},
		// "latest":  &graphql.Field{Type: graphql.String},
	},
})

var versionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Version",
	Fields: graphql.Fields{
		"current": &graphql.Field{Type: graphql.String},
		"latest":  &graphql.Field{Type: graphql.String},
	},
})

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
		"doctor": &graphql.Field{
			Type:        graphql.String,
			Description: "Homebrew doctor",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				out, err := exec.Command("brew", "doctor").CombinedOutput()
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
		"installed": &graphql.Field{
			Type:        graphql.NewList(formulaType),
			Description: "Homebrew installed formula list",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				command1 := exec.Command("brew", "list")
				command2 := exec.Command("brew", "list", "--versions")
				command2.Stdin, _ = command1.StdoutPipe()
				var out bytes.Buffer
				command2.Stdout = &out
				_ = command2.Start()
				_ = command1.Run()
				_ = command2.Wait()

				command11 := exec.Command("brew", "list")
				command12 := exec.Command("xargs", "brew", "info")
				command13 := exec.Command("grep", "stable")
				command12.Stdin, _ = command11.StdoutPipe()
				command13.Stdin, _ = command12.StdoutPipe()
				var out1 bytes.Buffer
				command13.Stdout = &out1
				_ = command13.Start()
				_ = command12.Start()
				_ = command11.Run()
				_ = command12.Wait()
				_ = command13.Wait()

				latestVersions := make(map[string]string)
				for _, line := range strings.Split(string(out1.Bytes()), "\n") {
					info := strings.Split(string(line), " ")
					if info[0] == "" {
						continue
					}
					latestVersions[strings.Trim(info[0], ":")] = info[2]
				}

				var formulas []FormulaInfo
				for _, line := range strings.Split(string(out.Bytes()), "\n") {
					info := strings.Split(string(line), " ")
					if info[0] == "" {
						continue
					}
					version := Version{Current: info[1], Latest: latestVersions[info[0]]}
					formula := FormulaInfo{Name: info[0], Version: version}
					formulas = append(formulas, formula)
				}

				fmt.Printf("formulas: %v", formulas)
				return formulas, nil
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

func Start() {
	// cmd.Execute()
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8082", nil)
}
