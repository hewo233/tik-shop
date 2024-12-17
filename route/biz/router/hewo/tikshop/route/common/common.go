package common

import (
	agp "aidanwoods.dev/go-paseto"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/paseto"
	"github.com/hewo/tik-shop/route/config"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/hewo/tik-shop/shared/errno"
	"log"
	"net/http"
)

func ComMiddleWare() []app.HandlerFunc {
	return []app.HandlerFunc{
		Cors(),
		Recovery(),
		gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".jpg", ".mp4", ".png"})),
	}
}

func PasetoAuth(audience string) app.HandlerFunc {
	pasetoInfo := config.GlobalServerConfig.PasetoInfo

	pasetoParse, err := paseto.NewV4PublicParseFunc(
		pasetoInfo.PubKey,
		[]byte(pasetoInfo.Implicit),
		paseto.WithAudience(audience),
		paseto.WithNotBefore(),
	)
	if err != nil {
		log.Fatal(err)
	}

	successHandler := func(ctx context.Context, c *app.RequestContext, token *agp.Token) {
		accountID, err := token.GetString("id")
		if err != nil {
			errHere := errno.StatusBadRequest.WithMessage("bad request: missing accountID in token")
			c.JSON(http.StatusUnauthorized, errHere)
			c.Abort()
			return
		}
		c.Set(consts.AccountID, accountID)
	}

	failHandler := func(ctx context.Context, c *app.RequestContext) {
		errHere := errno.StatusBadRequest.WithMessage("bad request: invalid token")
		c.JSON(http.StatusUnauthorized, errHere)
		c.Abort()
	}

	return paseto.New(
		paseto.WithTokenPrefix("Bearer "),
		paseto.WithParseFunc(pasetoParse),
		paseto.WithSuccessHandler(successHandler),
		paseto.WithErrorFunc(failHandler),
	)
}
