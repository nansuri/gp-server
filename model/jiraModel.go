package model

type JiraRequest struct {
	Project      string `json:"project"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	Assignee     string `json:"assignee"`
	Reporter     string `json:"reporter"`
	Type         string `json:"type"`
	AssigneeRole string `json:"AssigneeRole"`
	ExtendInfo   string `json:"ExtendInfo"`
}

type JiraResult struct {
	Status bool   `json:"success"`
	Key    string `json:"jira_key"`
}

type JiraTicketDetails struct {
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	Assignee     string `json:"assignee"`
	Reporter     string `json:"reporter"`
	Type         string `json:"type"`
	AssigneeRole string `json:"AssigneeRole"`
	ExtendInfo   string `json:"ExtendInfo"`
}
