package http

import "net/http"

func GenericErrorHandler(w http.ResponseWriter, _ *http.Request, err error) {
	SendErrorMessageWithStatus(w, http.StatusBadRequest, MsgBadRequest, err)
}
