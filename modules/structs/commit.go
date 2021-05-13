package structs

// UploadCommitOption is to upload commit info to blockchain
type UploadCommitOption struct {
	CommitHash     string `json:"commit_hash"`
	RepoID         string `json:"repo_id"`
	Author         string `json:"author"`
	AuthorEmail    string `json:"author_email"`
	Committer      string `json:"committer"`
	CommitterEmail string `json:"committer_email"`
	AuthorTime     string `json:"author_time"`
	CommitterTime  string `json:"committer_time"`
	Content        string `json:"content"`
	CommitDiff     string `json:"commit_diff"`
}

type ResponseCommit struct {
	UploadCommitOption
	Response
}
