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
	CommitHash     string `json:"commit_hash"`
	RepoID         string `json:"repo_id"`
	Author         string `json:"author"`
	AuthorEmail    string `json:"author_email"`
	Committer      string `json:"committer"`
	CommitterEmail string `json:"committer_email"`
	AuthorTime     string `json:"author_time"`
	CommitterTime  string `json:"committer_time"`
	Content        string `json:"content"`
	CommitDiff     string `json:"commit_diff"`

	// repo push 相关
	PushID     string   `json:"push_Id"`
	PushNumber uint64   `json:"push_number"`
	Reponame   string   `json:"reponame"`
	Ownername  string   `json:"ownername"`
	Branch     string   `json:"branch"`
	Size       uint64   `json:"size"`
	CommitShas []string `json:"commit_shas"`
	Time       string   `json:"time"`
	// repo_id, username, time

	// repo pull request 相关
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestNumber uint64   `json:"pull_request_number"`
	Action            string   `json:"action"`
	Title             string   `json:"title"`
	SourceBranch      string   `json:"source_branch"`
	TargetRepoID      string   `json:"target_repo_id"`
	TargetBranch      string   `json:"target_branch"`
	Reviewers         []string `json:"reviewers"`
	MergeUser         string   `json:"merge_user"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
	// repo_id, reponame, ownername, username, content, commit_shas

	// repo pull request comment 相关
	PullRequestCommentID     string `json:"pull_request_comment_id"`
	PullRequestCommentNumber uint64 `json:"pull_request_comment_number"`
	ParentID                 string `json:"parent_id"`
	// pull_request_number, repo_id, reponame, ownername, username, action, content, created_at, updated_at

	// repo issue 相关
	IssueID     string `json:"issue_id"`
	IssueNumber uint64 `json:"issue_number"`
	// repo_id, reponame, ownername, action, title, content, created_at, updated_at

	// repo issue comment 相关
	IssueCommentID     string `json:"issue_comment_id"`
	IssueCommentNumber uint64 `json:"issue_comment_number"`
	// issue_number, repo_id, parent_id, reponame, ownername, username, action, content, created_at, updated_at

	// user 相关
	Username string `json:"username"` // 用户名
	Amount   uint64 `json:"amount"`   // 转账金额
	Payer    string `json:"payer"`    // 转账人
	Payee    string `json:"payee"`    // 收款人

	// page 相关
	Page    int64 `json:"page"`     // 分页的第几页
	Pagenum int64 `json:"page_num"` // 分页每一页的需求数量
}
