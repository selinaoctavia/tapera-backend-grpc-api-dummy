package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bri"
)

// Subscription godoc
func (c *Controller) Subscription(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bri.SubscriptionParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}
	var result *bri.SubscriptionResponse
	result, _ = c.srv.Subscription(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
