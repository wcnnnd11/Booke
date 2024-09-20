package routers

import "GVB_server/api"

func (router RouterGroup) DataRouter() {
	app := api.ApiGroupApp.DataApi
	router.GET("data_login", app.SevenLoginView)
	router.GET("data_sum", app.DataSumView)

}
