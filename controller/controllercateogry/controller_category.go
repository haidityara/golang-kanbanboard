package controllercateogry

import (
	"github.com/arfan21/golang-kanbanboard/helper"
	"github.com/arfan21/golang-kanbanboard/model/modelcategory"
	"github.com/arfan21/golang-kanbanboard/service/servicecategory"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerCategory interface {
	Create(ctx *gin.Context)
	Gets(ctx *gin.Context)
}

type Controller struct {
	srv servicecategory.ServiceCategory
}

func (c *Controller) Gets(ctx *gin.Context) {
	resp, err := c.srv.Gets()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, resp, nil))
	return
}

func (c *Controller) Create(ctx *gin.Context) {
	request := new(modelcategory.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	create, err := c.srv.Create(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, create, nil))
}

func New(srv servicecategory.ServiceCategory) ControllerCategory {
	return &Controller{srv: srv}
}
