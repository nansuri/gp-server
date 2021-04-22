package util

import (
	"log"

	jira "github.com/andygrunwald/go-jira"
	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
)

func CreateJiraIssue(ticketDetail model.JiraRequest) (key string, errorMessage string) {

	var jiraKey string
	_ = jiraKey

	base := config.JiraUrl
	tp := jira.BasicAuthTransport{
		Username: config.JiraUsername,
		Password: config.JiraToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		// panic(err)
		errorMessage = "system error"
		log.Fatal(errorMessage)
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				EmailAddress: ticketDetail.Assignee,
			},
			Reporter: &jira.User{
				EmailAddress: ticketDetail.Reporter,
			},
			Description: ticketDetail.Description,
			Type: jira.IssueType{
				Name: ticketDetail.Type,
			},
			Project: jira.Project{
				Key: ticketDetail.Project,
			},
			Priority: &jira.Priority{
				Name: ticketDetail.Priority,
			},
			Summary: ticketDetail.Summary,
		},
	}
	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		// panic(err)
		errorMessage = "value on request body is invalid"
	}

	if issue == nil {
		jiraKey = ""
	} else {
		jiraKey = issue.Key
	}
	return jiraKey, errorMessage
}
