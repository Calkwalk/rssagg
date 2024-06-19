package service

import (
	"net/http"

	"github.com/Calkwalk/rssagg/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}
