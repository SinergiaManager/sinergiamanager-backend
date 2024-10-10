package auth

import (
	"fmt"
	"strings"
	"time"

	Enum "github.com/SinergiaManager/sinergiamanager-backend/config/utils"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	secret   = []byte("signature_hmac_secret_shared_key")
	Signer   = jwt.NewSigner(jwt.HS256, secret, 15*time.Minute)
	Verifier = jwt.NewVerifier(jwt.HS256, secret).WithDefaultBlocklist()
)

type UserClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func GenerateToken(signer *jwt.Signer, user *Models.UserDb) ([]byte, error) {
	claims := UserClaims{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(Enum.EnumUserRole.ADMIN),
	}

	token, err := signer.Sign(claims)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func Logout(ctx iris.Context) {
	err := ctx.Logout()
	if err != nil {
		ctx.WriteString(err.Error())
	} else {
		ctx.Writef("token invalidated, a new token is required to access the protected API")
	}
}

func JWTMiddleware(roles []string) iris.Handler {
	return func(ctx iris.Context) {
		tokenAuth := ctx.GetHeader("Authorization")
		if tokenAuth == "" {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "No token provided"})
			return
		}

		tokenParts := strings.Split(tokenAuth, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "Invalid token format"})
			return
		}

		token := tokenParts[1]

		verifiedToken, err := Verifier.VerifyToken([]byte(token))
		if err != nil {
			if strings.Contains(err.Error(), "token expired") {
				ctx.StatusCode(iris.StatusUnauthorized)
				ctx.JSON(iris.Map{"error": "Token expired. Please login again or refresh your token."})
				return
			}
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "Invalid token: " + err.Error()})
			return
		}

		var claims UserClaims
		if err := verifiedToken.Claims(&claims); err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "Failed to extract claims: " + err.Error()})
			return
		}

		fmt.Printf("Authenticated user claims: %+v\n", claims)

		ctx.Values().Set("user", claims)

		authorized := false
		if len(roles) > 0 {
			currentRole := claims.Role
			for _, role := range roles {
				if role == currentRole {
					authorized = true
					break
				}
			}
		} else {
			authorized = true
		}

		if !authorized {
			ctx.StatusCode(iris.StatusForbidden)
			ctx.JSON(iris.Map{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
