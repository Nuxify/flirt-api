package application

import (
	"context"

	"api-flirt/module/record/domain/entity"
)

// RecordQueryServiceInterface holds the implementable methods for the record query service
type RecordQueryServiceInterface interface {
	GetRecordByID(ctx context.Context, ID string) (entity.Record, error)
}
