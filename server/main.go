package main

import (
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!")
	//Nota, si lo dejas en nil, solo hace que use un multiplexe default
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
