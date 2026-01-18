package handler

import (
	"bnbot/app/bff/market_data/handler/req"
	"bnbot/app/bff/market_data/handler/res"
	"bnbot/app/bff/market_data/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getKlineHandler struct {
	service service.IService
}

func NewGetKlineHandler(service service.IService) *getKlineHandler {
	return &getKlineHandler{service: service}
}

func (h *getKlineHandler) Handle(c *gin.Context) {
	req := &req.KlineRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	klines, err := h.service.GetKline(c.Request.Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	res := &res.KlineResponse{}
	res.FromDomainList(klines)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	response.SendGinResponse(http.StatusOK, c)
}
