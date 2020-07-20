package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bri"
)

// CancelRedemption godoc
// @Summary Retrieves response cancel Redemption
// @Produce json
// @Success 200 {object} bri.CancelRedemptionResponse
// @Failure 400 {object} httpx.Error
// @Router /cancelredemption
func (c *Controller) CancelRedemption(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bri.CancelRedemptionParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}

	var result *bri.CancelRedemptionResponse

	result, _ = c.srv.CancelRedemption(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
