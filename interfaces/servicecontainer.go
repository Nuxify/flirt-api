/*
|--------------------------------------------------------------------------
| Service Container
|--------------------------------------------------------------------------
|
| This file performs the compiled dependency injection for your middlewares,
| controllers, services, providers, repositories, etc..
|
*/
package interfaces

import (
	"log"
	"os"
	"sync"

	"api-flirt/infrastructures/database/mysql"

	recordRepository "api-flirt/module/record/infrastructure/repository"
	recordService "api-flirt/module/record/infrastructure/service"
	recordREST "api-flirt/module/record/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterRecordRESTCommandController() recordREST.RecordCommandController
	RegisterRecordRESTQueryController() recordREST.RecordQueryController
}

type kernel struct{}

var (
	m              sync.Mutex
	k              *kernel
	containerOnce  sync.Once
	mysqlDBHandler *mysql.MySQLDBHandler
)

//==========================================================================

//================================= REST ===================================
// RegisterRecordRESTCommandController performs dependency injection to the RegisterRecordRESTCommandController
func (k *kernel) RegisterRecordRESTCommandController() recordREST.RecordCommandController {
	service := k.recordCommandServiceContainer()

	controller := recordREST.RecordCommandController{
		RecordCommandServiceInterface: service,
	}

	return controller
}

// RegisterRecordRESTQueryController performs dependency injection to the RegisterRecordRESTQueryController
func (k *kernel) RegisterRecordRESTQueryController() recordREST.RecordQueryController {
	service := k.recordQueryServiceContainer()

	controller := recordREST.RecordQueryController{
		RecordQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) recordCommandServiceContainer() *recordService.RecordCommandService {
	repository := &recordRepository.RecordCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &recordService.RecordCommandService{
		RecordCommandRepositoryInterface: &recordRepository.RecordCommandRepositoryCircuitBreaker{
			RecordCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) recordQueryServiceContainer() *recordService.RecordQueryService {
	repository := &recordRepository.RecordQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &recordService.RecordQueryService{
		RecordQueryRepositoryInterface: &recordRepository.RecordQueryRepositoryCircuitBreaker{
			RecordQueryRepositoryInterface: repository,
		},
	}

	return service
}

func registerHandlers() {
	// create new mysql database connection
	mysqlDBHandler = &mysql.MySQLDBHandler{}
	err := mysqlDBHandler.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("[SERVER] mysql database is not responding %v", err)
	}
}

// ServiceContainer export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	m.Lock()
	defer m.Unlock()

	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}
	return k
}
