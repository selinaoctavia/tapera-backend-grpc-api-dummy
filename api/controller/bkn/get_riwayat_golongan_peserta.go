package bkn

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bkn"
)

// GetRiwayatGolonganPeserta godoc
// @Summary Retrieves response get riwayat golongan peserta
// @Produce json
// @Success 200 {object} bkn.RiwayatGolonganPesertaResponse
// @Failure 400 {object} httpx.Error
// @Router /riwayatgolonganpeserta
func (c *Controller) GetRiwayatGolonganPeserta(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bkn.RiwayatGolonganPesertaParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}

	var result *bkn.RiwayatGolonganPesertaResponse

	result, _ = c.srv.GetRiwayatGolonganPeserta(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
