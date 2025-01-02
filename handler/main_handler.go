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

type MainHandler struct {
	mainUseCase usecase.MainUseCase
}

func (mainHandler MainHandler) InsertMain(ctx *gin.Context) {
	file, err := ctx.FormFile("fileUpload")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	defer fileContent.Close()

	reader := csv.NewReader(fileContent)

	records, err := reader.ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

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
	//fmt.Println(len(jsonData))
	err = mainHandler.mainUseCase.InsertMain(jsonData)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MainHandler.InsertMain() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MainHandler.InsertMain() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving logic data",
			})
		}
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/main")
}

func (mainHandler MainHandler) HomePage(ctx *gin.Context) {
	option := ctx.Request.FormValue("option")
	optionInt, _ := strconv.Atoi(option)
	//fmt.Println(optionInt)
	mainData, postgreData, jsonData, err := mainHandler.mainUseCase.MainPostgreSQL(optionInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	for _, v := range jsonData {
		for _, v2 := range mainData {
			if v.Nama+"JSON" == v2.Nama {
				for _, v3 := range v.OutlinerData {
					if v3 == v2.Duration {
						v2.Status = "out"
					}
				}
				if v.Top == v2.Duration {
					v2.Status = "top"
				}
			}
			//fmt.Println(v2)
		}
	}
	for _, v := range postgreData {
		for _, v2 := range mainData {
			if v.Nama+"PostgreSQL" == v2.Nama {
				for _, v3 := range v.OutlinerData {
					if v3 == v2.Duration {
						v2.Status = "out"
					}
				}
				if v.Top == v2.Duration {
					v2.Status = "top"
				}
			}
			//fmt.Println(v2)
		}
	}

	// for _, v := range mainData {
	// 	fmt.Println("main",v)
	// }
	// for _, v := range postgreData {
	// 	fmt.Println("psql",v)
	// }
	// for _, v := range jsonData {
	// 	fmt.Println("json",v)
	// }

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"mainData":    mainData,
		"postgreData": postgreData,
		"jsonData":    jsonData,
	})
}

func (mainHandler MainHandler) Restart(ctx *gin.Context) {
	err := mainHandler.mainUseCase.Restart()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/main")
}

func NewMainHandler(srv *gin.Engine, mainUseCase usecase.MainUseCase) *MainHandler {
	mainHandler := &MainHandler{
		mainUseCase: mainUseCase,
	}

	// route
	srv.POST("/lgc", mainHandler.InsertMain)
	srv.GET("/main", mainHandler.HomePage)
	srv.POST("/restart", mainHandler.Restart)

	return mainHandler
}
