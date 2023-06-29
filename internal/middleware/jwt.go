package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hd2yao/blog/pkg/app"
	"github.com/hd2yao/blog/pkg/err_code"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			token string
			ecode = err_code.Success
		)
		if s, exist := context.GetQuery("token"); exist {
			token = s
		} else {
			token = context.GetHeader("token")
		}
		if token == "" {
			ecode = err_code.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = err_code.UnauthorizedTokenTimeout
				default:
					ecode = err_code.UnauthorizedTokenError
				}
			}
		}

		if ecode != err_code.Success {
			response := app.NewResponse(context)
			response.ToErrorResponse(ecode)
			context.Abort()
			return
		}

		context.Next()
	}
}
