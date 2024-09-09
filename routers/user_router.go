package routers

import (
	"GVB_server/api"
	"GVB_server/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("KG87gfGg8gsr7G78HL7gljB7sh22"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserAPi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", app.EmailLoginView)
	router.POST("login", app.QQLoginView)
	router.POST("users", middleware.JwtAuth(), app.UserCreateView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAdmin(), app.UserBindEmailView)

}
