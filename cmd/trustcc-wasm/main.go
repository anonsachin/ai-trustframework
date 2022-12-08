package main

import (
	"ai-trustframework/pkg/chaincode"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/sirupsen/logrus"
)

func main(){
	log := logrus.NewEntry(logrus.New())

	// Getting the connection info
	ccid := os.Getenv("CHAINCODE_ID")
	if ccid == "" {
		log.Fatalln("No Chaincode ID")
	} else {
		log.Println("ID : " + ccid)
	}
	add := os.Getenv("CHAINCODE_SERVER_ADDRESS")
	if add == "" {
		log.Fatalln("No Address assigned")
	} else {
		log.Println("ADD : " + add)
	}
	// Setting up chaincode
	log.Info("Setting up the server.")
	cc := chaincode.NewTrustWASMCC(log)
	cc.InitRegistry()

	// setting up server
	server := &shim.ChaincodeServer{
		CCID:    ccid,
		Address: add,
		CC:      cc,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}
	// Starting
	log.Info("Starting the server...")
	err := server.Start()
	if err != nil {
		log.Fatalln("Error starting chaincode server")
	}
}