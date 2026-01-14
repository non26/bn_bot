package handler

import (
	"bnbot/app/core/bot_register_template/handler/res"
	"bnbot/app/core/bot_register_template/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllTemplateHandler struct {
	service service.IBotBNTemplateService
}

func NewGetAllTemplateHandler(service service.IBotBNTemplateService) *getAllTemplateHandler {
	return &getAllTemplateHandler{service: service}
}

func (h *getAllTemplateHandler) Handle(c *gin.Context) {
	templates, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}

	res := &res.GetAllTemplateResponse{}
	res.FromDomain(templates)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	response.SendGinResponse(http.StatusOK, c)
}
