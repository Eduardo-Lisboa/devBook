package respostas

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dados)
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}

func JsonMensagem(w http.ResponseWriter, statusCode int, mensagem string) {
	JSON(w, statusCode, struct {
		Mensagem string `json:"mensagem"`
	}{
		Mensagem: mensagem,
	})
}	
