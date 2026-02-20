package handlers

import (
	"net/http"

	"github.com/Nascfn/Go-Bingo.git/logic"
)

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	game := logic.NewGame()
	// Add game to a global list or store
}
