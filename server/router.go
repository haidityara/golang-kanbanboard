package server

import (
	_ "github.com/arfan21/golang-kanbanboard/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
