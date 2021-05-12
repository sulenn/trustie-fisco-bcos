package structs

type UploadPullReuqestOption struct {
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestNumber uint64   `json:"pull_request_number"`
	Action            string   `json:"action"`
	Title             string   `json:"title"`
	SourceBranch      string   `json:"source_branch"`
	TargetRepoID      string   `json:"target_repo_id"`
	TargetBranch      string   `json:"target_branch"`
	Reviewers         []string `json:"reviewers"`
	MergeUser         string   `json:"merge_user"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_ad"`
	RepoID            string   `json:"repo_id"`
	Reponame          string   `json:"reponame"`
	Ownername         string   `json:"ownername"`
	Username          string   `json:"username"`
	Content           string   `json:"content"`
	CommitShas        []string `json:"commit_shas"`
}

type ResponsePullRequest struct {
	UploadPullReuqestOption
	Response
}

type ResponsePullRequestArr struct {
	UploadPullReuqestOptionArr
	Response
}

type UploadPullReuqestOptionArr []UploadPullReuqestOption
