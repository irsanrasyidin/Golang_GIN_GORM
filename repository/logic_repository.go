package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/utils"
)

type LogicRepo interface {
	InsertLogic(lgc []*model.LogicModel) error
	GetLogicById(id string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetAllLogic(page int) ([]*model.LogicModel, *model.Pagination, error)
	EditLogicById(lgc *model.LogicModel) error
	DeleteLogicById(id string) error
	DeleteAllLogic() error
	MainLogicExec(exec *model.ExecutionModel) error
}

type logicRepoImpl struct {
}

func (lgcRepo *logicRepoImpl) InsertLogic(lgc []*model.LogicModel) error {
	jsonFile, err := os.Create("output.json")
	if err != nil {
		return &utils.AppError{
			ErrorCode:    101,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(lgc)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    102,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (lgcRepo *logicRepoImpl) GetLogicById(id string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var pagging model.Pagination
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    201,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    202,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		if jsonData[i].ID == id {
			limitData = append(limitData, &jsonData[i])
		}
	}
	pagging.Page = page
	pagging.PageSize = 1
	pagging.TotalItems = len(jsonData)
	return limitData, &pagging, nil
}

func (lgcRepo *logicRepoImpl) GetLogicByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var pagging model.Pagination
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    301,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    302,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData []model.LogicModel
	var first []model.LogicModel
	var last []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		if strings.Contains(jsonData[i].First_name, nama) {
			first = append(first, jsonData[i])
		} else if strings.Contains(jsonData[i].Last_name, nama) {
			last = append(last, jsonData[i])
		}
	}
	for i := 0; i < len(first); i++ {
		limitData = append(limitData, &first[i])
	}
	for i := 0; i < len(last); i++ {
		limitData = append(limitData, &last[i])
	}
	pagging.Page = page
	pagging.PageSize = 1
	pagging.TotalItems = len(jsonData)
	return limitData, &pagging, nil
}

func (lgcRepo *logicRepoImpl) GetLogicByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var pagging model.Pagination
	// limit := 20 * page
	// g := limit - 20
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    401,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    402,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData, searchData []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		if strings.Contains(jsonData[i].Email, email) {
			searchData = append(searchData, jsonData[i])
		}
	}
	for i := 0; i < len(searchData); i++ {
		limitData = append(limitData, &searchData[i])
	}

	pagging.Page = page
	pagging.PageSize = 1
	pagging.TotalItems = len(jsonData)
	return limitData, &pagging, nil
}

func (lgcRepo *logicRepoImpl) GetLogicByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var pagging model.Pagination
	// limit := 20 * page
	// g := limit - 20
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    501,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    502,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData, searchData []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		if jsonData[i].Gender == gender {
			limitData = append(limitData, &jsonData[i])
		}
	}
	for i := 0; i < len(searchData); i++ {
		limitData = append(limitData, &searchData[i])
	}
	pagging.Page = page
	pagging.PageSize = 1
	pagging.TotalItems = len(jsonData)
	return limitData, &pagging, nil
}

func (lgcRepo *logicRepoImpl) GetAllLogic(page int) ([]*model.LogicModel, *model.Pagination, error) {
	var pagging model.Pagination
	// limit := 20 * page
	// g := limit - 20
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    601,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    602,
			ErrorMessage: err.Error(),
		}
	}

	var jsonData []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		limitData = append(limitData, &jsonData[i])
	}
	pagging.Page = page
	pagging.PageSize = (len(jsonData) / 20)
	pagging.TotalItems = len(jsonData)
	return limitData, &pagging, nil
}

func (lgcRepo *logicRepoImpl) EditLogicById(lgc *model.LogicModel) error {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return &utils.AppError{
			ErrorCode:    701,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    702,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		if jsonData[i].ID == lgc.ID {
			limitData = append(limitData, lgc)
		} else {
			limitData = append(limitData, &jsonData[i])
		}

	}

	//fmt.Println(limitData)
	err = lgcRepo.InsertLogic(limitData)
	if err != nil {
		return err
	}
	return nil
}

func (lgcRepo *logicRepoImpl) DeleteLogicById(id string) error {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return &utils.AppError{
			ErrorCode:    801,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    802,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData []model.LogicModel
	var limitData []*model.LogicModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		if jsonData[i].ID != id {
			limitData = append(limitData, &jsonData[i])
		}
	}

	//fmt.Println(limitData)
	err = lgcRepo.InsertLogic(limitData)
	if err != nil {
		return err
	}
	return nil
}

func (lgcRepo *logicRepoImpl) DeleteAllLogic() error {
	err := lgcRepo.InsertLogic(nil)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    902,
			ErrorMessage: err.Error(),
		}

	}
	return nil
}

func (lgcRepo *logicRepoImpl) InsertLogicExe(exec []*model.ExecutionModel) error {
	jsonFile, err := os.Create("execution.json")
	if err != nil {
		return &utils.AppError{
			ErrorCode:    1001,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(exec)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    1002,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (lgcRepo *logicRepoImpl) MainLogicExec(exec *model.ExecutionModel) error {
	coba := 1
	var existData, sortData []*model.ExecutionModel
	jsonFile, err := os.Open("execution.json")
	if err != nil {
		return &utils.AppError{
			ErrorCode:    801,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()
	existData, err = lgcRepo.GetLogicExec(jsonFile)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    802,
			ErrorMessage: err.Error(),
		}
	}
	for _, value := range existData {
		if value.Nama == exec.Nama {
			coba++
		}
	}
	if coba == 31 {
		exec.Coba = 30
		sortData = lgcRepo.SortLogicExec(existData, exec)
		err = lgcRepo.InsertLogicExe(sortData)
		if err != nil {
			return &utils.AppError{
				ErrorCode:    803,
				ErrorMessage: err.Error(),
			}
		}
		//fmt.Println("Selesai sortir ", sortData)
	} else {
		exec.Coba = coba
		existData = append(existData, exec)
		err = lgcRepo.InsertLogicExe(existData)
		if err != nil {
			return &utils.AppError{
				ErrorCode:    804,
				ErrorMessage: err.Error(),
			}
		}
		//fmt.Println("tidak ada sortir ", existData)
	}
	return nil
}

func (lgcRepo *logicRepoImpl) SortLogicExec(existData []*model.ExecutionModel, newData *model.ExecutionModel) []*model.ExecutionModel {
	var sortData []*model.ExecutionModel
	for _, value := range existData {
		if value.Nama == newData.Nama {
			switch value.Coba {
			case 1:
				continue
			case 30:
				value.Coba = value.Coba - 1
				sortData = append(sortData, value)
				sortData = append(sortData, newData)
			default:
				value.Coba = value.Coba - 1
				sortData = append(sortData, value)
			}
		} else {
			sortData = append(sortData, value)
		}
	}
	fmt.Println("sortir ", sortData)
	return sortData
}

func (lgcRepo *logicRepoImpl) GetLogicExec(file *os.File) ([]*model.ExecutionModel, error) {
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    1202,
			ErrorMessage: err.Error(),
		}
	}
	var jsonData []model.ExecutionModel
	var limitData []*model.ExecutionModel
	json.Unmarshal(byteValue, &jsonData)
	//fmt.Println(len(jsonData))
	for i := 0; i < len(jsonData); i++ {
		limitData = append(limitData, &jsonData[i])
	}

	return limitData, nil
}

func NewLogicRepo() LogicRepo {
	return &logicRepoImpl{}
}
