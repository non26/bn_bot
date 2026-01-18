package handler

import (
	"bnbot/app/core/bot_opening/handler/req"
	"bnbot/app/core/bot_opening/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type upsertHandler struct {
	service service.IService
}

func NewUpsertHandler(service service.IService) *upsertHandler {
	return &upsertHandler{service: service}
}

func (h *upsertHandler) Handle(c *gin.Context) {
	req := &req.UpsertRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	err := h.service.Upsert(c.Request.Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
