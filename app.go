package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	mocksetup "mta-hosting-optimizer/mockSetup"
)

func main() {

	os.Setenv("tempKey", "1")

	mocksetup.InitData() //please remove this function for custom data set

	http.HandleFunc("/get/host", Handler)

	log.Println("Listening at :9000")
	http.ListenAndServe(":9000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		tempKey := os.Getenv("tempKey")
		intVal, err := strconv.Atoi(tempKey)
		if err != nil {
			log.Println("got error while converting environment variable value: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := mocksetup.GetData(intVal)

		bytRes, err := json.Marshal(res)
		if err != nil {
			log.Println("error while marshaling the response: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(bytRes)
		if err != nil {
			log.Println("fail to write the response: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
