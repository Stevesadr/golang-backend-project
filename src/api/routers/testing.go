package routers

import (
	"github.com/Stevesadr/golang-backend-project/api/handlers"
	"github.com/gin-gonic/gin"
)

func Testing(r *gin.RouterGroup){
	t := handlers.NewTesting()
	r.GET("/", t.TestingHandler)
	r.GET("/users", t.Users)
	r.GET("/user/:id", t.UserById)
	r.GET("/user/get-user-by-username/:username", t.UserByUsername)
	r.GET("/user/:id/account", t.Account)
	r.POST("/add-user", t.AddUser)

	r.GET("/binder/header1", t.HeaderBinderWithGetHeader)
	r.GET("/binder/header2", t.HeaderBinderWithBindHeader)

	r.GET("/binder/query1", t.QueryBinderWithGetQueryArrayAndGetQuery)
	r.GET("/binder/query2", t.QueryBinderWithGetQueryMap)

	r.GET("/binder/uli1/:id/:name", t.UliBinderWithParam)

	r.GET("/binder/body1", t.BodyBinderWithBindJson)

	r.GET("/binder/form1", t.FormBinderWithBind)

	r.GET("/binder/file1", t.FileBinderWithFormFile)
}