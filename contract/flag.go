package contract

var (
	RepoExisted     int64 = -1001 // 在合约中创建项目时，项目已存在
	BalanceOverFlow int64 = -1002 // 在合约中创建项目时，预分配的用户Token累积值大于项目总Token

	ContractAddress string = "0xE46EE87f19049d5705989f49305b136d0732a0Cf"
)
