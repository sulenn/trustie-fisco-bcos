package structs

type UploadPushOption struct {
	PushID     string   `json:"push_Id"`
	PushNumber uint64   `json:"push_number"`
	Reponame   string   `json:"reponame"`
	Ownername  string   `json:"ownername"`
	Branch     string   `json:"branch"`
	CommitShas []string `json:"commit_shas"`
	RepoID     string   `json:"repo_id"`
	Username   string   `json:"username"` // 用户名
	Time       string   `json:"time"`
}

type ResponsePush struct {
	UploadPushOption
	Response
}
