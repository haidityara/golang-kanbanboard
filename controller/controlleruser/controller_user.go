package controlleruser

import (
	"net/http"

	"github.com/arfan21/golang-kanbanboard/helper"
	"github.com/arfan21/golang-kanbanboard/model/modeluser"
	"github.com/arfan21/golang-kanbanboard/service/serviceuser"
	"github.com/gin-gonic/gin"
)

type ControllerUser interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv serviceuser.ServiceUser
}

func New(srv serviceuser.ServiceUser) ControllerUser {
	return &controller{srv}
}

func (c *controller) Create(ctx *gin.Context) {
	data := new(modeluser.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Create(*data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusCreated, response, nil))
}

func (c *controller) Login(ctx *gin.Context) {
	data := new(modeluser.RequestLogin)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Login(*data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (c *controller) Update(ctx *gin.Context) {
	data := new(modeluser.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	id := ctx.MustGet("user_id")

	data.ID = id.(uint)

	response, err := c.srv.Update(*data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (c *controller) DeleteByID(ctx *gin.Context) {
	id := ctx.MustGet("user_id")

	err := c.srv.DeleteByID(id.(uint))

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, map[string]interface{}{"message": "your account has been successfully deleted"}, nil))
}
