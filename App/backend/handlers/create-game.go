package handlers

import (
	"net/http"

	"github.com/Nascfn/Go-Bingo.git/logic"
)

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	game := logic.NewGame()

	// game do paul safadinho
}
