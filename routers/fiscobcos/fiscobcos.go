package fiscobcos

import (
	"log"
	"net/http"

	"gitea.com/macaron/macaron"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"
	"github.com/sulenn/trustie-fisco-bcos/routers/repo"
	"github.com/sulenn/trustie-fisco-bcos/routers/user"
)

// 路由统一处理
func HandleAllRoutes(ctx *macaron.Context, opt api.FiscoBcos, logger *log.Logger) {
	requestType := opt.RequestType
	switch requestType {
	case "create repo":
		repo.Create(ctx,
			api.CreateRepoOption{
				Owner:        opt.Owner,
				TokenName:    opt.TokenName,
				TotalSupply:  opt.TotalSupply,
				CurSupply:    opt.CurSupply,
				TokenBalance: opt.TokenBalance,
			},
			logger)
	case "query repo basic info":
		repo.QueryBasic(ctx, opt.TokenName, logger)
	case "upload commit info":
		repo.UploadCommitInfo(ctx,
			api.UploadCommitOption{
				CommitHash: opt.CommitHash,
				RepoID:     opt.RepoID,
				Author:     opt.Author,
				Email:      opt.Email,
				Time:       opt.Time,
				Content:    opt.Content,
				CommitDiff: opt.CommitDiff,
			},
			logger)
	case "query commit info":
		repo.SelectCommitInfo(ctx, opt.CommitHash, logger)
	case "upload push info":
		repo.UploadPushInfo(ctx,
			api.UploadPushOption{
				PushID:     opt.PushID,
				PushNumber: opt.PushNumber,
				RepoID:     opt.RepoID,
				Reponame:   opt.Reponame,
				Ownername:  opt.Ownername,
				Username:   opt.Username,
				Branch:     opt.Branch,
				CommitShas: opt.CommitShas,
				Time:       opt.Time,
			},
			logger)
	case "query push info":
		repo.SelectPushInfo(ctx, opt.PushID, logger)
	case "query user balance of single repo":
		user.SelectUserBalance(ctx, opt.Username, opt.TokenName, logger)
	case "query user balance of all repos":
		user.SelectUserAllBalance(ctx, opt.Username, logger)
	case "add user balance":
		user.AddUserAmount(ctx, opt.Username, opt.TokenName, opt.Amount, logger)
	case "minus user balance":
		user.MinusUserAmount(ctx, opt.Username, opt.TokenName, opt.Amount, logger)
	case "lock user balance":
		user.LockUserAmount(ctx, opt.Username, opt.TokenName, opt.Amount, logger)
	case "unlock user balance":
		user.UnlockUserAmount(ctx, opt.Username, opt.TokenName, opt.Amount, logger)
	case "transfer amount":
		user.TransferAmount(ctx, opt.Payer, opt.Payee, opt.TokenName, opt.Amount, logger)
	default:
		ctx.JSON(http.StatusNotFound, "the request type not found!")
	}
}
