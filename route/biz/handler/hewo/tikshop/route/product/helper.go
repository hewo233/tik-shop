package product

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hewo/tik-shop/route/biz/model/hewo/tikshop/route/base"
	"github.com/hewo/tik-shop/shared/consts"
	"net/http"
	"strconv"
)

// NormalMerchantChecker : get merchant id and check role
func NormalMerchantChecker[T any](ctx context.Context, c *app.RequestContext, buildResp func(response *base.BaseResponse) T) (id int64, ok bool) {
	idCtx, exi := c.Get(consts.AccountID)
	if !exi {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40080,
			Message: "cannot get account id from context",
		}))
		return -1, false
	}

	idCtxStr, ok := idCtx.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40081,
			Message: "account id in context has invalid type",
		}))
		return -1, false
	}

	roleCtx, exi := c.Get(consts.Audience)
	if !exi {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40080,
			Message: "cannot get role from context",
		}))
		return -1, false
	}

	roleCtxStr, ok := roleCtx.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40081,
			Message: "role in context has invalid type",
		}))
		return -1, false
	}

	if roleCtxStr != consts.RoleMerchant && roleCtxStr != consts.RoleAdmin {
		c.JSON(http.StatusUnauthorized, buildResp(&base.BaseResponse{
			Code:    40081,
			Message: "permission denied: not a merchant or admin",
		}))
		return -1, false
	}

	id, err := strconv.ParseInt(idCtxStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40082,
			Message: "account id in context has invalid format",
		}))
		return -1, false
	}

	return id, true
}
