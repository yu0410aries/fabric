package main

import (
	"dapp-base-template/pkg/context"
	"dapp-base-template/pkg/httpserver"
	"log"
	"net/http"
)

func main() {
	router := httpserver.NewRouter()

	context.SharedDataContext = context.New(
		"Admin",
		"Org1MSP",
		"mychannel",
		"chaincode-share-with-me",
		"Org1MSP",
		"Org1MSP",
		"/home/ricky_yu/baas/project/leadtek/222/dapp-share-with-me/ccp.yaml",
		"/home/ricky_yu/baas/project/leadtek/222/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp",
		"/home/ricky_yu/baas/project/leadtek/222/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem",
		false, //discoveryAsLocalhost，這個會把 peer 的路徑打到 locahost 上，實測上好像只有呼叫 initLedge 時候才會生效，但是如果 dapp 跟 fabric 網路不在同一台機器上，這邊應該要給 false，沒仔細研究這個屬性
	)

	err := http.ListenAndServe("0.0.0.0:8501", router)

	if err != nil {
		log.Printf("Failed to http listenerAndServe: %v", err)
	}
}
