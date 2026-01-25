package handler

import (
	"bnbot/app/bff/bot_trade_management/handler/req"
	"bnbot/app/bff/bot_trade_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type newOrderHandler struct {
	service service.IService
}

func NewNewOrderHandler(service service.IService) *newOrderHandler {
	return &newOrderHandler{service: service}
}

func (h *newOrderHandler) Handle(c *gin.Context) {
	req := &req.NewOrderRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	err := h.service.NewOrder(c.Request.Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
