package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
)

// Test godoc
func (c *Controller) Test(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	ext.JSON(http.StatusOK, "Hello")
}
