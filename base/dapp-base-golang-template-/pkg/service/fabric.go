package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type FabricContext struct {
	Wallet       *gateway.Wallet
	Gateway      *gateway.Gateway
	Network      *gateway.Network
	Contract     *gateway.Contract
	ContractName string
	FabricDetail
}

type FabricDetail struct {
	Sdk     *fabsdk.FabricSDK
	Channel *channel.Client
}

func (f *FabricContext) BuildDetail(channelID string, userName string, orgName string, ccpPath string) {

	log.Println("============ BuildDetail starts ============")

	sdk, err := fabsdk.New(config.FromFile(ccpPath))

	if err != nil {
		log.Printf("Failed to create new channel client: %s \n ", err)
	}

	log.Printf("sdk : %+v \n ", sdk)

	f.Sdk = sdk

	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))

	ctx, err := channel.New(clientChannelContext)

	if err != nil {
		log.Printf("Failed to create new channel client: %s \n", err)
	}

	log.Printf("ctx : %+v \n ", ctx)

	f.Channel = ctx
}

func (f *FabricContext) Build(
	userName string,
	walletName string,
	channelName string,
	contractName string,
	mspName string,
	ccpPath string,
	credPath string,
	certPath string,
	discoveryAsLocalhost bool,
) {
	f.ContractName = contractName

	log.Println("============ application-golang starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", strconv.FormatBool(discoveryAsLocalhost))
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet(walletName)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists(userName) {
		err = populateWallet(wallet, userName, mspName, credPath, certPath)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, userName),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channelName)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract(contractName)
	f.Wallet = wallet
	f.Gateway = gw
	f.Network = network
	f.Contract = contract

}

func populateWallet(wallet *gateway.Wallet, userName string,
	orgName string, credPath string, certPath string) error {

	log.Println("============ Populating wallet ============")

	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")

	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}

	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}

	keyPath := filepath.Join(keyDir, files[0].Name())

	key, err := ioutil.ReadFile(filepath.Clean(keyPath))

	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(orgName, string(cert), string(key))

	return wallet.Put(userName, identity)

}
