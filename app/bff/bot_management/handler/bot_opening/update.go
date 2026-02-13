package handler

import (
	req "bnbot/app/bff/bot_management/handler/req/bot_opening"
	"bnbot/app/bff/bot_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type updateHandler struct {
	service service.IBotOpeningService
}

func NewUpdateHandler(service service.IBotOpeningService) *updateHandler {
	return &updateHandler{service: service}
}

func (h *updateHandler) Handle(c *gin.Context) {
	req := &req.UpsertRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	err := h.service.Update(c.Request.Context(), req.ToDomain())
	if err != nil {
		if err.Error() == appresponse.BOTNOTFOUNDCODE {
			response := appresponse.NewAppResponse(
				appresponse.BOTNOTFOUNDCODE,
				appresponse.BOTMAPPING[appresponse.BOTNOTFOUNDCODE],
				nil)
			response.SendGinResponse(http.StatusOK, c)
			return
		}
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
