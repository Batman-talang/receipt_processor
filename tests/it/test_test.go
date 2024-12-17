package it

import (
	"receipt_processor/internal/models"
	"receipt_processor/internal/ports"

	"receipt_processor/internal/services"
	"receipt_processor/internal/storage"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestReceiptTestSuite(t *testing.T) {
	suite.Run(t, new(ReceiptTestSuite))
}

type ReceiptTestSuite struct {
	suite.Suite
	db  ports.MemoryStore
	svc ports.Receipt
}

func (s *ReceiptTestSuite) SetupSuite() {

	db := storage.New()

	s.db = db

	s.svc = services.New(db)

}

func (s *ReceiptTestSuite) TearDownSuite() {
	s.db.Clear()

}

func (s *ReceiptTestSuite) TestPoints() {

	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}

	id := s.svc.CalculatePoints(receipt)
	s.Require().NotEqual(id, "")
	points, found := s.svc.GetPoints(id)
	s.Require().True(found)
	s.Require().Equal(points.Points, 28)

}
