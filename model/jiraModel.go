package model

type JiraRequest struct {
	Project      string `json:"project"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	Assignee     string `json:"assignee"`
	Priority     string `json:"priority"`
	Reporter     string `json:"reporter"`
	Type         string `json:"type"`
	AssigneeRole string `json:"assignee_role"`
	IsUrgent     string `json:"is_urgent"`
	ExtendInfo   string `json:"extend_info"`
}

type JiraResult struct {
	Status bool   `json:"success"`
	Key    string `json:"jira_key"`
	Error  string `json:"error_detail"`
}
