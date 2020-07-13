package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bri"
)

// CancelSubscribe godoc
// @Summary Retrieves response cancel subscribe
// @Produce json
// @Success 200 {object} bri.CancelSubscribeResponse
// @Failure 400 {object} httpx.Error
// @Router /cancelsubscribe
func (c *Controller) CancelSubscribe(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bri.CancelSubscribeParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}

	var result *bri.CancelSubscribeResponse

	result, _ = c.srv.CancelSubscribe(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
