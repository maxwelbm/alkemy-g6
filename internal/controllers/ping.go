package controllers

import (
	"fmt"
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "ping")
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusOK, "ping"))
}
