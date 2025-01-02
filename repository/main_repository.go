package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/utils"

	"gorm.io/gorm"
)

type MainRepo interface {
	MainPostgreSQL(option int) ([]*model.ExecutionModel, []*model.ExecutionResultModel, error)
	MainJSON(option int) ([]*model.ExecutionModel, []*model.ExecutionResultModel, error)
	Restart() error
}

type mainRepoImpl struct {
	db *gorm.DB
}

func (mainRepo *mainRepoImpl) MainPostgreSQL(option int) ([]*model.ExecutionModel, []*model.ExecutionResultModel, error) {
	var a, b int
	mainData, err := mainRepo.GetAllPostgreSQL()
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    121,
			ErrorMessage: err.Error(),
		}
	}

	switch option {
	case 2:
		a = 0
		b = 1
	case 3:
		a = 1
		b = 6
	case 4:
		a = 6
		b = 7
	case 5:
		a = 7
		b = 9
	default:
		a = 0
		b = 9
	}
	resultData, newData, err := mainRepo.ProcessPostgreSQL(mainData, a, b)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    221,
			ErrorMessage: err.Error(),
		}
	}
	if newData != nil {
		return newData, resultData, nil
	}
	return mainData, resultData, nil
}

func (mainRepo *mainRepoImpl) GetAllPostgreSQL() ([]*model.ExecutionModel, error) {
	var result []*model.ExecutionModel
	err := mainRepo.db.Order("nama, coba").Find(&result).Error
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    221,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil
}

func (mainRepo *mainRepoImpl) ProcessPostgreSQL(mainData []*model.ExecutionModel, a int, b int) ([]*model.ExecutionResultModel, []*model.ExecutionModel, error) {
	var resultName string
	var resultData []*model.ExecutionResultModel
	var newData []*model.ExecutionModel
	for i := a; i < b; i++ {
		switch i {
		case 0:
			resultName = "Insert"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 1:
			resultName = "GetByID"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 2:
			resultName = "GetByName"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 3:
			resultName = "GetByEmail"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 4:
			resultName = "GetByGender"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 5:
			resultName = "GetAll"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 6:
			resultName = "UpdateByID"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 7:
			resultName = "DeleteByID"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 8:
			resultName = "DeleteAll"
			result, abc := mainRepo.ResultPostgreSQL(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		}
	}
	if a != 0 || b != 9 {
		return resultData, newData, nil
	}
	return resultData, nil, nil
}

func (mainRepo *mainRepoImpl) ResultPostgreSQL(name string, mainData []*model.ExecutionModel) (*model.ExecutionResultModel, []*model.ExecutionModel) {
	var resultData model.ExecutionResultModel
	var newData []*model.ExecutionModel
	resultData.Nama = name
	for _, value := range mainData {
		if value.Nama == name+"PostgreSQL" {
			if value.Duration < resultData.Top && resultData.Top != 0 {
				resultData.Top = value.Duration
			}
			if resultData.Top == 0 {
				resultData.Top = value.Duration
			}
			newData = append(newData, value)
			resultData.Data = append(resultData.Data, value.Duration)
		}
	}
	var deletedData float64
	for {
		resultData.S_Deviasi = utils.StandardDeviation(resultData.Data)
		if resultData.S_Deviasi > 10 {
			resultData.Data, resultData.Outliner, deletedData = utils.Filter(resultData.Data, resultData.Outliner)
			resultData.OutlinerData = append(resultData.OutlinerData, deletedData)
		} else {
			break
		}
	}

	resultData.Average = utils.AverageDuration(resultData.Data)
	return &resultData, newData
}

func (mainRepo *mainRepoImpl) MainJSON(option int) ([]*model.ExecutionModel, []*model.ExecutionResultModel, error) {
	var a, b int
	mainData, err := mainRepo.GetAllJSON()
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    121,
			ErrorMessage: err.Error(),
		}
	}
	switch option {
	case 2:
		a = 0
		b = 1
	case 3:
		a = 1
		b = 6
	case 4:
		a = 6
		b = 7
	case 5:
		a = 7
		b = 9
	default:
		a = 0
		b = 9
	}
	resultData, newData, err := mainRepo.ProcessJSON(mainData, a, b)
	if err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    221,
			ErrorMessage: err.Error(),
		}
	}
	if newData != nil {
		return newData, resultData, nil
	}
	return nil, resultData, nil
}

func (mainRepo *mainRepoImpl) GetAllJSON() ([]*model.ExecutionModel, error) {
	jsonFile, err := os.Open("execution.json")
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    601,
			ErrorMessage: err.Error(),
		}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    602,
			ErrorMessage: err.Error(),
		}
	}

	var jsonData []model.ExecutionModel
	var limitData []*model.ExecutionModel
	json.Unmarshal(byteValue, &jsonData)
	for i := 0; i < len(jsonData); i++ {
		limitData = append(limitData, &jsonData[i])
	}
	return limitData, nil
}

func (mainRepo *mainRepoImpl) ProcessJSON(mainData []*model.ExecutionModel, a int, b int) ([]*model.ExecutionResultModel, []*model.ExecutionModel, error) {
	var resultName string
	var resultData []*model.ExecutionResultModel
	var newData []*model.ExecutionModel
	for i := a; i < b; i++ {
		switch i {
		case 0:
			resultName = "Insert"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 1:
			resultName = "GetByID"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 2:
			resultName = "GetByName"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 3:
			resultName = "GetByEmail"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 4:
			resultName = "GetByGender"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 5:
			resultName = "GetAll"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 6:
			resultName = "UpdateByID"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 7:
			resultName = "DeleteByID"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		case 8:
			resultName = "DeleteAll"
			result, abc := mainRepo.ResultJSON(resultName, mainData)
			newData = append(newData, abc...)
			resultData = append(resultData, result)
		}
	}
	if a != 0 || b != 9 {
		return resultData, newData, nil
	}
	return resultData, nil, nil
}

func (mainRepo *mainRepoImpl) ResultJSON(name string, mainData []*model.ExecutionModel) (*model.ExecutionResultModel, []*model.ExecutionModel) {
	var resultData model.ExecutionResultModel
	var newData []*model.ExecutionModel
	resultData.Nama = name
	for _, value := range mainData {
		if value.Nama == name+"JSON" {
			if value.Duration < resultData.Top && resultData.Top != 0 {
				resultData.Top = value.Duration
			}
			if resultData.Top == 0 {
				resultData.Top = value.Duration
			}
			newData = append(newData, value)
			resultData.Data = append(resultData.Data, value.Duration)
		}
	}
	var deletedData float64
	for {
		resultData.S_Deviasi = utils.StandardDeviation(resultData.Data)
		if resultData.S_Deviasi > 10 {
			resultData.Data, resultData.Outliner, deletedData = utils.Filter(resultData.Data, resultData.Outliner)
			resultData.OutlinerData = append(resultData.OutlinerData, deletedData)
		} else {
			break
		}
	}

	resultData.Average = utils.AverageDuration(resultData.Data)
	return &resultData, newData
}

func (mainRepo *mainRepoImpl) Restart() error {
	dataJSON, err := mainRepo.GetAllJSON()
	if err != nil {
		return err
	}
	sort.Slice(dataJSON, func(i, j int) bool {
		if dataJSON[i].Nama == dataJSON[j].Nama {
			return dataJSON[i].Coba < dataJSON[j].Coba
		}
		return dataJSON[i].Nama < dataJSON[j].Nama
	})

	file, err := os.Create("sorted_people.csv")
	if err != nil {
		fmt.Println("Tidak bisa membuat file CSV:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Nama", "Masuk", "Keluar", "Duration", "Coba"}
	writer.Write(header)

	for _, record := range dataJSON {
		durationString := fmt.Sprintf("%.3f", record.Duration)
		cobaString := strconv.Itoa(record.Coba)

		recordSlice := []string{
			record.Nama,
			record.Masuk,
			record.Keluar,
			durationString,
			cobaString,
		}

		if err := writer.Write(recordSlice); err != nil {
			return &utils.AppError{
				ErrorCode:    803,
				ErrorMessage: err.Error(),
			}
		}
	}

	return nil
}

func NewMainRepo(db *gorm.DB) MainRepo {
	return &mainRepoImpl{db: db}
}
