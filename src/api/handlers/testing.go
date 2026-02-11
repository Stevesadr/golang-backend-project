package handlers

import (
	"net/http"

	"github.com/Stevesadr/golang-backend-project/api/helper"
	"github.com/gin-gonic/gin"
)

type Testing struct{}

func NewTesting() *Testing{
	return &Testing{}
}

func (t *Testing)TestingHandler(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("TestHandler", true, 0))
}

func (t *Testing)Users(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("Users", true, 0))
}

func (t *Testing)UserById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result" : "TestingById",
		"id": id,
	}, true, 0))
}

func (t *Testing)UserByUsername(c *gin.Context){
	username := c.Param("username")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result" : "TestingByUsername",
		"username": username,
	}, true, 0))
}

func(t *Testing)Account(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("Account", true, 0))
}

func(t *Testing)AddUser(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("AddUser", true, 0))
}

func(t *Testing)HeaderBinderWithGetHeader(c *gin.Context){
	h := c.GetHeader("token")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result" : "HeaderBinderWithGetHeader",
		"token": h,
	}, true, 0))
} 

type headerBind struct{
	Token string
	Browser string
}
func(t *Testing)HeaderBinderWithBindHeader(c *gin.Context){
	headerData := headerBind{}
	err := c.ShouldBindHeader(&headerData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, false, 1, err))
	}
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
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "QueryBinderWithGetQueryArrayAndGetQuery",
		"page": page,
		"tags": tags,
		}, true, 0),
	)	
}
func(t *Testing)QueryBinderWithGetQueryMap(c *gin.Context){
	// The query most be like this : /users?user[name]=ali&user[age]=20
	user, _:= c.GetQueryMap("user")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "QueryBinderWithGetQueryMap",
		"user": user,
		}, true, 0),
	)	
}

func(t *Testing)UliBinderWithParam(c *gin.Context){
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "UliBinderWithParam",
		"id": id,
		"name": name,
		}, true, 0),
	)		
}

type person struct{
	FirstName string `binding:"required"`
	LastName string
	Mobile string `binding:"required,mobile"`
}
func(t *Testing)BodyBinderWithBindJson(c *gin.Context){
	p := person{}
	err := c.ShouldBindJSON(&p) 
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithValidation(nil, false, 1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "BodyBinderWithBindJson",
		"data": p,
		"error": err,
		}, true, 0),
	)			
}

func(t *Testing)FormBinderWithBind(c *gin.Context){
	p := person{}
	err := c.Bind(&p)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusNotAcceptable, helper.GenerateResponseWithValidation(nil, false, 1, err))
	}
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "BodyBinderWithBindJson",
		"data": p,
	}, true, 0))		
}

func(t *Testing)FileBinderWithFormFile(c *gin.Context){
	file, _:= c.FormFile("file")
	err := c.SaveUploadedFile(file,"file")
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateResponseWithError(nil, false, 1, err))
	}
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "FileBinderWithFormFile",
		"file": file.Filename,
		}, true, 0),
	)				
}