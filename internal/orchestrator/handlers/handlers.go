package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	expression "github.com/Aspikk/Distributed-Calculator/internal/entities/expression"
	storage "github.com/Aspikk/Distributed-Calculator/internal/storage"
)

// Handlers

func AddExpression(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "something went wrong... (resolving body)", http.StatusInternalServerError)
		return
	}

	exp := expression.New()
	err = json.Unmarshal(body, exp)
	if err != nil {
		http.Error(w, "something went wrong... (resolving body)", http.StatusInternalServerError)
		return
	}

	if exp.AddSpaces(); exp.IsInvalid() {
		http.Error(w, "invalid expression", http.StatusUnprocessableEntity)
		return
	}

	exp.AddTo(storage.DB.Expressions)

	response := fmt.Sprintf("{\"id\": %d}", exp.ID)
	writeResponse(w, http.StatusCreated, []byte(response))
}

func GetExpressions(w http.ResponseWriter, r *http.Request) {
	data, err := storage.DB.MarshallJSON()
	if err != nil {
		http.Error(w, "something went wrong... (stringifying storage)", http.StatusInternalServerError)
		return
	}

	writeResponse(w, http.StatusOK, data)
}

func GetExpressioinById(w http.ResponseWriter, r *http.Request) {
	idStr, ok := getId(r)
	if !ok {
		http.Error(w, "invalid request path", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusNotFound)
		return
	}

	for _, exp := range *storage.DB.Expressions {
		if exp.ID == id {
			buf, err := json.Marshal(exp)
			if err != nil {
				http.Error(w, "something went wrong... (stringifying expression)", http.StatusInternalServerError)
				return
			}

			result := fmt.Sprintf("{\"expression\": %s}", string(buf))
			writeResponse(w, http.StatusOK, []byte(result))
			return
		}
	}

	http.Error(w, "expression not found", http.StatusNotFound)
}

// Help functions

func writeResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(data)
}

func getId(r *http.Request) (string, bool) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 5 {
		return "", false
	}

	return parts[4], true
}
