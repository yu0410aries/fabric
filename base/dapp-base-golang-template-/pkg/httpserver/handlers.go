package httpserver

import (
	"dapp-base-template/pkg/context"
	"log"
	"net/http"
)

func InitLedger(w http.ResponseWriter, r *http.Request) {

	contract := context.SharedDataContext.MyFabric.Contract

	result, err := contract.SubmitTransaction("InitLedger")

	if err != nil {
		log.Printf("Failed to Submit transaction: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))

}
