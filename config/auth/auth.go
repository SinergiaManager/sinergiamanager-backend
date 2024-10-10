package auth

import (
	"time"

	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	secret           = []byte("signature_hmac_secret_shared_key")
	Signer           = jwt.NewSigner(jwt.HS256, secret, 15*time.Minute)
	VerifyMiddleware = jwt.NewVerifier(jwt.HS256, secret).Verify(func() interface{} {
		return new(UserClaims)
	})
	Verifier = jwt.NewVerifier(jwt.HS256, secret).WithDefaultBlocklist()
)

type UserClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GenerateToken(signer *jwt.Signer, user *Models.UserDb) iris.Handler {
	return func(ctx iris.Context) {
		claims := UserClaims{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		token, err := signer.Sign(claims)
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		ctx.Write(token)
	}
}

func Protected(ctx iris.Context) {
	claims := jwt.Get(ctx).(*UserClaims)

	standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
	expiresAtString := standardClaims.ExpiresAt().
		Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
	timeLeft := standardClaims.Timeleft()

	ctx.Writef("foo=%s\nexpires at: %s\ntime left: %s\n", claims.Username, expiresAtString, timeLeft)
}

func Logout(ctx iris.Context) {
	err := ctx.Logout()
	if err != nil {
		ctx.WriteString(err.Error())
	} else {
		ctx.Writef("token invalidated, a new token is required to access the protected API")
	}
}
