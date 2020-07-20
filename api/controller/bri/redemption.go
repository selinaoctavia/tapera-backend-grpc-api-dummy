package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bri"
)

// Redemption godoc
func (c *Controller) Redemption(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bri.RedemptionParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}
	var result *bri.RedemptionResponse
	result, _ = c.srv.Redemption(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
