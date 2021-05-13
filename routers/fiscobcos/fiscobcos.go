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
				Size:       opt.Size,
				CommitShas: opt.CommitShas,
				Time:       opt.Time,
			},
			logger)
	case "query push info":
		repo.SelectPushInfo(ctx, opt.PushID, logger)
	case "upload pull request info":
		repo.UploadPullRequestInfo(ctx,
			api.UploadPullReuqestOption{
				PullRequestID:     opt.PullRequestID,
				PullRequestNumber: opt.PullRequestNumber,
				RepoID:            opt.RepoID,
				Reponame:          opt.Reponame,
				Ownername:         opt.Ownername,
				Username:          opt.Username,
				Action:            opt.Action,
				Title:             opt.Title,
				Content:           opt.Content,
				SourceBranch:      opt.SourceBranch,
				TargetRepoID:      opt.TargetRepoID,
				TargetBranch:      opt.TargetBranch,
				Reviewers:         opt.Reviewers,
				CommitShas:        opt.CommitShas,
				MergeUser:         opt.MergeUser,
				CreatedAt:         opt.CreatedAt,
				UpdatedAt:         opt.UpdatedAt,
			},
			logger)
	case "query pull request latest info":
		repo.SelectPullRequestLatestInfo(ctx, opt.PullRequestID, logger)
	case "query pull request all info":
		repo.SelectPullRequestAllInfo(ctx, opt.PullRequestID, logger)
	case "upload pull request comment info":
		repo.UploadPullRequestCommentInfo(ctx,
			api.UploadPullReuqestCommentOption{
				PullRequestCommentID:     opt.PullRequestCommentID,
				PullRequestCommentNumber: opt.PullRequestCommentNumber,
				PullRequestNumber:        opt.PullRequestNumber,
				RepoID:                   opt.RepoID,
				ParentID:                 opt.ParentID,
				Reponame:                 opt.Reponame,
				Ownername:                opt.Ownername,
				Username:                 opt.Username,
				Action:                   opt.Action,
				CreatedAt:                opt.CreatedAt,
				UpdatedAt:                opt.UpdatedAt,
				Content:                  opt.Content,
				PullRequestID:            opt.PullRequestID,
			},
			logger)
	case "query pull request comment latest info":
		repo.SelectPullRequestCommentLatestInfo(ctx, opt.PullRequestCommentID, logger)
	case "query pull request comment all info":
		repo.SelectPullRequestCommentAllInfo(ctx, opt.PullRequestCommentID, logger)
	case "upload issue info":
		repo.UploadIssueInfo(ctx,
			api.UploadIssueOption{
				IssueID:     opt.IssueID,
				IssueNumber: opt.IssueNumber,
				RepoID:      opt.RepoID,
				Reponame:    opt.Reponame,
				Ownername:   opt.Ownername,
				Username:    opt.Username,
				Action:      opt.Action,
				Title:       opt.Title,
				CreatedAt:   opt.CreatedAt,
				UpdatedAt:   opt.UpdatedAt,
				Content:     opt.Content,
			},
			logger)
	case "query issue latest info":
		repo.SelectIssueLatestInfo(ctx, opt.IssueID, logger)
	case "query issue all info":
		repo.SelectIssueAllInfo(ctx, opt.IssueID, logger)
	case "upload issue comment info":
		repo.UploadIssueCommentInfo(ctx,
			api.UploadIssueCommentOption{
				IssueCommentID:     opt.IssueCommentID,
				IssueCommentNumber: opt.IssueCommentNumber,
				IssueNumber:        opt.IssueNumber,
				RepoID:             opt.RepoID,
				ParentID:           opt.ParentID,
				Reponame:           opt.Reponame,
				Ownername:          opt.Ownername,
				Username:           opt.Username,
				Action:             opt.Action,
				CreatedAt:          opt.CreatedAt,
				UpdatedAt:          opt.UpdatedAt,
				Content:            opt.Content,
				IssueID:            opt.IssueID,
			},
			logger)
	case "query issue comment latest info":
		repo.SelectIssueCommentLatestInfo(ctx, opt.IssueCommentID, logger)
	case "query issue comment all info":
		repo.SelectIssueCommentAllInfo(ctx, opt.IssueCommentID, logger)
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
