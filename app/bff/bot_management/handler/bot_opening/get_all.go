package handler

import (
	res "bnbot/app/bff/bot_management/handler/res/bot_opening"
	"bnbot/app/bff/bot_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllHandler struct {
	service service.IBotOpeningService
}

func NewGetAllHandler(service service.IBotOpeningService) *getAllHandler {
	return &getAllHandler{service: service}
}

func (h *getAllHandler) Handle(c *gin.Context) {

	items, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	res := &res.GetAllResponse{}
	res.FromDomainList(items)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	response.SendGinResponse(http.StatusOK, c)
}
