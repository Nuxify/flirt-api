package repository

import (
	"errors"
	"fmt"
	"strings"

	"gomora/infrastructures/database/mysql/types"
	apiError "gomora/internal/errors"
	"gomora/module/record/domain/entity"
	repositoryTypes "gomora/module/record/infrastructure/repository/types"
)

// RecordCommandRepository handles the record command repository logic
type RecordCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// InsertRecord creates a new record
func (repository *RecordCommandRepository) InsertRecord(data repositoryTypes.CreateRecord) (entity.Record, error) {
	record := entity.Record{
		ID:   data.ID,
		Data: data.Data,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (id, data) VALUES (:id, :data)", record.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, record)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return entity.Record{}, errors.New(apiError.DuplicateRecord)
		}
		return entity.Record{}, errors.New(apiError.DatabaseError)
	}

	return record, nil
}
