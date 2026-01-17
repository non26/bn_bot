package handler

import (
	"bnbot/app/bff/bot_register/handler/res"
	"bnbot/app/bff/bot_register/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllHandler struct {
	service service.IService
}

func NewGetAllHandler(service service.IService) *getAllHandler {
	return &getAllHandler{service: service}
}

func (h *getAllHandler) Handle(c *gin.Context) {
	templates, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	res := &res.GetAllResponse{}
	res.FromDomain(templates)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	response.SendGinResponse(http.StatusOK, c)
}
