package utils

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
)

func transToThree(x int64) int {
	for x >= 1000 {
		x /= 10
	}
	return int(x)
}

func HandleRPCError(c *app.RequestContext, err error) {
	var errResp *base.ErrorResponse
	if errors.As(err, &errResp) {
		c.JSON(transToThree(errResp.Code), base.ErrorResponse{
			Code:    errResp.Code,
			Message: errResp.Message,
		})
		return
	}

	c.JSON(consts.StatusInternalServerError, base.ErrorResponse{
		Code:    consts.StatusInternalServerError,
		Message: err.Error(),
	})
}
