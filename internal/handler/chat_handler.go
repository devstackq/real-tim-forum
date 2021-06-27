package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chat handler")
}

// ednpoint - api/caht/listuser; api/chat/historyuser/; api/chat/message
