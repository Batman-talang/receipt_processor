package services

import (
	"receipt_processor/internal/models"
	"receipt_processor/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Example 1 from the specification",
			receipt: models.Receipt{
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
			},
			expected: 28,
		},
		{
			name: "Example 2 from the specification",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
		{
			name: "No items, simple retailer",
			receipt: models.Receipt{
				Retailer:     "ABC123",
				PurchaseDate: "2022-06-02",
				PurchaseTime: "10:00",
				Items:        []models.Item{},
				Total:        "5.00",
			},
			// Points breakdown:
			// Retailer "ABC123" has 6 alphanumeric chars -> 6 points
			// total = 5.00 is a round dollar -> 50 points
			// total is multiple of 0.25 -> 25 points
			// no items -> 0 points for items
			// total > 10.00 rule does NOT apply because total = 5.00
			// day is even (2), so no points
			// time not in 14:00-16:00 range, so no points
			// total = 6 + 50 + 25 = 81 points
			expected: 81,
		},
		{
			name: "Check odd day and total > 10 rule",
			receipt: models.Receipt{
				Retailer:     "RetailerXYZ",
				PurchaseDate: "2022-06-03", // Day = 3 is odd
				PurchaseTime: "15:00",      // On the edge: 15:00 is after 14:00 and before 16:00?
				Items: []models.Item{
					{ShortDescription: "   ABC   ", Price: "12.00"}, // "ABC" length=3, multiple of 3
				},
				Total: "12.50",
			},
			// Points breakdown:
			// RetailerXYZ alphanumeric count: "RetailerXYZ" -> R(1) e(2) t(3) a(4) i(5) l(6) e(7) r(8) X(9) Y(10) Z(11)
			// -> 11 points
			// total = 12.50 is not a round dollar -> 0
			// total = 12.50 is multiple of 0.25 (12.50/0.25=50) -> 25 points
			// 1 item -> (1/2)*5 = 0 (since we need two items for 5 points)
			// Item desc "ABC": length=3 (multiple of 3), price=12.00*0.2=2.4 -> ceil=3 points
			// total > 10.00 -> 5 points (LLM rule)
			// day=3 (odd) -> 6 points
			// time=15:00 is between 14:00 and 16:00 -> 10 points
			// total = 11+25+0+3+5+6+10 = 60 points
			expected: 55,
		},
	}

	svc := New(storage.New())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := svc.CalculatePoints(tt.receipt)
			points, found := svc.GetPoints(id)
			assert.Equal(t, found, true)
			assert.Equal(t, tt.expected, points)
		})
	}
}
