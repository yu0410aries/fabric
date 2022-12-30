package context

import (
	"dapp-base-template/pkg/service"
	"log"
	"strconv"
)

type GlobalContext struct {
	MyFabric *service.FabricContext
}

func New(
	userName string,
	walletName string,
	channelName string,
	contractName string,
	mspName string,
	orgName string,
	ccpPath string,
	credPath string,
	certPath string,
	discoveryAsLocalhost bool,
) *GlobalContext {
	log.Println("========[ new context ]==========")
	log.Println("userName: " + userName)
	log.Println("walletName: " + walletName)
	log.Println("channelName: " + channelName)
	log.Println("contractName: " + contractName)
	log.Println("orgName: " + mspName)
	log.Println("ccpPath: " + ccpPath)
	log.Println("credPath: " + credPath)
	log.Println("certPath: " + certPath)
	log.Println("discoveryAsLocalhost " + strconv.FormatBool(discoveryAsLocalhost))

	ctx := &service.FabricContext{}
	ctx.Build(userName, walletName, channelName, contractName, mspName, ccpPath, credPath, certPath, discoveryAsLocalhost)
	ctx.BuildDetail(channelName, userName, orgName, ccpPath)
	return &GlobalContext{
		MyFabric: ctx,
	}
}

var SharedDataContext *GlobalContext
