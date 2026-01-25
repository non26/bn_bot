package handler

import (
	"bnbot/app/bff/bot_management/handler/req"
	"bnbot/app/bff/bot_management/handler/res"
	"bnbot/app/bff/bot_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getHandler struct {
	service service.IService
}

func NewGetHandler(service service.IService) *getHandler {
	return &getHandler{service: service}
}

func (h *getHandler) Handle(c *gin.Context) {
	req := &req.GetRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	item, err := h.service.Get(c.Request.Context(), req.BotId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	res := &res.GetResponse{}
	res.FromDomain(item)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	response.SendGinResponse(http.StatusOK, c)
}
