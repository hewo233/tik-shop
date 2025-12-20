package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hewo/tik-shop/route/biz/model/hewo/tikshop/route/base"
	"github.com/hewo/tik-shop/shared/consts"
	"net/http"
	"strconv"
)

// NormalFuncChecker : self and admin can pass
func NormalFuncChecker[T any](ctx context.Context, c *app.RequestContext, id int64, buildResp func(response *base.BaseResponse) T) bool {
	idStr := strconv.FormatInt(id, 10)
	idCtx, exi := c.Get(consts.AccountID)
	if !exi {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40080,
			Message: "cannot get account id from context",
		}))
		return false
	}

	idCtxStr, ok := idCtx.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40081,
			Message: "account id in context has invalid type",
		}))
		return false
	}

	roleCtx, exi := c.Get(consts.Audience)
	if !exi {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40080,
			Message: "cannot get role from context",
		}))
		return false
	}

	roleCtxStr, ok := roleCtx.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, buildResp(&base.BaseResponse{
			Code:    40081,
			Message: "role in context has invalid type",
		}))
		return false
	}

	if idCtxStr != idStr && roleCtxStr != consts.Admin {
		c.JSON(http.StatusUnauthorized, buildResp(&base.BaseResponse{
			Code:    40081,
			Message: "permission denied",
		}))
		return false
	}
	return true
}
