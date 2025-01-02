package manager

import (
	"sync"
	"Golang_GIN_GORM/usecase"
)

type UsecaseManager interface {
	GetLogicUsecase() usecase.LogicUseCase
	GetPostgreSqlUsecase() usecase.PostgreSqlUseCase
	GetMainUsecase() usecase.MainUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	custUsecase   usecase.LogicUseCase
	pstsqlUsecase usecase.PostgreSqlUseCase
	mainUseCase   usecase.MainUseCase
}

var onceLoadLogicUsecase sync.Once
var onceLoadPostgreSqlUsecase sync.Once
var onceLoadMainUsecase sync.Once

func (um *usecaseManager) GetLogicUsecase() usecase.LogicUseCase {
	onceLoadLogicUsecase.Do(func() {
		um.custUsecase = usecase.NewLogicUseCase(um.repoManager.GetLogicRepo(), um.repoManager.GetPostgreSqlRepo())
	})
	return um.custUsecase
}

func (um *usecaseManager) GetPostgreSqlUsecase() usecase.PostgreSqlUseCase {
	onceLoadPostgreSqlUsecase.Do(func() {
		um.pstsqlUsecase = usecase.NewPostgreSqlUseCase(um.repoManager.GetPostgreSqlRepo(), um.repoManager.GetLogicRepo())
	})
	return um.pstsqlUsecase
}

func (um *usecaseManager) GetMainUsecase() usecase.MainUseCase {
	onceLoadMainUsecase.Do(func() {
		um.mainUseCase = usecase.NewMainUseCase(um.repoManager.GetMainRepo(), um.repoManager.GetLogicRepo(), um.repoManager.GetPostgreSqlRepo())
	})
	return um.mainUseCase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
