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
	case "query user balance of single repo":
		user.SelectUserBalance(ctx, opt.Username, opt.TokenName, logger)
	case "query user balance of all repos":
		user.SelectUserAllBalance(ctx, opt.Username, logger)
	default:
		ctx.JSON(http.StatusNotFound, "the request type not found!")
	}
}