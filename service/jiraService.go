package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	jira "github.com/andygrunwald/go-jira"
	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
)

func CreateJiraIssue(ticketDetail model.JiraRequest, assignee string) (key string, errorMessage string) {

	var jiraKey string
	_ = jiraKey

	base := config.JiraUrl
	tp := jira.BasicAuthTransport{
		Username: config.JiraUsername,
		Password: Decrypt(config.EncryptedJiraToken),
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		// panic(err)
		errorMessage = "system error"
		println(errorMessage)
	}

	// var labels []string
	labels := make([]string, 5)
	labels[0] = ticketDetail.Label

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				AccountID: GetAccountIdByEmail(ticketDetail.Assignee),
			},
			Reporter: &jira.User{
				AccountID: GetAccountIdByEmail(ticketDetail.Reporter),
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
			Labels:  labels,
		},
	}
	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		// panic(err)
		errorMessage = "value on request body is invalid"
		println(errorMessage)
	}

	if issue == nil {
		jiraKey = ""
	} else {
		jiraKey = issue.Key
	}
	return jiraKey, errorMessage
}

func GetAccountIdByEmail(email string) string {

	var userRes model.JiraUserResponse

	url := config.JiraUrl + "/rest/api/2/user/search?query=" + email
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", Decrypt(config.EncryptedAuth))
	// req.Header.Add("Cookie", "atlassian.xsrf.token=82a67398-2cf2-4f41-9d44-44664d6f572e_cb08ba08f15d6a87097e21db7430e1a27c377ff9_lin")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	stringResponse := string(body)
	trimmedString1 := strings.Trim(stringResponse, "[")
	trimmedString2 := strings.Trim(trimmedString1, "]")

	json.Unmarshal([]byte(trimmedString2), &userRes)
	// println(string(body))

	return userRes.AccountId
}
