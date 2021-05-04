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

type JiraUserResponse struct {
	Self         string `json:"self"`
	AccountId    string `json:"accountId"`
	AccountType  string `json:"accountType"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls   AvatarUrls
	DisplayName  string `json:"displayName"`
	Active       string `json:"active"`
	TimeZone     string `json:"timeZone"`
	Locale       string `json:"locale"`
}

type AvatarUrls struct {
	A string `json:"48x48"`
	B string `json:"24x24"`
	C string `json:"16x16"`
	D string `json:"32x32"`
}
