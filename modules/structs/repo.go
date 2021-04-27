package structs

// CreateRepoOption is options when to create a repository
type CreateRepoOption struct {
	Owner        string          `json:"owner"`         // 创始人
	TokenName    string          `json:"token_name"`    // 项目代号（默认等同于项目名）
	TotalSupply  uint64          `json:"total_supply"`  // 项目初始金额
	CurSupply    uint64          `json:"cur_supply"`    // 项目当前金额
	TokenBalance [][]interface{} `json:"token_balance"` // 用户初始Token分配
}

type Repo struct {
	Owner       string `json:"username"`     // 创始人
	TokenName   string `json:"token_name"`   // 项目代号（默认等同于项目名）
	TotalSupply uint64 `json:"total_supply"` // 项目初始金额
	CurSupply   uint64 `json:"cur_supply"`   // 项目当前金额
}

type ReponseRepo struct {
	Repo
	Response
}
