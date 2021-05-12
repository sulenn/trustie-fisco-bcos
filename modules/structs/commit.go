package structs

// UploadCommitOption is to upload commit info to blockchain
type UploadCommitOption struct {
	CommitHash string `json:"commit_hash"`
	RepoID     string `json:"repo_id"`
	Author     string `json:"author"`
	Email      string `json:"email"`
	Time       string `json:"time"`
	Content    string `json:"content"`
	CommitDiff string `json:"commit_diff"`
}

type ResponseCommit struct {
	UploadCommitOption
	Response
}
