package handler

import (
	"bnbot/app/core/bot_opening/handler/req"
	"bnbot/app/core/bot_opening/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type deleteHandler struct {
	service service.IService
}

func NewDeleteHandler(service service.IService) *deleteHandler {
	return &deleteHandler{service: service}
}

func (h *deleteHandler) Handle(c *gin.Context) {
	req := &req.DeleteRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	err := h.service.Delete(c.Request.Context(), req.BotId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
