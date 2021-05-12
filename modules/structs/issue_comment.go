package structs

type UploadIssueCommentOption struct {
	IssueCommentID     string `json:"issue_comment_id"`
	IssueCommentNumber uint64 `json:"issue_comment_number"`
	IssueNumber        uint64 `json:"issue_number"`
	RepoID             string `json:"repo_id"`
	ParentID           string `json:"parent_id"`
	Reponame           string `json:"reponame"`
	Ownername          string `json:"ownername"`
	Username           string `json:"username"`
	Action             string `json:"action"`
	Content            string `json:"content"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

type ResponseIssueComment struct {
	UploadIssueCommentOption
	Response
}

type ResponseIssueCommentArr struct {
	UploadIssueCommentOptionArr
	Response
}

type UploadIssueCommentOptionArr []UploadIssueCommentOption
