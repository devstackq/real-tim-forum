package handler

import (
	"net/http"
)

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {

	// switch r.Method {

	// case "GET":
	// 	fmt.Println("chat")
	// case "POST":
	// 	// chat,_, _, _, err := GetJsonData(w, r, "chat")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	if chat.Type == "listUser" {
	// 		//get online user & and chated users with current user
	// 		result, err := h.Services.GetListUser(chat.userId)
	// 	}
	// 	if chat.Type == "history" {
	// 		//if no have rromid, create new roomId -> sender, receiver, roomId save db, relation
	// 		//message table, message[id, text, from, who]
	// 		result, err := h.Services.GetHistoryUser(chat.roomId)
	// 	}
	// 	if chat.Type == "sendMessage" {
	// 		result, err := h.Services.sendMessage(chat.From, chat.Who)
	// 		//if no err -> h.Services.showNotify(chat.Who)
	// 	}

	// 	JsonResponse(w, r, http.StatusOK, result)
	// default:
	// 	JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	// }
}

// ednpoint - api/caht/listuser; api/chat/historyuser/; api/chat/message
