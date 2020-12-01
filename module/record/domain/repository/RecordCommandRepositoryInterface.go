package repository

import (
	"api-flirt/module/record/domain/entity"
	"api-flirt/module/record/infrastructure/repository/types"
)

// RecordCommandRepositoryInterface holds the implementable methods for record command repository
type RecordCommandRepositoryInterface interface {
	InsertRecord(data types.CreateRecord) (entity.Record, error)
}
