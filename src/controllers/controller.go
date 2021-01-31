package controllers

import (
	"commands"
	"encoding/json"
	"net/http"
)

func Get_DHT(rw http.ResponseWriter, req *http.Request) {

	js, err := json.Marshal(commands.Get_TEMP_HUMITY())

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

	rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(200)

	rw.Write(js)
}
