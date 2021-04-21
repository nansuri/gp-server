package util

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
)

func CreateIssue(ticketDetail model.JiraRequest) string {
	base := config.JiraUrl
	tp := jira.BasicAuthTransport{
		Username: config.JiraUsername,
		Password: config.JiraToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				AccountID: ticketDetail.Assignee,
			},
			Reporter: &jira.User{
				AccountID: ticketDetail.Reporter,
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
		panic(err)
	}

	fmt.Printf("\nThe key is %s", issue.Key)
	return issue.Key
}
