package handler

import (
	"bnbot/app/core/bot_register_template/handler/req"
	"bnbot/app/core/bot_register_template/handler/res"
	"bnbot/app/core/bot_register_template/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getTemplateHandler struct {
	service service.IBotBNTemplateService
}

func NewGetTemplateHandler(service service.IBotBNTemplateService) *getTemplateHandler {
	return &getTemplateHandler{service: service}
}

func (h *getTemplateHandler) Handle(c *gin.Context) {
	req := &req.GetTemplateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	template, err := h.service.Get(c.Request.Context(), req.ExchangeId, req.TemplateId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	res := &res.GetTemplateResponse{}
	res.FromDomain(template)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	response.SendGinResponse(http.StatusOK, c)
}
