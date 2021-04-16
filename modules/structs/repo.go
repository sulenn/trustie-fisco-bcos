package structs

// CreateRepoOption is options when to create a repository
type CreateRepoOption struct {
	Owner        string          `json:"username" binding:"Required"`     // 创始人
	TotalSupply  uint64          `json:"total_supply" binding:"Required"` // 项目初始金额
	TokenName    string          `json:"token_name" binding:"Required"`   // 项目名
	TokenBalance [][]interface{} `json:"token_balance"`                   // 用户初始Token分配
}

type Repo struct {
	Owner       string `json:"username"`     // 创始人
	TokenName   string `json:"token_name"`   // 项目代号（默认等同于项目名）
	TotalSupply uint64 `json:"total_supply"` // 项目初始金额
	CurSupply   uint64 `json:"cur_supply"`   // 项目当前金额
}

// TransferTokenOption is options when to transfer amount from one person to another
type TransferTokenOption struct {
	TokenSymbol string `json:"reponame" binding:"Required"` // 项目名
	From        string `json:"payer" binding:"Required"`    // 付款人
	To          string `json:"payee" binding:"Required"`    // 收款人
	Amount      uint64 `json:"amount" binding:"Required"`   // 金额
}

// AddTokenOption is options when user add amount
type AddTokenOption struct {
	TokenSymbol string `json:"reponame" binding:"Required"` // 项目名
	Username    string `json:"username" binding:"Required"` // 用户名
	Amount      uint64 `json:"amount" binding:"Required"`   // 金额
}

// MinusTokenOption is options when user minus amount
type MinusTokenOption struct {
	TokenSymbol string `json:"reponame" binding:"Required"` // 项目名
	Username    string `json:"username" binding:"Required"` // 用户名
	Amount      uint64 `json:"amount" binding:"Required"`   // 金额
}

// // UserBalance represents a user's balance
// type UserBalance struct {
// 	Username    string `json:"username" binding:"Required"` // 用户名
// 	TokenSymbol string `json:"reponame" binding:"Required"` // 项目名
// 	Balance     uint64 `json:"balance"`                     // 用户余额
// }

// // UserBalanceList represents all user's balance list in repository
// type UserBalanceList []*UserBalance
