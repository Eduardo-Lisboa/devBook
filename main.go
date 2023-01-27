package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

func init() {

	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
	}

	
}

func main() {

	config.Carregar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
