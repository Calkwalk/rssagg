package service

import (
	"net/http"

	"github.com/Calkwalk/rssagg/utils"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}
