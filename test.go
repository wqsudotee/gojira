package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"strings"
)

type issew struct {
	key     string
	status  string
	summary string
}

func getissues(client jira.Client) []jira.Issue {
	var searchopts jira.SearchOptions
	searchopts.MaxResults = 1000
	x, y, z := client.Issue.Search("fixVersion = Production", &searchopts)
	if y != nil || z != nil {
		// oh, dis be bad code mon.
	}
	return x
}

func prnl() {
	fmt.Printf("\n")
}

func main() {
	tp := jira.BasicAuthTransport{
		Username: "",
		Password: "",
	}

	// get client
	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(""))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	var production_issues []issew
	thingy := getissues(*client)
	for i := 0; i < len(thingy); i++ {
		if thingy[i].Fields.FixVersions[0].Name == "Production" {
			var issue issew
			issue.key = thingy[i].Key
			issue.status = thingy[i].Fields.Status.Name
			issue.summary = thingy[i].Fields.Summary
			production_issues = append(production_issues, issue)
		}
	}

	for i := 0; i < len(production_issues); i++ {
		fmt.Printf(production_issues[i].key)
		prnl()
		fmt.Printf(production_issues[i].status)
		prnl()
		fmt.Printf(production_issues[i].summary)
		prnl()
		prnl()
	}
}
