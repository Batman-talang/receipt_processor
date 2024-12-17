package ports

import "receipt_processor/internal/models"

type Receipt interface {
	CalculatePoints(r models.Receipt) string
	GetPoints(id string) (*models.Point, bool)
}
