package httpResponses

import (
	"net/http"

	"github.com/Natannms/go_api/models"
)

type HttpResponseMessage struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	HttpResponseMessage HttpResponseMessage
	User                models.User
	Cookie              http.Cookie
}
