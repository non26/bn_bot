package handler

import (
	"bnbot/app/core/bot_register_template/handler/req"
	"bnbot/app/core/bot_register_template/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type registerTemplateHandler struct {
	service service.IBotBNTemplateService
}

func NewRegisterTemplateHandler(service service.IBotBNTemplateService) *registerTemplateHandler {
	return &registerTemplateHandler{service: service}
}

func (h *registerTemplateHandler) Handle(c *gin.Context) {
	req := &req.UpsertTemplateRequest{}
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
