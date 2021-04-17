package structs

type User struct {
	Username  string `json:"username"`
	TokenName string `json:"token_name"`
}

type UserBalance struct {
	User
	Balance uint64 `json:"balance"`
}

type UserBalanceList []*UserBalance
