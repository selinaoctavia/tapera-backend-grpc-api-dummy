package bkn

import (
	"net/http"

	"tapera.mitraintegrasi/api/util/httpx"
	"tapera.mitraintegrasi/service/bkn"
)

// GetPeserta godoc
// @Summary Retrieves response get peserta
// @Produce json
// @Success 200 {object} bkn.PesertaResponse
// @Failure 400 {object} httpx.Error
// @Router /peserta
func (c *Controller) GetPeserta(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param bkn.PesertaParam
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}

	var result *bkn.PesertaResponse

	result, _ = c.srv.GetPeserta(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
