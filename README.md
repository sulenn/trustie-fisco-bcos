## 目标

使用区块链技术存储开源开发数据，保障数据不可篡改、支持数据溯源、支持记录用户开源激励积分等能力

## 技术方案

- 区块链技术：国内成熟联盟链  - [fiscso bcos](https://github.com/FISCO-BCOS/FISCO-BCOS)
- go-sdk：FISCO BCOS Go语言版本的SDK - [go-sdk](https://github.com/FISCO-BCOS/go-sdk)
- macaron：高性能、模块化的 go web 框架 - [go-macaron](https://go-macaron.com/)

## 项目介绍

### 目录结构

```shell
├── 0x83309d045a19c44dc3722d15a6abd472f95866ac.pem
├── bin  # 忘记干啥用的了，如果删掉后不影响项目正常启动，可以删掉
│   └── contract
│       └── opensource
│           ├── OpenSource.abi
│           ├── OpenSource.bin
│           ├── OpenSource.json
│           └── opensource-solc-output.json
├── ca.crt # fisco-bcos 的证书
├── cert.cnf # fisco-bcos 的证书
├── config.toml # 项目配置文件
├── contract # 智能合约、以及.sol转.go的中间文件
│   ├── deploy_call
│   │   ├── call
│   │   │   └── insert.go
│   │   └── deploy
│   │       └── deploy.go
│   ├── flag.go
│   └── opensource
│       ├── Condition.abi
│       ├── Condition.bin
│       ├── Entries.abi
│       ├── Entries.bin
│       ├── Entry.abi
│       ├── Entry.bin
│       ├── KVTable.abi
│       ├── KVTable.bin
│       ├── KVTableFactory.abi
│       ├── KVTableFactory.bin
│       ├── OpenSource.abi
│       ├── OpenSource.bin
│       ├── OpenSource.go
│       ├── opensource.sol
│       ├── Table.abi
│       ├── Table.bin
│       ├── TableFactory.abi
│       ├── TableFactory.bin
│       └── Table.sol
├── go.mod
├── go.sum
├── main.go # web 项目主入口
├── modules # 项目基础包、关键数据结构
│   ├── format
│   │   ├── chaincode_args.go
│   │   └── chaincode_args_test.go
│   ├── os
│   │   └── common.go
│   └── structs
│       ├── commit.go
│       ├── fiscobcos.go
│       ├── issue_comment.go
│       ├── issue.go
│       ├── pull_request_comment.go
│       ├── pull_request.go
│       ├── push.go
│       ├── repo.go
│       ├── response.go
│       └── user.go
├── routers # web 路由、定义核心接口处理方法
│   ├── fiscobcos
│   │   └── fiscobcos.go
│   ├── repo
│   │   ├── query.go
│   │   ├── repo.go
│   │   └── storage.go
│   ├── routes
│   │   └── routes.go
│   └── user
│       └── user.go
├── sdk.crt # fisco-bcos 的证书
├── sdk.key # fisco-bcos 的证书
├── tabletest # 智能合约测试文件
│   ├── contract
│   │   ├── tabletest_insert.go
│   │   ├── tabletest_main.go
│   │   ├── tabletest_remove.go
│   │   ├── tabletest_select.go
│   │   └── tabletest_update.go
│   ├── Table.abi
│   ├── Table.bin
│   ├── Table.sol
│   ├── TableTest.abi
│   ├── TableTest.bin
│   ├── TableTest.go
│   └── TableTest.sol
└── tools # 用于将 .sol 转译为 .go 文件的工具
    ├── abigen
    └── solc-0.4.25
```

### 环境搭建

- **fisco-bcos 联盟链搭建自行参考官方文档**：https://fisco-bcos-documentation.readthedocs.io/zh-cn/stable/docs/installation.html
- **go环境搭建**：自行参考网上教程，版本可用 `1.22.7`。需要开启 `GO111MODULE='on'`，可使用命令 `执行 go` 安装项目依赖包
- 将 `fisco-bcos` 关键证书拷贝到 go 项目主目录中 `ca.crt`、`sdk.crt`、`sdk.key` 等

#### 部署指定合约

- 执行 `go run contract/deploy_call/deploy/deploy.go` 部署智能合约，获取合约地址。如：`0x0a68F060B46e0d8f969383D260c34105EA13a9dd`

- 将  `./contract/flag.go` 中的 `ContractAddress` 变量值替换成新的合约地址

#### 运行 web 项目

执行 `go run main.go` 即可

## 接口介绍

参考 `./routers/fiscobcos/fiscobcos.go` 即可