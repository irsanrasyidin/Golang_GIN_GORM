package usecase

import (
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/repository"
	"time"
)

type PostgreSqlUseCase interface {
	InsertPostgreSql(pstsql []*model.LogicModel) error
	GetPostgreSqlById(id string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetPostgreSqlByIdNoExec(id string, page int) ([]*model.LogicModel, *model.Pagination, error) 
	GetPostgreSqlByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetPostgreSqlByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetPostgreSqlByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetAllPostgreSql(page int) ([]*model.LogicModel, *model.Pagination, error)
	EditPostgreSqlById(pstsql *model.LogicModel) error
	DeletePostgreSqlById(id string) error
	DeleteAllPostgreSql() error
}

type postgresqlUseCaseImpl struct {
	pstsqlRepo repository.PostgreSqlRepo
	lgcRepo    repository.LogicRepo
}

func (pstsqlUseCase *postgresqlUseCaseImpl) GetAllPostgreSql(page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetAllPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := pstsqlUseCase.pstsqlRepo.GetAllPostgreSql(page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) GetPostgreSqlById(id string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetByIDPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := pstsqlUseCase.pstsqlRepo.GetPostgreSqlById(id, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) GetPostgreSqlByIdNoExec(id string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	data, page2, err := pstsqlUseCase.pstsqlRepo.GetPostgreSqlById(id, page)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) GetPostgreSqlByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetByNamePostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := pstsqlUseCase.pstsqlRepo.GetPostgreSqlByName(nama, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) GetPostgreSqlByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var execution model.ExecutionModel
	execution.Nama = "GetByEmailPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := pstsqlUseCase.pstsqlRepo.GetPostgreSqlByEmail(email, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) GetPostgreSqlByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	if gender == "male" {
		gender = "Male"
	}
	if gender == "female" {
		gender = "Female"
	}
	var execution model.ExecutionModel
	execution.Nama = "GetByGenderPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	data, page2, err := pstsqlUseCase.pstsqlRepo.GetPostgreSqlByGender(gender, page)
	if err != nil {
		return nil, nil, err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return nil, nil, err
	}
	return data, page2, nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) InsertPostgreSql(pstsql []*model.LogicModel) error {
	return pstsqlUseCase.pstsqlRepo.InsertPostgreSql(pstsql)
}

func (pstsqlUseCase *postgresqlUseCaseImpl) EditPostgreSqlById(pstsql *model.LogicModel) error {
	var execution model.ExecutionModel
	execution.Nama = "UpdateByIDPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	lgc, _, err := pstsqlUseCase.pstsqlRepo.GetPostgreSqlById(pstsql.ID, 0)
	if lgc == nil {
		return err
	}
	err = pstsqlUseCase.pstsqlRepo.EditPostgreSqlById(pstsql)
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}
	return nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) DeletePostgreSqlById(id string) error {
	var execution model.ExecutionModel
	execution.Nama = "DeleteByIDPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	err := pstsqlUseCase.pstsqlRepo.DeletePostgreSqlById(id)
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}
	return nil
}

func (pstsqlUseCase *postgresqlUseCaseImpl) DeleteAllPostgreSql() error {
	var execution model.ExecutionModel
	execution.Nama = "DeleteAllPostgreSQL"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	err := pstsqlUseCase.pstsqlRepo.DeleteAllPostgreSql()
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds())/1000
	//fmt.Println(execution)
	err = pstsqlUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = pstsqlUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}
	return nil
}

func NewPostgreSqlUseCase(pstsqlRepo repository.PostgreSqlRepo, lgcRepo repository.LogicRepo) PostgreSqlUseCase {
	return &postgresqlUseCaseImpl{
		pstsqlRepo: pstsqlRepo,
		lgcRepo:    lgcRepo,
	}
}
