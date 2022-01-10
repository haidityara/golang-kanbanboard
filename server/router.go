package server

import (
	"github.com/arfan21/golang-kanbanboard/controller/controlleruser"
	_ "github.com/arfan21/golang-kanbanboard/docs"
	"github.com/arfan21/golang-kanbanboard/middleware"
	"github.com/arfan21/golang-kanbanboard/repository/repositoryuser"
	"github.com/arfan21/golang-kanbanboard/service/serviceuser"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/update-account", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("/delete-account", middleware.Authorization, ctrlUser.DeleteByID)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
