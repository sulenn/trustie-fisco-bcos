package structs

type FiscoBcos struct {
	RequestType string `json:"request-type" binding:"Required"`

	// repo 相关
	Owner        string          `json:"owner"`         // 创始人
	TokenName    string          `json:"token_name"`    // 项目代号（默认等同于项目名）
	TotalSupply  uint64          `json:"total_supply"`  // 项目初始金额
	CurSupply    uint64          `json:"cur_supply"`    // 项目当前金额
	TokenBalance [][]interface{} `json:"token_balance"` // 用户初始Token分配

	// user 相关
	Username string `json:"username"` // 用户名
}
