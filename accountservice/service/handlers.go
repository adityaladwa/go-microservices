package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adityaladwa/go-microservices/accountservice/dbservice"
	"github.com/gorilla/mux"
)

var DBClient dbservice.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountId = mux.Vars(r)["accountId"]
	account, err := DBClient.QueryAccount(accountId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
