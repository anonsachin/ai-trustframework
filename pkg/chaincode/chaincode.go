package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/sirupsen/logrus"
)


type TrustWASMCC struct{
	regisrty ActionsRegistry
	log *logrus.Entry
}

func NewTrustWASMCC (log *logrus.Entry) *TrustWASMCC {
	return &TrustWASMCC{
		log: log,
	}
}

func (t *TrustWASMCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	t.log.Println("Chaincode Initiated.")
	return shim.Success(nil)
}

func (t *TrustWASMCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()

	if funcName == "" {
		return shim.Error("UnExpected function")
	}
	//Getting the action
	function, ok := t.regisrty[funcName]
	if !ok {
		return shim.Error(fmt.Sprintf("Unknown function [%s].", funcName))
	}

	t.log.Printf("Calling : %s", funcName)

	return function(stub, args)
}

func (t *TrustWASMCC) InitRegistry() {
	t.regisrty[Test] = t.Test
}

func (t *TrustWASMCC) Test (stub shim.ChaincodeStubInterface, args []string) pb.Response{
	t.log.Println("Chaincode test successful.")
	return shim.Success([]byte("Test successful."))
}