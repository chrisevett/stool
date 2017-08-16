package game

import (
	"net/http"
	"strconv"

	"github.com/gapi/util"
	"github.com/gorilla/mux"
)

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	g := Game{Id: gameId}

	if err := g.DeleteGame(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
