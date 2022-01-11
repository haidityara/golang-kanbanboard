package controllertask

import (
	"github.com/arfan21/golang-kanbanboard/helper"
	"github.com/arfan21/golang-kanbanboard/model/modeltask"
	"github.com/arfan21/golang-kanbanboard/service/servicetask"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerTask interface {
	Create(ctx *gin.Context)
	Gets(ctx *gin.Context)
}

type controller struct {
	srv servicetask.ServiceTask
}

func (c *controller) Gets(ctx *gin.Context) {
	response, err := c.srv.Gets()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (c *controller) Create(ctx *gin.Context) {
	request := new(modeltask.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	request.UserID = ctx.MustGet("user_id").(uint)

	response, err := c.srv.Create(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusCreated, response, nil))
}

func New(srv servicetask.ServiceTask) ControllerTask {
	return &controller{srv: srv}
}
