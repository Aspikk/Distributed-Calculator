package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	entities "github.com/Aspikk/Distributed-Calculator/internal/entities/expression"
	"github.com/Aspikk/Distributed-Calculator/internal/storage"
)

var (
	id = 1
)

func AddExpression(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "something went wrong... (resolving body)", http.StatusInternalServerError)
		return
	}

	exp := &entities.Expression{}
	err = json.Unmarshal(body, exp)
	if err != nil {
		http.Error(w, "something went wrong... (resolving body)", http.StatusInternalServerError)
		return
	}

	if exp.RemoveSpaces().IsInvalid() {
		http.Error(w, "invalid expression", http.StatusUnprocessableEntity)
		return
	}
	exp.SetID(&id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := fmt.Sprintf("{\"id\": %d}", id-1)
	w.Write([]byte(response))

	exp.Status = "processing..."

	storage.DB.Expressions = append(storage.DB.Expressions, exp)
}

func GetExpressions(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(storage.DB)
	if err != nil {
		http.Error(w, "something went wrong... (stringifying storage)", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(data)
}

func GetExpressioinById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 5 {
		http.Error(w, "invalid request path", http.StatusNotFound)
		return
	}

	idStr := parts[4]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusNotFound)
		return
	}

	for _, exp := range storage.DB.Expressions {
		if exp.ID == id {
			buf, err := json.Marshal(exp)
			if err != nil {
				http.Error(w, "something went wrong... (stringifying expression)", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write([]byte("{\"expression\": "))
			w.Write(buf)
			w.Write([]byte("}"))

			return
		}
	}

	http.Error(w, "expression not found", http.StatusNotFound)
}
