package structs

type UploadIssueOption struct {
	IssueID     string `json:"issue_id"`
	IssueNumber uint64 `json:"issue_number"`
	RepoID      string `json:"repo_id"`
	Reponame    string `json:"reponame"`
	Ownername   string `json:"ownername"`
	Username    string `json:"username"`
	Action      string `json:"action"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ResponseIssue struct {
	UploadIssueOption
	Response
}

type ResponseIssueArr struct {
	UploadIssueOptionArr
	Response
}

type UploadIssueOptionArr []UploadIssueOption
