package gateway_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/gznrf/go-reader/pkg/models"
	response "github.com/gznrf/go-reader/pkg/utils"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	const op = "gateway.handler.auth.register"
	var input models.RegisterUser
	var userID int64

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.WriteError(w, 400, err)
		log.Printf("error with parsing req at %s", op)
		return
	}

	userID, err := h.services.Authorization.Register(input.Name, input.Email, input.Email)
	if err != nil {
		response.WriteError(w, 400, err)
		log.Printf("error with registration at %s", op)
		return
	}

	response.WriteJson(w, 200, map[string]string{
		"New Users id": strconv.Itoa(int(userID)),
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	const op = "gateway.handler.auth.login"
	var input models.LoginUser
	var token string

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.WriteError(w, 400, err)
		log.Printf("error with parsing req at %s", op)
		return
	}

	token, err := h.services.Authorization.Login(input.Email, input.Password)
	if err != nil {
		response.WriteError(w, 400, err)
		log.Printf("error with registration at %s", op)
		return
	}

	response.WriteJson(w, 200, map[string]string{
		"User token": token,
	})
}
