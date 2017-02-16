package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-json-restapi/managers"
	"simple-json-restapi/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome dear User!")
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := managers.GetUsers()

	json.NewEncoder(w).Encode(users)
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	switch r.Method {

	case "GET":

		userItem := managers.GetUserById(userId)
		json.NewEncoder(w).Encode(userItem)

	case "DELETE":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method", r.Method),
		})

	case "POST":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method ", r.Method),
		})
	}
}
