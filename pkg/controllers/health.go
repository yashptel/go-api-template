package controllers

import (
	"net/http"

	"github.com/yashptel/go-api-template/pkg/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JSONSuccess(w, http.StatusOK, "OK")
}
