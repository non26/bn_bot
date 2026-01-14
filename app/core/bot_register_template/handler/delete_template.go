package handler

import (
	"bnbot/app/core/bot_register_template/handler/req"
	"bnbot/app/core/bot_register_template/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type deleteTemplateHandler struct {
	service service.IBotBNTemplateService
}

func NewDeleteTemplateHandler(service service.IBotBNTemplateService) *deleteTemplateHandler {
	return &deleteTemplateHandler{service: service}
}

func (h *deleteTemplateHandler) Handle(c *gin.Context) {
	req := &req.DeleteTemplateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	err := h.service.Delete(c.Request.Context(), req.ExchangeId, req.TemplateId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
