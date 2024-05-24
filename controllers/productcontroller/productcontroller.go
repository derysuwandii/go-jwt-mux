package productcontroller

import (
	"go-jwt-mux/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "Kemeja",
			"stok":         1000,
		},
		{
			"id":           2,
			"nama_product": "Celana",
			"stok":         1000,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
