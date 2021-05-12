package structs

type FiscoBcos struct {
	RequestType string `json:"request-type" binding:"Required"`

	// repo 相关
	Owner        string          `json:"owner"`         // 创始人
	TokenName    string          `json:"token_name"`    // 项目代号（默认等同于项目名）
	TotalSupply  uint64          `json:"total_supply"`  // 项目初始金额
	CurSupply    uint64          `json:"cur_supply"`    // 项目当前金额
	TokenBalance [][]interface{} `json:"token_balance"` // 用户初始Token分配

	// repo commit 相关
	CommitHash string `json:"commit_hash"`
	RepoID     string `json:"repo_id"`
	Author     string `json:"author"`
	Email      string `json:"email"`
	Time       string `json:"time"`
	Content    string `json:"message"`
	CommitDiff string `json:"commit_diff"`

	// repo push 相关
	PushID     string   `json:"push_Id"`
	PushNumber uint64   `json:"push_number"`
	Reponame   string   `json:"reponame"`
	Ownername  string   `json:"ownername"`
	Branch     string   `json:"branch"`
	CommitShas []string `json:"commit_shas"`
	// repo_id, username, time

	// user 相关
	Username string `json:"username"` // 用户名
	Amount   uint64 `json:"amount"`   // 转账金额
	Payer    string `json:"payer"`    // 转账人
	Payee    string `json:"payee"`    // 收款人
}
