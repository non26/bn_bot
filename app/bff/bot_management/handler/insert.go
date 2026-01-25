package handler

import (
	"bnbot/app/bff/bot_management/handler/req"
	"bnbot/app/bff/bot_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type insertHandler struct {
	service service.IService
}

func NewInsertHandler(service service.IService) *insertHandler {
	return &insertHandler{service: service}
}

func (h *insertHandler) Handle(c *gin.Context) {
	req := &req.InsertRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	err := h.service.Insert(c.Request.Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
