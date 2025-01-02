package repository

import (
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/utils"

	"gorm.io/gorm"
)

type PostgreSqlRepo interface {
	InsertPostgreSql(pstsql []*model.LogicModel) error
	GetPostgreSqlById(id string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetPostgreSqlByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetPostgreSqlByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetPostgreSqlByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error)
	GetAllPostgreSql(page int) ([]*model.LogicModel, *model.Pagination, error)
	EditPostgreSqlById(pstsql *model.LogicModel) error
	DeletePostgreSqlById(id string) error
	DeleteAllPostgreSql() error
	MainPostgreSqlExec(pstsql *model.ExecutionModel) error
}

type postgresqlRepoImpl struct {
	db *gorm.DB
}

func (pstsqlRepo *postgresqlRepoImpl) InsertPostgreSql(pstsql []*model.LogicModel) error {
	pstsqlRepo.DeleteAllPostgreSql()
	if err := pstsqlRepo.db.Create(&pstsql).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    111,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (pstsqlRepo *postgresqlRepoImpl) GetPostgreSqlById(id string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var result []*model.LogicModel
	if err := pstsqlRepo.db.Where("id = ?", id).Find(&result).Error; err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    211,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil, nil
}

func (pstsqlRepo *postgresqlRepoImpl) GetPostgreSqlByName(nama string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var result []*model.LogicModel
	if err := pstsqlRepo.db.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+nama+"%", "%"+nama+"%").Find(&result).Error; err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    311,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil, nil
}

func (pstsqlRepo *postgresqlRepoImpl) GetPostgreSqlByEmail(email string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var result []*model.LogicModel
	if err := pstsqlRepo.db.Where("email ILIKE ?", "%"+email+"%").Find(&result).Error; err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    411,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil, nil
}

func (pstsqlRepo *postgresqlRepoImpl) GetPostgreSqlByGender(gender string, page int) ([]*model.LogicModel, *model.Pagination, error) {
	var result []*model.LogicModel
	if err := pstsqlRepo.db.Where("gender = ?", gender).Find(&result).Error; err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    511,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil, nil
}

func (pstsqlRepo *postgresqlRepoImpl) GetAllPostgreSql(page int) ([]*model.LogicModel, *model.Pagination, error) {
	var result []*model.LogicModel
	if err := pstsqlRepo.db.Order("CAST(id AS INTEGER)").Find(&result).Error; err != nil {
		return nil, nil, &utils.AppError{
			ErrorCode:    612,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil, nil
}

func (pstsqlRepo *postgresqlRepoImpl) EditPostgreSqlById(pstsql *model.LogicModel) error {
	if err := pstsqlRepo.db.Model(&model.LogicModel{}).Where("id = ?", pstsql.ID).Updates(pstsql).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    711,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (pstsqlRepo *postgresqlRepoImpl) DeletePostgreSqlById(id string) error {
	if err := pstsqlRepo.db.Where("id = ?", id).Delete(&model.LogicModel{}).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    811,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (pstsqlRepo *postgresqlRepoImpl) DeleteAllPostgreSql() error {
	if err := pstsqlRepo.db.Exec("DELETE FROM logic_models").Error; err != nil {
		return &utils.AppError{
			ErrorCode:    911,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (pstsqlRepo *postgresqlRepoImpl) MainPostgreSqlExec(pstsql *model.ExecutionModel) error {
	var existData []*model.ExecutionModel
	if err := pstsqlRepo.db.Where("nama = ?", pstsql.Nama).Find(&existData).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    1111,
			ErrorMessage: err.Error(),
		}
	}
	if len(existData) != 30 {
		pstsql.Coba = len(existData) + 1
		if err := pstsqlRepo.db.Create(pstsql).Error; err != nil {
			return &utils.AppError{
				ErrorCode:    1112,
				ErrorMessage: err.Error(),
			}
		}
	} else {
		if err := pstsqlRepo.db.Where("nama = ?", pstsql.Nama).Delete(&model.ExecutionModel{}).Error; err != nil {
			return &utils.AppError{
				ErrorCode:    1113,
				ErrorMessage: err.Error(),
			}
		}
		for _, data := range existData {
			if data.Coba > 1 {
				data.Coba--
				if err := pstsqlRepo.db.Save(data).Error; err != nil {
					return &utils.AppError{
						ErrorCode:    1114,
						ErrorMessage: err.Error(),
					}
				}
			}
		}
		pstsql.Coba = 30
		if err := pstsqlRepo.db.Create(pstsql).Error; err != nil {
			return &utils.AppError{
				ErrorCode:    1115,
				ErrorMessage: err.Error(),
			}
		}
	}
	return nil
}

func NewPostgreSqlRepo(db *gorm.DB) PostgreSqlRepo {
	return &postgresqlRepoImpl{
		db: db,
	}
}
