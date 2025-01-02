package manager

import (
	"fmt"
	"log"
	"sync"
	"Golang_GIN_GORM/config"
	"Golang_GIN_GORM/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	GetDB() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg config.Config
}

var onceLoadDb sync.Once

func (im *infraManager) GetDB() *gorm.DB {
	onceLoadDb.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", im.cfg.Host, im.cfg.Port, im.cfg.User, im.cfg.Password, im.cfg.Name)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Cannot start app, Error when connect to DB ", err.Error())
		}
		im.db = db

		err = db.AutoMigrate(&model.LogicModel{})
		if err != nil {
			log.Fatal(err)
		}

		err = db.AutoMigrate(&model.ExecutionModel{})
		if err != nil {
			log.Fatal(err)
		}
	})

	return im.db
}

func (i *infraManager) DbConn() *gorm.DB {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.GetDB()
	return &infra
}
