package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Testing struct{}

func NewTesting() *Testing{
	return &Testing{}
}

func (t *Testing)TestingHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result" : "TestHandler",
	})
}

func (t *Testing)Users(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

func (t *Testing)UserById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "TestingById",
		"id": id,
	})
}

func (t *Testing)UserByUsername(c *gin.Context){
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result": "TestingByUsername",
		"username": username,
	})
}

func(t *Testing)Account(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result": "Account",
	})
}

func(t *Testing)AddUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
	})
}

func(t *Testing)HeaderBinderWithGetHeader(c *gin.Context){
	h := c.GetHeader("token")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinderWithGetHeader",
		"token": h,
	})
} 

type headerBind struct{
	Token string
	Browser string
}
func(t *Testing)HeaderBinderWithBindHeader(c *gin.Context){
	headerData := headerBind{}
	err := c.BindHeader(&headerData)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinderWithBindHeader",
		"error": err,
		"data": headerData,
	})
}

func(t *Testing)QueryBinderWithGetQueryArrayAndGetQuery(c *gin.Context){
	// The query most be like this : /users?page=1&tag=go&tag=api
	page, _:= c.GetQuery("page")  
	tags, _:= c.GetQueryArray("tag")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinderWithGetQueryArrayAndGetQuery",
		"page":page,
		"tags": tags,
	})
}
func(t *Testing)QueryBinderWithGetQueryMap(c *gin.Context){
	// The query most be like this : /users?user[name]=ali&user[age]=20
	user, _:= c.GetQueryMap("user")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinderWithGetQueryMap",
		"user":user,
	})	
}

func(t *Testing)UliBinderWithParam(c *gin.Context){
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "UliBinderWithParam",
		"id":id,
		"name": name,
	})		
}

type person struct{
	FirstName string
	LastName string
}
func(t *Testing)BodyBinderWithBindJson(c *gin.Context){
	p := person{}
	err := c.BindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinderWithBindJson",
		"data": p,
		"error": err,
	})			
}

func(t *Testing)FormBinderWithBind(c *gin.Context){
	p := person{}
	err := c.Bind(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinderWithBindJson",
		"data": p,
		"error": err,
	})		
}

func(t *Testing)FileBinderWithFormFile(c *gin.Context){
	file, _:= c.FormFile("file")
	c.SaveUploadedFile(file,"file")
	c.JSON(http.StatusOK, gin.H{
		"result": "FileBinderWithFormFile",
		"file": file.Filename,
	})			
}