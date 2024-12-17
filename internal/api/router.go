package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "receipt_processor/docs"
	"receipt_processor/internal/ports"
	"time"
)

func NewRouter(svs ports.Receipt) *gin.Engine {
	h := Handler{
		svs: svs,
	}
	router := gin.Default()
	setupSwagger(router)
	router.Use(CORS())
	router.POST("/receipts/process", h.ProcessReceipt)
	router.GET("/receipts/:id/points", h.GetPoint)

	return router
}

// setupSwagger sets up Swagger documentation
func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func CORS() gin.HandlerFunc {
	c := cors.Config{

		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}
	c.AllowAllOrigins = true
	return cors.New(c)
}
