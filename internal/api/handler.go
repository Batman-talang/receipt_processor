package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt_processor/internal/models"
	"receipt_processor/internal/ports"
)

type Handler struct {
	svs ports.Receipt
}

// ProcessReceipt godoc
//
// @Summary      Process Receipts
// @Description  Takes in a JSON receipt and returns a JSON object with an ID generated by the service. This ID can be used to retrieve the points awarded to the receipt.
// @Tags         Api-Services
// @Accept       json
// @Produce      json
// @Param        receipt  body      models.Receipt  true  "Receipt data"
// @Success      200      {object}  map[string]string  "JSON with the generated receipt ID"
// @Failure      400      {object}  map[string]string  "Invalid JSON"
// @Router       /receipts/process [post]
func (h Handler) ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	id := h.svs.CalculatePoints(receipt)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetPoint godoc
//
// @Summary      Get Points
// @Description  Given a receipt ID, returns the number of points awarded to that receipt.
// @Tags         Api-Services
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Receipt ID"
// @Success      200  {object}  models.Point    "Points awarded"
// @Failure      404  {object}  map[string]string "Receipt not found"
// @Router       /receipts/{id}/points [get]
func (h Handler) GetPoint(c *gin.Context) {
	id := c.Param("id")
	points, found := h.svs.GetPoints(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
