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
}

type Version struct {
	Current string
	Latest  string
}

type InstallResult struct {
	Name    string
	Version string
	Status  bool
}

var formulaType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Formula",
	Fields: graphql.Fields{
		"name":    &graphql.Field{Type: graphql.String},
		"version": &graphql.Field{Type: versionType},
		"latest":  &graphql.Field{Type: graphql.String},
	},
})

var versionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Version",
	Fields: graphql.Fields{
		"current": &graphql.Field{Type: graphql.String},
		"latest":  &graphql.Field{Type: graphql.String},
	},
})

var installType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Install",
	Fields: graphql.Fields{
		"name":    &graphql.Field{Type: graphql.String},
		"version": &graphql.Field{Type: graphql.String},
		"status":  &graphql.Field{Type: graphql.Boolean},
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

				return formulas, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"install": &graphql.Field{
			Type:        installType,
			Description: "Install formula",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"version": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var result = InstallResult{Name: "", Version: "", Status: false}
				name, isOk := params.Args["name"].(string)
				if !isOk {
					return result, nil
				}
				fmt.Printf("Param name: %v \n", name)
				// fmt.Printf("Param version: %v \n", params.Args["version"].(string))
				result.Name = name
				out, err := exec.Command("brew", "install", name).CombinedOutput()
				if err != nil {
					fmt.Println(err.Error())
					return "cannot get version", nil
				}
				stdout := string(out)
				lines := strings.Split(stdout, "\n")
				versionLine := strings.Split(lines[len(lines)-2], " ")
				fmt.Printf("versionLine: %v\n", versionLine)
				repoInfo := strings.Split(versionLine[2], "/")
				version := strings.TrimSuffix(repoInfo[len(repoInfo)-1], ":")
				fmt.Printf("version: %v\n", version)

				result.Version = version
				result.Status = true

				return result, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
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
