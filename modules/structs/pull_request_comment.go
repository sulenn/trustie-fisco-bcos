package structs

type UploadPullReuqestCommentOption struct {
	PullRequestCommentID     string `json:"pull_request_comment_id"`
	PullRequestCommentNumber uint64 `json:"pull_request_comment_number"`
	ParentID                 string `json:"parent_id"`
	PullRequestNumber        uint64 `json:"pull_request_number"`
	RepoID                   string `json:"repo_id"`
	Reponame                 string `json:"reponame"`
	Ownername                string `json:"ownername"`
	Username                 string `json:"username"`
	Action                   string `json:"action"`
	Content                  string `json:"content"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
}

type ResponsePullRequestComment struct {
	UploadPullReuqestCommentOption
	Response
}

type ResponsePullRequestCommentArr struct {
	UploadPullReuqestCommentOptionArr
	Response
}

type UploadPullReuqestCommentOptionArr []UploadPullReuqestCommentOption
