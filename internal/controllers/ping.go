package controllers

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "ping")
}
