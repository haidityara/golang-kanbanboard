package server

import (
	"github.com/arfan21/golang-kanbanboard/controller/controllercateogry"
	"github.com/arfan21/golang-kanbanboard/controller/controllertask"
	"github.com/arfan21/golang-kanbanboard/controller/controlleruser"
	_ "github.com/arfan21/golang-kanbanboard/docs"
	"github.com/arfan21/golang-kanbanboard/middleware"
	"github.com/arfan21/golang-kanbanboard/repository/repositorycategory"
	"github.com/arfan21/golang-kanbanboard/repository/repositorytask"
	"github.com/arfan21/golang-kanbanboard/repository/repositoryuser"
	"github.com/arfan21/golang-kanbanboard/service/servicecategory"
	"github.com/arfan21/golang-kanbanboard/service/servicetask"
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

	repoTask := repositorytask.New(db)
	srvTask := servicetask.New(repoTask)
	ctrlTask := controllertask.New(srvTask)

	repoCategory := repositorycategory.New(db)
	srvCategory := servicecategory.New(repoCategory)
	ctrlCategory := controllercateogry.New(srvCategory)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/update-account", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("/delete-account", middleware.Authorization, ctrlUser.DeleteByID)

	// route task
	r.POST("tasks", middleware.Authorization, ctrlTask.Create)
	r.GET("tasks", middleware.Authorization, ctrlTask.Gets)
	r.PUT("tasks/:taskID", middleware.Authorization, ctrlTask.Update)
	r.PATCH("tasks/update-status/:taskID", middleware.Authorization, ctrlTask.UpdateStatus)
	r.PATCH("tasks/update-category/:taskID", middleware.Authorization, ctrlTask.UpdateCategory)
	r.DELETE("tasks/:taskID", middleware.Authorization, ctrlTask.Delete)

	// route category
	r.POST("categories", middleware.Authorization, ctrlCategory.Create)
	r.GET("categories", middleware.Authorization, ctrlCategory.Gets)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
