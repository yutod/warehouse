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

type Doctor struct {
	Errors, Warnings []Message
}

type Message struct {
	Subject string
	Detail  []string
}

type FormulaInfo struct {
	Name    string
	Version Version
}

type Version struct {
	Current string
	Latest  string
}

type UpdateResult struct {
	Name   string
	Status bool
}

type InstallResult struct {
	Name    string
	Status  bool
	Version string
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

var doctorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Doctor",
	Fields: graphql.Fields{
		"errors":   &graphql.Field{Type: graphql.NewList(graphql.String)},
		"warnings": &graphql.Field{Type: graphql.NewList(graphql.String)},
	},
})

var installType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Install",
	Fields: graphql.Fields{
		"name":    &graphql.Field{Type: graphql.String},
		"status":  &graphql.Field{Type: graphql.Boolean},
		"version": &graphql.Field{Type: graphql.String},
	},
})

var updateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Update",
	Fields: graphql.Fields{
		"name":   &graphql.Field{Type: graphql.String},
		"status": &graphql.Field{Type: graphql.Boolean},
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
			Type:        doctorType,
			Description: "Homebrew doctor",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var result = Doctor{Errors: []Message{}, Warnings: []Message{}}
				out, _ := exec.Command("brew", "doctor").CombinedOutput()

				var log []string
				var logType *[]Message
				for _, line := range strings.Split(string(out), "\n") {
					if strings.Contains(line, "Error:") {
						logType = &result.Errors
					} else if strings.Contains(line, "Warning:") {
						logType = &result.Warnings
					}

					if strings.Contains(line, "Error:") || strings.Contains(line, "Warning:") {
						if log != nil {
							message := Message{Subject: log[0], Detail: log[1:]}
							*logType = append(*logType, message)
							log = nil
						}
						log = append(log, line)
					} else if len(log) > 0 && line != "" {
						log = append(log, line)
					}
				}
				message := Message{Subject: log[0], Detail: log[1:]}
				*logType = append(*logType, message)
				log = nil
				// fmt.Printf("%v", result)

				return result, nil
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
					fullName := strings.Split(strings.Trim(info[0], ":"), "/")
					name := fullName[len(fullName)-1]
					latestVersions[name] = strings.Trim(info[2], ",")
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
				var result = InstallResult{Name: "", Status: false, Version: ""}
				name, isOk := params.Args["name"].(string)
				if !isOk {
					return result, nil
				}
				result.Name = name
				out, err := exec.Command("brew", "install", name).CombinedOutput()
				if err != nil {
					fmt.Println(err.Error())
					return result, nil
				}
				stdout := string(out)
				lines := strings.Split(stdout, "\n")
				versionLine := strings.Split(lines[len(lines)-2], " ")
				repoInfo := strings.Split(versionLine[2], "/")
				version := strings.TrimSuffix(repoInfo[len(repoInfo)-1], ":")

				result.Version = version
				result.Status = true

				return result, nil
			},
		},
		"upgrade": &graphql.Field{
			Type:        installType,
			Description: "Upgrade Formula",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"version": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var result = InstallResult{Name: "", Status: false, Version: ""}
				name, isOk := params.Args["name"].(string)
				if !isOk {
					return result, nil
				}
				version, isOk := params.Args["version"].(string)
				if !isOk {
					return result, nil
				}
				result.Name = name
				out, err := exec.Command("brew", "upgrade", name).CombinedOutput()
				if err != nil {
					fmt.Println(err.Error())
					return result, nil
				}
				for _, line := range strings.Split(string(out), "\n") {
					info := strings.Split(string(line), " ")
					if info[0] == name && info[1] == version && info[2] == "->" {
						result.Version = info[3]
						result.Status = true
					}
				}

				return result, nil
			},
		},
		"delete": &graphql.Field{
			Type:        updateType,
			Description: "Delete Formula",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var result = UpdateResult{"", false}
				name, isOk := params.Args["name"].(string)
				if !isOk {
					return result, nil
				}
				result.Name = name
				out, err := exec.Command("brew", "uninstall", name).CombinedOutput()
				if err != nil {
					fmt.Println(err.Error())
					return result, nil
				}
				stdout := string(out)
				info := strings.Split(stdout, " ")
				if info[0] == "Uninstalling" {
					result.Status = true
				}

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
