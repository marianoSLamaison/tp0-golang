package main

import (
	"bytes"
	"client/globals"
	"client/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageRequest struct {
	Mensaje string
}

type respuestaAMensaje struct {
	Mensaje string
}

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargadhttp.NewRequest("POST", url, json.Marshal(g))
	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	bodyReq, erro := json.Marshal(messageRequest{Mensaje: globals.ClientConfig.Mensaje})
	if erro != nil {
		panic("No pudimos ni crear el json!!")
	}
	//creamos un cliente puntero, quien se encarga de mandar nuestros mensajes
	cliente := &http.Client{}
	//cargamos la url que necesitamos
	url := fmt.Sprintf("http://localhost:8080/mensaje")
	//creamos una request con un POST para darle informacion al servidor
	request, erro := http.NewRequest("POST", url, bytes.NewBuffer(bodyReq))
	if erro != nil {
		panic("No pudimos ni crear la request!!")
	}
	//aclaramos que estamos enviando un json en nuestro header de nuestra request
	request.Header.Set("Content-Type", "application/json")
	//enviamos la request al servidor
	respuesta, erro := cliente.Do(request)
	if erro != nil || respuesta.StatusCode != http.StatusOK {
		panic("Huvo un problema con el reuqest!!")
	}
	//decodificamos la respuesta del servidor la cual viene en formato json
	//Esto era por si venia en formato json, pero la respuesta es texto simple
	//a si que no hay por que decodificarlo
	//erro = json.NewDecoder(respuesta.Body).Decode(&fRespuesta)
	//if erro != nil {
	//	panic("Algo paso decodificando la respuesta" + erro.Error())
	//}
	finalResp, _ := io.ReadAll(respuesta.Body)
	log.Println(string(finalResp))

	// leer de la consola el mensaje
	misEntradas := utils.LeerConsola()
	log.Println(misEntradas)
	url = fmt.Sprintf("http://localhost:8080/paquetes")
	// generamos un paquete y lo enviamos al servidor

	//guarda todas las entradas en un json
	//de algun modo go puede leer el tipo de datos que se almacenan y los campos
	//de un struct, lo cual es loquisimo para mi
	bodyReq, erro = json.Marshal(utils.Paquete{Valores: misEntradas})
	if erro != nil {
		panic("Ni No pudimos convertir a json las entradas!!!")
	}

	request, erro = http.NewRequest("POST", url, bytes.NewBuffer(bodyReq))
	if erro != nil {
		panic("No pudimos hacer la request!!")
	}

	request.Header.Set("Content-Type", "application/json")
	respuesta, erro = cliente.Do(request)
	if erro != nil || respuesta.StatusCode != http.StatusOK {
		panic("No pudimos relizar la request!!")
	}
	finalResp, _ = io.ReadAll(respuesta.Body)
	log.Println("El servidor dice " + string(finalResp))

	// utils.GenerarYEnviarPaquete()
}
