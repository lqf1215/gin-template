package intercept

import (
	"fmt"
	"gin-template/config"
	"gin-template/global"
	"gin-template/model"
	"gin-template/pkg"
	"github.com/gin-gonic/gin"
)

// AuthApp 是一个中间件，用于保护路由
func AuthApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userId int64
			token  = c.GetHeader(config.LocalToken)
			db     = global.DB
			err    error
		)

		// 打印请求地址
		global.Log.Info("Request URL: " + c.Request.URL.Path)
		if token == "" || len(token) < 10 {
			pkg.MessageResponse(c, config.TOKEN_FAIL, "token is null", "令牌为null")
			c.Abort()
			return
		}

		user, err := model.UserSelectIdByToken(db, token)
		if err != nil {
			if err.Error() != "record not found" {
				fmt.Println(err.Error())
			}
			pkg.MessageResponse(c, config.TOKEN_FAIL, "token invalid or network error", "token失效或网络故障")
			c.Abort()
			return
		}
		userId = int64(user.ID)

		if !pkg.CheckSpecialCharacters(&token) {
			pkg.MessageResponse(c, config.TOKEN_FAIL, "token is invalid", "令牌无效")
			c.Abort()
			return
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&user.Token) {
			pkg.MessageResponse(c, config.TOKEN_FAIL, "token is exceed", "令牌超过")
			c.Abort()
			return
		}

		//刷新token有效时间
		if err = model.UserRefreshToken(db, userId, user.Token); err != nil {
			pkg.MessageResponse(c, config.TOKEN_FAIL, "db UserRefreshAppToken err", "刷新失败")
			c.Abort()
			return
		}

		c.Set(config.LocalUseridUint, uint(userId))
		c.Set(config.LocalUseridInt64, userId)
		c.Next()
	}
}

// AuthWebOperationProtected 用于保护网页操作的路由
func AuthWebOperationProtected(rights string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authority = c.GetHeader(config.LocalAuthority)
		)

		//刷新token有效时间
		if authority != "all" {
			if authority != rights {
				pkg.MessageResponse(c, config.OPERATION_FAIL, "un authorized operation", "无权限操作")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// ManagerAuthProtected 是用于保护管理员操作的路由
//func ManagerAuthProtected() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var (
//			token = c.GetHeader(config.LocalToken)
//			db    = global.DB
//			err   error
//		)
//
//		// 打印请求地址
//		global.Log.Info("Request URL: ", zap.Field{Key: c.Request.URL.Path})
//		if token == "" || len(token) < 10 {
//			c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
//			c.Abort()
//			return
//		}
//
//		user, err := model.UserSelectIdByManagerToken(db, token)
//		if err != nil {
//			c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
//			c.Abort()
//			return
//		}
//
//		if !pkg.CheckSpecialCharacters(&token) {
//			c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
//			c.Abort()
//			return
//		}
//		//检查token 有效时间
//		if !pkg.CheckTokenValidityTime(&user.ManagerToken) {
//			c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
//			c.Abort()
//			return
//		}
//
//		//刷新manager_token有效时间 要加一个时间戳
//		if err = model.UserRefreshManagerToken(db, int64(user.ID), user.ManagerToken); err != nil {
//			c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshManagerToken ", err.Error()))
//			c.Abort()
//			return
//		}
//		c.Set(config.AdminUseridInt64, int64(user.ID))
//		//c.Set(config.AdminUsername, user.Name)
//		//c.Set(config.ManageRole, user.Role)
//		c.Set(config.ManageUser, user)
//		c.Next()
//	}
//}
