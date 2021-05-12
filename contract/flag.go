package contract

var (
	RepoExisted                int64 = -1001 // 在合约中创建项目时，项目已存在
	BalanceOverFlow            int64 = -1002 // 在合约中创建项目时，预分配的用户Token累积值大于项目总Token
	RepoUnexisted              int64 = -1003 // 项目不存在
	RepoCurBalanceUpOverFlow   int64 = -1004 // 当前项目总Token 数超过最大 Token 量
	RepoCurBalanceDownOverFlow int64 = -1005 // 当前项目总 Token 数低于0
	UserBblanceNotEnough       int64 = -1006 // 用户余额不足

	ContractAddress string = "0xE46EE87f19049d5705989f49305b136d0732a0Cf"
)
