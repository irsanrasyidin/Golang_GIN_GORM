package usecase

import (
	"fmt"
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type LogicUseCase interface {
	InsertLogic(lgc []*model.LogicModel, ctx *gin.Context) error
	GetLogicById(id string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByIdNoExec(id string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetLogicByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetAllLogic(page int) ([]*model.LogicModel, *model.Pagination, error)
	EditLogicById(lgc *model.LogicModel) error
	DeleteLogicById(id string) error
	DeleteAllLogic() error
}

type logicUseCaseImpl struct {
	lgcRepo    repository.LogicRepo
	pstsqlRepo repository.PostgreSqlRepo
}

func (lgcUseCase *logicUseCaseImpl) GetAllLogic(page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetAllJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := lgcUseCase.lgcRepo.GetAllLogic(page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (lgcUseCase *logicUseCaseImpl) GetLogicById(id string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetByIDJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := lgcUseCase.lgcRepo.GetLogicById(id, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (lgcUseCase *logicUseCaseImpl) GetLogicByIdNoExec(id string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	data, page2, err := lgcUseCase.lgcRepo.GetLogicById(id, page)
	if err != nil {
		return nil, nil, err
	}

	return data, page2, nil
}

func (lgcUseCase *logicUseCaseImpl) GetLogicByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetByNameJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := lgcUseCase.lgcRepo.GetLogicByName(nama, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (lgcUseCase *logicUseCaseImpl) GetLogicByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetByEmailJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := lgcUseCase.lgcRepo.GetLogicByEmail(email, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (lgcUseCase *logicUseCaseImpl) GetLogicByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	if gender == "male" {
		gender = "Male"
	}
	if gender == "female" {
		gender = "Female"
	}
	var execution model.ExecutionModel
	execution.Nama = "GetByGenderJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := lgcUseCase.lgcRepo.GetLogicByGender(gender, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (lgcUseCase *logicUseCaseImpl) InsertLogic(lgc []*model.LogicModel, ctx *gin.Context) error {
	return lgcUseCase.lgcRepo.InsertLogic(lgc)
}

func (lgcUseCase *logicUseCaseImpl) EditLogicById(lgc *model.LogicModel) error {
	var execution model.ExecutionModel
	execution.Nama = "UpdateByIDJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	_, _, err := lgcUseCase.lgcRepo.GetLogicById(lgc.ID, 0)
	if err != nil {
		return err
	}
	err = lgcUseCase.lgcRepo.EditLogicById(lgc)
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}

	return nil

}

func (lgcUseCase *logicUseCaseImpl) DeleteLogicById(id string) error {
	var execution model.ExecutionModel
	execution.Nama = "DeleteByIDJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	err := lgcUseCase.lgcRepo.DeleteLogicById(id)
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}
	return nil
}

func (lgcUseCase *logicUseCaseImpl) DeleteAllLogic() error {
	var execution model.ExecutionModel
	execution.Nama = "DeleteAllJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	err := lgcUseCase.lgcRepo.DeleteAllLogic()
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = lgcUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = lgcUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}
	return nil
}

func NewLogicUseCase(lgcRepo repository.LogicRepo, pstsqlRepo repository.PostgreSqlRepo) LogicUseCase {
	return &logicUseCaseImpl{
		lgcRepo:    lgcRepo,
		pstsqlRepo: pstsqlRepo,
	}
}
