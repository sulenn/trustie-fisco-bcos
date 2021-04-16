package structs

type User struct {
	Username  string `json:"username" binding:"Required"`
	TokenName string `json:"token_name"`
}

type UserBalance struct {
	User
	Balance uint64 `json:"balance"`
}

type UserBalanceList []*UserBalance

// type CreateUserOption struct {
// 	Username string `json:"username" binding:"Required"`
// 	Orgname  string `json:"orgnization_name" binding:"Required"`
// }
