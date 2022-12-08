package chaincode

import (
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type Action func(stub shim.ChaincodeStubInterface, args []string) pb.Response
type ActionsRegistry map[string]Action

//List of the actions for the chaincode

const (
	Test         = "Test"
)