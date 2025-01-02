package handler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/usecase"
	"Golang_GIN_GORM/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type PostgreSqlHandler struct {
	pstsqlUseCase usecase.PostgreSqlUseCase
}

func (pstsqlHandler PostgreSqlHandler) GetAllPostgreSql(ctx *gin.Context) {
	idStr := ctx.Query("id")
	namaStr := ctx.Query("nama")
	emailStr := ctx.Query("email")
	genderStr := ctx.Query("gender")
	pageStr := ctx.Query("page")
	fmt.Println("parameter" + idStr)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		// Handle error, misalnya, tetapkan page ke nilai default 1
		page = 1
	}

	if idStr == "" && namaStr == "" && emailStr == "" && genderStr == "" {
		pstsql, pagging, err := pstsqlHandler.pstsqlUseCase.GetAllPostgreSql(page)
		if err != nil {
			fmt.Printf("PostgreSqlHandler.GetAllPostgreSql(): %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": err,
			})
			return
		}
		ctx.HTML(http.StatusOK, "ui-tablepostgresql.html", gin.H{
			"success": true,
			"page":    pagging,
			"data":    pstsql,
		})
	} else if idStr != "" {
		pstsqlHandler.GetPostgreSqlById(ctx, idStr, page)
	} else if namaStr != "" {
		pstsqlHandler.GetPostgreSqlByName(ctx, namaStr, page)
	} else if emailStr != "" {
		pstsqlHandler.GetPostgreSqlByEmail(ctx, emailStr, page)
	} else if genderStr != "" {
		pstsqlHandler.GetPostgreSqlByGender(ctx, genderStr, page)
	}

}

func (pstsqlHandler PostgreSqlHandler) GetPostgreSqlById(ctx *gin.Context, id string, page int) {

	pstsql, pagging, err := pstsqlHandler.pstsqlUseCase.GetPostgreSqlById(id, page)
	if err != nil {
		fmt.Printf("PostgreSqlHandler.GetAllPostgreSql(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablepostgresql.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    pstsql,
	})
}

func (pstsqlHandler PostgreSqlHandler) GetPostgreSqlByName(ctx *gin.Context, nama string, page int) {
	pstsql, pagging, err := pstsqlHandler.pstsqlUseCase.GetPostgreSqlByName(nama, page)
	if err != nil {
		fmt.Printf("PostgreSqlHandler.GetAllPostgreSql(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablepostgresql.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    pstsql,
	})
}

func (pstsqlHandler PostgreSqlHandler) GetPostgreSqlByEmail(ctx *gin.Context, email string, page int) {
	pstsql, pagging, err := pstsqlHandler.pstsqlUseCase.GetPostgreSqlByEmail(email, page)
	if err != nil {
		fmt.Printf("PostgreSqlHandler.GetAllPostgreSql(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablepostgresql.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    pstsql,
	})
}

func (pstsqlHandler PostgreSqlHandler) GetPostgreSqlByGender(ctx *gin.Context, gender string, page int) {
	pstsql, pagging, err := pstsqlHandler.pstsqlUseCase.GetPostgreSqlByGender(gender, page)
	if err != nil {
		fmt.Printf("PostgreSqlHandler.GetAllPostgreSql(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.HTML(http.StatusOK, "ui-tablepostgresql.html", gin.H{
		"success": true,
		"page":    pagging,
		"data":    pstsql,
	})
}

func (pstsqlHandler PostgreSqlHandler) InsertPostgreSql(ctx *gin.Context) {
	file, err := ctx.FormFile("fileUpload")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Open the CSV file
	fileContent, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	defer fileContent.Close()

	// Read CSV file using encoding/csv library
	reader := csv.NewReader(fileContent)

	// Read all lines from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Process CSV data (e.g., print to console)
	var jsonData []*model.LogicModel
	headers := records[0]
	for _, record := range records[1:] {
		entry := &model.LogicModel{}
		for i, value := range record {
			switch headers[i] {
			case "id":
				entry.ID = value
			case "first_name":
				entry.First_name = value
			case "last_name":
				entry.Last_name = value
			case "email":
				entry.Email = value
			case "gender":
				entry.Gender = value
			case "avatar":
				entry.Avatar = value
			}
		}
		jsonData = append(jsonData, entry)
	}
	fmt.Println(len(jsonData))
	err = pstsqlHandler.pstsqlUseCase.InsertPostgreSql(jsonData)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("PostgreSqlHandler.InsertPostgreSql() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("PostgreSqlHandler.InsertPostgreSql() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving logic2 data",
			})
		}
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/main")
}

func (pstsqlHandler PostgreSqlHandler) EditPostgreSql(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	fname := ctx.Request.FormValue("fname")
	lname := ctx.Request.FormValue("lname")
	email := ctx.Request.FormValue("email")
	gender := ctx.Request.FormValue("gender")
	avatar := ctx.Request.FormValue("avatar")

	pstsql := model.LogicModel{
		ID:         id,
		First_name: fname,
		Last_name:  lname,
		Email:      email,
		Gender:     gender,
		Avatar:     avatar,
	}

	err := pstsqlHandler.pstsqlUseCase.EditPostgreSqlById(&pstsql)
	if err != nil {
		fmt.Println(err)
	}

	ctx.Redirect(http.StatusSeeOther, "/pstsqlu")
}

func (pstsqlHandler PostgreSqlHandler) DeletePostgreSql(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	all := ctx.Request.FormValue("all")

	if id != "" {
		err := pstsqlHandler.pstsqlUseCase.DeletePostgreSqlById(id)
		if err != nil {
			fmt.Println(err)
		}
	}
	if all == "true" {
		err := pstsqlHandler.pstsqlUseCase.DeleteAllPostgreSql()
		if err != nil {
			fmt.Println(err)
		}
	}
	ctx.Redirect(http.StatusSeeOther, "/pstsqld")
}

func (pstsqlHandler PostgreSqlHandler) EditPageP(ctx *gin.Context) {

	id := ctx.Request.FormValue("id")
	if id == "" {
		ctx.HTML(http.StatusOK, "ui-updatepostgresql.html", gin.H{})
	} else {
		pstsql, _, err := pstsqlHandler.pstsqlUseCase.GetPostgreSqlById(id, 1)
		if err != nil {
			ctx.HTML(http.StatusOK, "ui-updatepostgresql.html", gin.H{})
		} else {
			ctx.HTML(http.StatusOK, "ui-updatepostgresql.html", gin.H{
				"success": true,
				"page":    0,
				"data":    pstsql,
			})
		}

	}

}

func (pstsqlHandler PostgreSqlHandler) DeletePageP(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	if ctx.Query("id") != "" {
		id = ctx.Query("id")
	}
	if id == "" {
		ctx.HTML(http.StatusOK, "ui-deletepostgresql.html", gin.H{})
	} else {
		pstsql, _, err := pstsqlHandler.pstsqlUseCase.GetPostgreSqlByIdNoExec(id, 1)
		if err != nil {
			ctx.HTML(http.StatusOK, "ui-deletepostgresql.html", gin.H{})
		} else {
			ctx.HTML(http.StatusOK, "ui-deletepostgresql.html", gin.H{
				"success": true,
				"page":    0,
				"data":    pstsql,
			})
		}

	}
}

func NewPostgreSqlHandler(srv *gin.Engine, pstsqlUseCase usecase.PostgreSqlUseCase) *PostgreSqlHandler {
	pstsqlHandler := &PostgreSqlHandler{
		pstsqlUseCase: pstsqlUseCase,
	}

	// route
	srv.POST("/pstsql", pstsqlHandler.InsertPostgreSql)
	srv.POST("/pstsqlu", pstsqlHandler.EditPostgreSql)
	srv.POST("/pstsqld", pstsqlHandler.DeletePostgreSql)
	srv.GET("/pstsqld", pstsqlHandler.DeletePageP)
	srv.GET("/pstsqlu", pstsqlHandler.EditPageP)
	srv.GET("/pstsqls", pstsqlHandler.GetAllPostgreSql)

	return pstsqlHandler
}
