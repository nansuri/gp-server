package util

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	model "github.com/nansuri/gp-server/model"
)

func CreateIssue(ticketDetail model.JiraRequest) string {
	base := "https://danaindonesia.atlassian.net"
	tp := jira.BasicAuthTransport{
		Username: "onduty.bot@dana.id",
		Password: "tnTynAPy8TCQz32cfnC7DB47",
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: "myuser",
			},
			Reporter: &jira.User{
				Name: "youruser",
			},
			Description: ticketDetail.Description,
			Type: jira.IssueType{
				Name: ticketDetail.Type,
			},
			Project: jira.Project{
				Key: ticketDetail.Project,
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
