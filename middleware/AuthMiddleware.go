package middleware

import (
	"gin_demo2/common"
	"gin_demo2/model"
	"gin_demo2/response"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		log.Println("ctx111=", tokenString)
		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {

			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		log.Println("ctx222=", tokenString)
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}
		log.Println("ctx333=", claims)
		// 验证通过后获取claim中的userId
		userId := claims.UserId
		log.Println("ctx444=", claims.UserId)
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		log.Println("ctx555=", user)
		// 用户不存在
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")

			ctx.Abort()
			return
		}

		// 用户存在 将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()

	}

}
