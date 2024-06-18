package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
