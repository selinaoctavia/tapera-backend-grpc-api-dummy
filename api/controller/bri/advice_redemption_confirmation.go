package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bri"
)

// AdviceRedemptionConfirmation godoc
// @Summary Retrieves response cancel Redemption
// @Produce json
// @Success 200 {object} bri.AdviceRedemptionConfirmationResponse
// @Failure 400 {object} httpx.Error
// @Router /AdviceRedemptionConfirmation
func (c *Controller) AdviceRedemptionConfirmation(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bri.AdviceRedemptionConfirmationParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}

	var result *bri.AdviceRedemptionConfirmationResponse

	result, _ = c.srv.AdviceRedemptionConfirmation(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
