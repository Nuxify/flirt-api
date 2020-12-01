package service

import (
	"context"

	"github.com/segmentio/ksuid"

	"api-flirt/module/record/domain/entity"
	"api-flirt/module/record/domain/repository"
	repositoryTypes "api-flirt/module/record/infrastructure/repository/types"
	"api-flirt/module/record/infrastructure/service/types"
)

// RecordCommandService handles the record command service logic
type RecordCommandService struct {
	repository.RecordCommandRepositoryInterface
}

// CreateRecord create a record
func (service *RecordCommandService) CreateRecord(ctx context.Context, data types.CreateRecord) (entity.Record, error) {
	record := repositoryTypes.CreateRecord{
		ID:   data.ID,
		Data: data.Data,
	}

	// check id if empty create new unique id
	if len(record.ID) == 0 {
		record.ID = generateID()
	}

	res, err := service.RecordCommandRepositoryInterface.InsertRecord(record)
	if err != nil {
		return entity.Record{}, err
	}

	return res, nil
}

// generateID generates unique id
func generateID() string {
	return ksuid.New().String()
}
