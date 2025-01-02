package handler

import (
	"fmt"
	"strconv"
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type LogicHandler struct {
	lgcUseCase usecase.LogicUseCase
}

func (lgcHandler LogicHandler) GetAllLogic(ctx *gin.Context) {
	idStr := ctx.Query("id")
	namaStr := ctx.Query("nama")
	emailStr := ctx.Query("email")
	genderStr := ctx.Query("gender")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	if idStr == "" && namaStr == "" && emailStr == "" && genderStr == "" {
		lgc, pagging, err := lgcHandler.lgcUseCase.GetAllLogic(page)
		if err != nil {
			fmt.Printf("LogicHandler.GetAllLogic(): %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching logic data",
			})
			return
		}
		ctx.HTML(http.StatusOK, "ui-tablejson.html", gin.H{
			"success": true,
			"page":    pagging,
			"data":    lgc,
		})
	} else if idStr != "" {
		lgcHandler.GetLogicById(ctx, idStr, page)
	} else if namaStr != "" {
		lgcHandler.GetLogicByName(ctx, namaStr, page)
	} else if emailStr != "" {
		lgcHandler.GetLogicByEmail(ctx, emailStr, page)
	} else if genderStr != "" {
		lgcHandler.GetLogicByGender(ctx, genderStr, page)
	}

}

func (lgcHandler LogicHandler) GetLogicById(ctx *gin.Context, id string, page int) {

	lgc, pagging, err := lgcHandler.lgcUseCase.GetLogicById(id, page)
	if err != nil {
		fmt.Printf("LogicHandler.GetAllLogic(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching logic data",
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablejson.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    lgc,
	})
}

func (lgcHandler LogicHandler) GetLogicByName(ctx *gin.Context, nama string, page int) {
	lgc, pagging, err := lgcHandler.lgcUseCase.GetLogicByName(nama, page)
	if err != nil {
		fmt.Printf("LogicHandler.GetAllLogic(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching logic data",
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablejson.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    lgc,
	})
}

func (lgcHandler LogicHandler) GetLogicByEmail(ctx *gin.Context, email string, page int) {
	lgc, pagging, err := lgcHandler.lgcUseCase.GetLogicByEmail(email, page)
	if err != nil {
		fmt.Printf("LogicHandler.GetAllLogic(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching logic data",
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablejson.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    lgc,
	})
}

func (lgcHandler LogicHandler) GetLogicByGender(ctx *gin.Context, gender string, page int) {
	lgc, pagging, err := lgcHandler.lgcUseCase.GetLogicByGender(gender, page)
	if err != nil {
		fmt.Printf("LogicHandler.GetAllLogic(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching logic data",
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablejson.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    lgc,
	})
}

func (lgcHandler LogicHandler) EditLogic(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	fname := ctx.Request.FormValue("fname")
	lname := ctx.Request.FormValue("lname")
	email := ctx.Request.FormValue("email")
	gender := ctx.Request.FormValue("gender")
	avatar := ctx.Request.FormValue("avatar")

	lgc := model.LogicModel{
		ID:         id,
		First_name: fname,
		Last_name:  lname,
		Email:      email,
		Gender:     gender,
		Avatar:     avatar,
	}

	err := lgcHandler.lgcUseCase.EditLogicById(&lgc)
	if err != nil {
		fmt.Println(err)
	}

	ctx.Redirect(http.StatusSeeOther, "/lgcu")
}

func (lgcHandler LogicHandler) DeleteLogic(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	all := ctx.Request.FormValue("all")

	if id != "" {
		err := lgcHandler.lgcUseCase.DeleteLogicById(id)
		if err != nil {
			fmt.Println(err)
		}
	}
	if all == "true" {
		err := lgcHandler.lgcUseCase.DeleteAllLogic()
		if err != nil {
			fmt.Println(err)
		}
	}
	ctx.Redirect(http.StatusSeeOther, "/lgcd")
}

func (lgcHandler LogicHandler) EditPage(ctx *gin.Context) {

	id := ctx.Request.FormValue("id")
	if id == "" {
		ctx.HTML(http.StatusOK, "ui-updatejson.html", gin.H{})
	} else {
		lgc, _, err := lgcHandler.lgcUseCase.GetLogicById(id, 1)
		if err != nil {
			ctx.HTML(http.StatusOK, "ui-updatejson.html", gin.H{})
		}
		ctx.HTML(http.StatusOK, "ui-updatejson.html", gin.H{
			"success": true,
			"page":    0,
			"data":    lgc,
		})
	}

}

func (lgcHandler LogicHandler) DeletePage(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	if ctx.Query("id") != "" {
		id = ctx.Query("id")
	}
	if id == "" {
		ctx.HTML(http.StatusOK, "ui-deletejson.html", gin.H{})
	} else {
		lgc, _, err := lgcHandler.lgcUseCase.GetLogicByIdNoExec(id, 1)
		if err != nil {
			ctx.HTML(http.StatusOK, "ui-deletejson.html", gin.H{})
		}
		ctx.HTML(http.StatusOK, "ui-deletejson.html", gin.H{
			"success": true,
			"page":    0,
			"data":    lgc,
		})
	}
}

func NewLogicHandler(srv *gin.Engine, lgcUseCase usecase.LogicUseCase) *LogicHandler {
	lgcHandler := &LogicHandler{
		lgcUseCase: lgcUseCase,
	}

	// route
	srv.POST("/lgcu", lgcHandler.EditLogic)
	srv.POST("/lgcd", lgcHandler.DeleteLogic)
	srv.GET("/lgcd", lgcHandler.DeletePage)
	srv.GET("/lgcu", lgcHandler.EditPage)
	srv.GET("/lgcs", lgcHandler.GetAllLogic)

	return lgcHandler
}
