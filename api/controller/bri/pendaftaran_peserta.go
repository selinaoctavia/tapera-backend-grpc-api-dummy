package bri

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bri"
)

// PendaftaranPeserta godoc
// @Summary Retrieves response pendaftaran peserta
// @Produce json
// @Success 200 {object} pendaftaranpeserta.Result
// @Failure 400 {object} httpx.Error
// @Router /pendaftaranpeserta
func (c *Controller) PendaftaranPeserta(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bri.PendaftaranPesertaParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// var result *bri.PendaftaranPesertaResponse

	result, er := c.srv.PendaftaranPeserta(r.Context(), &param)

	if er != nil {
		ext.JSONerr(http.StatusBadRequest, er.Error())
		return
	}

	ext.JSON(http.StatusOK, result)
}
