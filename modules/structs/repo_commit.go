package structs

// UploadCommitOption is to upload commit info to blockchain
type UploadCommitOption struct {
	Sha               string `json:"sha" binding:"Required"`                // sha 值
	Branch            string `json:"branch" binding:"Required"`             // 分支
	Ownername         string `json:"ownername" binding:"Required"`          // 项目创始人
	Reponame          string `json:"reponame" binding:"Required"`           // 项目名
	Username          string `json:"username" binding:"Required"`           // 用户名
	ModificationLines int    `json:"modification_lines" binding:"Required"` //修改行数
}
