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

// @Summary Start Testing
// @Description This handler Just for test
// @Accept json
// @Produce json
// @Tags Test
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test [get]
func (t *Testing)TestingHandler(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("TestHandler", true, 0))
}

// @Summary Start Testing
// @Description This handler Just another handler for test
// @Accept json
// @Produce json
// @Tags Test
// @Success 200 {object} helper.BaseResponse "Success"
// @Security BearerAuth
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/users [get]
func (t *Testing)Users(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("Users", true, 0))
}

// @Summary Test Get id
// @Description This handler want to get id from the path
// @Accept json
// @Produce json
// @Tags Test
// @Param id path string true "User Id"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/users/{id} [get]
func (t *Testing)UserById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result" : "TestingById",
		"id": id,
	}, true, 0))
}

// @Summary Test Get username
// @Description This handler want to get username from the path
// @Accept json
// @Produce json
// @Tags Test
// @Param username path string true "User username"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/user/get-user-by-username/{username} [get]
func (t *Testing)UserByUsername(c *gin.Context){
	username := c.Param("username")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result" : "TestingByUsername",
		"username": username,
	}, true, 0))
}

// @Summary Testing username account
// @Description This handler Just for get account from the path with id
// @Accept json
// @Produce json
// @Tags Test
// @Param id path string true "User username account id"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/user/{id}/account [get]
func(t *Testing)Account(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("Account", true, 0))
}

// @Summary Testing Post
// @Description This handler Just for test post
// @Accept json
// @Produce json
// @Tags Test
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/user/add-user [post]
func(t *Testing)AddUser(c *gin.Context){
	c.JSON(http.StatusOK, helper.GenerateResponse("AddUser", true, 0))
}

// @Summary Testing Get user token
// @Description This handler Just get token from header
// @Accept json
// @Produce json
// @Tags Test
// @Param token header string true "token"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/header1 [get]
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
// @Summary Testing Get User Token and Browser
// @Description This handler Just get token and browser from header
// @Accept json
// @Produce json
// @Tags Test
// @Param token header string true "token"
// @Param browser header string false "browser"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/header2 [get]
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

// @Summary Testing Params query
// @Description This handler Just for get page and tag from query
// @Accept json
// @Produce json
// @Tags Test
// @Param page query string false "get page" default(1)
// @Param tag query string false "get tag" default(golang)
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/query1 [get]
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

// @Summary Testing User query
// @Description This handler Just for get user data from query
// @Accept json
// @Produce json
// @Tags Test
// @Param user query string false "get user data" 
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/query2 [get]
func(t *Testing)QueryBinderWithGetQueryMap(c *gin.Context){
	// The query most be like this : /users?user[name]=ali&user[age]=20
	user, _:= c.GetQueryMap("user")
	c.JSON(http.StatusOK, helper.GenerateResponse(gin.H{
		"result": "QueryBinderWithGetQueryMap",
		"user": user,
		}, true, 0),
	)	
}

// @Summary Testing Get Id and Name
// @Description This handler get id and name from path together
// @Accept json
// @Produce json
// @Tags Test
// @Param id path string true "id"
// @Param name path string true "name"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/uli1/{id}/{name} [get]
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
// @Summary Testing for return response
// @Description this handler hear to get response from the request
// @Accept json
// @Produce json
// @Tags Test
// @Param person1 body person true "person data"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/body1 [post]
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

// @Summary Testing for return response from form
// @Description this handler hear to get response from the request as form
// @Accept json
// @Produce json
// @Tags Test
// @Param person formData person false "person data"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/form1 [post]
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

// @Summary Testing for return response from file
// @Description this handler hear to get response from the request as file
// @Accept json
// @Produce json
// @Tags Test
// @Param file formData file false "file data"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} []validations.ValidationError "Failure"
// @Router /v1/test/binder/file1 [post]
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