package usecase

import (
	"sort"
	"Golang_GIN_GORM/model"
	"Golang_GIN_GORM/repository"
	"time"
)

type MainUseCase interface {
	MainPostgreSQL(option int) ([]*model.ExecutionModel, []*model.ExecutionResultModel, []*model.ExecutionResultModel, error)
	InsertMain(main []*model.LogicModel) error
	Restart() error
}

type mainUseCaseImpl struct {
	mainRepo   repository.MainRepo
	lgcRepo    repository.LogicRepo
	pstsqlRepo repository.PostgreSqlRepo
}

func (mainUseCase *mainUseCaseImpl) MainPostgreSQL(option int) ([]*model.ExecutionModel, []*model.ExecutionResultModel, []*model.ExecutionResultModel, error) {
	mainData, postgreData, err := mainUseCase.mainRepo.MainPostgreSQL(option)
	if err != nil {
		return nil, nil, nil, err
	}
	mainDaGolang_GIN_GORM, jsonData, err := mainUseCase.mainRepo.MainJSON(option)
	if err != nil {
		return nil, nil, nil, err
	}

	if mainDaGolang_GIN_GORM != nil {
		mainData = append(mainData, mainDaGolang_GIN_GORM...)
		sort.Slice(mainData, func(i, j int) bool {
			// return utils.GetOrder(strings.ToLower(combinedData[i].Nama)) < utils.GetOrder(strings.ToLower(combinedData[j].Nama))
			if mainData[i].Nama == mainData[j].Nama {
				return mainData[i].Coba < mainData[j].Coba
			}
			return mainData[i].Nama < mainData[j].Nama
		})
	}
	// for _, data := range mainData {
	// 	fmt.Printf("Nama: %s, Coba: %d\n", data.Nama, data.Coba)
	// }
	// for _, v := range mainData {
	// 	fmt.Println("main", v)
	// }
	// for _, v := range postgreData {
	// 	fmt.Println("psql", v)
	// }
	// for _, v := range jsonData {
	// 	fmt.Println("json", v)
	// }
	return mainData, postgreData, jsonData, nil
}

func (mainUseCase *mainUseCaseImpl) InsertMain(main []*model.LogicModel) error {
	var execution model.ExecutionModel
	// JSON
	execution.Nama = "InsertJSON"
	in := time.Now()
	execution.Masuk = in.Format("15:04:05")
	err := mainUseCase.lgcRepo.InsertLogic(main)
	if err != nil {
		return err
	}
	out := time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds()) / 1000
	//fmt.Println(execution)
	err = mainUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = mainUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}

	// PostgreSQL
	execution.Nama = "InsertPostgreSQL"
	in = time.Now()
	execution.Masuk = in.Format("15:04:05")
	err = mainUseCase.pstsqlRepo.InsertPostgreSql(main)
	if err != nil {
		return err
	}
	out = time.Now()
	execution.Keluar = out.Format("15:04:05")
	execution.Duration = float64(out.Sub(in).Microseconds()) / 1000
	//fmt.Println(execution)
	err = mainUseCase.lgcRepo.MainLogicExec(&execution)
	if err != nil {
		return err
	}
	execution.Coba = 0
	err = mainUseCase.pstsqlRepo.MainPostgreSqlExec(&execution)
	if err != nil {
		return err
	}
	return nil
}

func (mainUseCase *mainUseCaseImpl) Restart() error {
	return mainUseCase.mainRepo.Restart()
}

func NewMainUseCase(mainRepo repository.MainRepo, lgcRepo repository.LogicRepo, pstsqlRepo repository.PostgreSqlRepo) MainUseCase {
	return &mainUseCaseImpl{
		mainRepo:   mainRepo,
		lgcRepo:    lgcRepo,
		pstsqlRepo: pstsqlRepo,
	}
}
