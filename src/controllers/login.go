package controllers

import (
	"api/src/auth"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, errors.New("ao ler o corpo da requisição"))
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoReq, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, errors.New("ao converter o usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, errors.New("ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, errors.New("ao buscar usuário no banco de dados"))
		return
	}

	if err = security.CheckPasswordHash(usuario.Senha, []byte(usuarioSalvoNoBanco.Senha)); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("senha incorreta"))
		return
	}

	token, _ := auth.GenerateToken(usuarioSalvoNoBanco.ID)
	fmt.Println(token)
	w.Write([]byte(token))

}
