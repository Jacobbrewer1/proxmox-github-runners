package http

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Jacobbrewer1/proxmox-github-runners/pkg/codegen/apis/common"
	"github.com/Jacobbrewer1/proxmox-github-runners/pkg/logging"
	"github.com/Jacobbrewer1/proxmox-github-runners/pkg/utils"
)

// NewMessage creates a new Message.
func NewMessage(message string, args ...any) *common.Message {
	var msg string
	if len(args) > 0 {
		msg = fmt.Sprintf(message, args...)
	} else {
		msg = message
	}
	return &common.Message{
		Message: &msg,
	}
}

func SendMessageWithStatus(w http.ResponseWriter, status int, message string, args ...any) {
	msg := NewMessage(message, args...)
	err := Encode(w, status, msg)
	if err != nil {
		slog.Error("Error encoding message", slog.String(logging.KeyError, err.Error()))
	}
}

func SendMessage(w http.ResponseWriter, message string, args ...any) {
	SendMessageWithStatus(w, http.StatusOK, message, args...)
}

func NewErrorMessage(message string, err error, args ...any) *common.ErrorMessage {
	var msg string
	if len(args) > 0 {
		msg = fmt.Sprintf(message, args...)
	} else {
		msg = message
	}
	if err == nil {
		err = errors.New("")
	}
	return &common.ErrorMessage{
		Message: &msg,
		Error:   utils.Ptr(err.Error()),
	}
}

func SendErrorMessageWithStatus(w http.ResponseWriter, status int, message string, err error, args ...any) {
	msg := NewErrorMessage(message, err, args...)
	err = Encode(w, status, msg)
	if err != nil {
		slog.Error("Error encoding error message", slog.String("error", err.Error()))
	}
}

func SendErrorMessage(w http.ResponseWriter, message string, err error, args ...any) {
	SendErrorMessageWithStatus(w, http.StatusInternalServerError, message, err, args...)
}
