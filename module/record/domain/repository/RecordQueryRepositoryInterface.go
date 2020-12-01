package repository

import (
	"api-flirt/module/record/domain/entity"
)

// RecordQueryRepositoryInterface holds the implementable method for record query repository
type RecordQueryRepositoryInterface interface {
	SelectRecordByID(ID string) (entity.Record, error)
}
