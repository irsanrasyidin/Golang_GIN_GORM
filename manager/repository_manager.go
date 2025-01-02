package manager

import (
	"sync"
	"Golang_GIN_GORM/repository"
)

type RepoManager interface {
	GetLogicRepo() repository.LogicRepo
	GetPostgreSqlRepo() repository.PostgreSqlRepo
	GetMainRepo() repository.MainRepo
}

type repoManager struct {
	infraManager InfraManager

	lgcRepo    repository.LogicRepo
	pstsqlRepo repository.PostgreSqlRepo
	mainRepo   repository.MainRepo
}

var onceLoadLogicRepo sync.Once
var onceLoadPostgreSqlRepo sync.Once
var onceLoadMainRepo sync.Once

func (rm *repoManager) GetLogicRepo() repository.LogicRepo {
	onceLoadLogicRepo.Do(func() {
		rm.lgcRepo = repository.NewLogicRepo()
	})

	return rm.lgcRepo
}

func (rm *repoManager) GetPostgreSqlRepo() repository.PostgreSqlRepo {
	onceLoadPostgreSqlRepo.Do(func() {
		rm.pstsqlRepo = repository.NewPostgreSqlRepo(rm.infraManager.GetDB())
	})
	return rm.pstsqlRepo
}

func (rm *repoManager) GetMainRepo() repository.MainRepo {
	onceLoadMainRepo.Do(func() {
		rm.mainRepo = repository.NewMainRepo(rm.infraManager.GetDB())
	})
	return rm.mainRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
