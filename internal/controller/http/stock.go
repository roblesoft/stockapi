package controller

import (
	"net/http"

	"github.com/roblesoft/stockapi/internal/helper"
)

func (s Server) Ping(w http.ResponseWriter, r *http.Request) {
	helper.RenderJSON(w, helper.H{"status": "OK"})
}
