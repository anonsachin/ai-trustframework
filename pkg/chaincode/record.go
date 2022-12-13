package chaincode

import (
	sig"ai-trustframework/pkg/lamportsig"
	"ai-trustframework/pkg/merkele"
)

type RecordOfExecutionOld struct {
	prediction interface{} // The ouput only
}

type RecordOfExecutionV1 struct {
	prediction interface{} // The ouput only
	explanation string // hash of explanation
}

type RecordOfExecutionV2 struct {
	prediction interface{} // The ouput only
	explanation string // hash of explanation
	model string // hash of model
	input string // hash of input
}

type RecordOfExecutionV3 struct {
	Prediction interface{} // The ouput only
	Explanation string // hash of explanation
	Model string // hash of model
	InputArchive string // hash of input
	OutputArchive string // hash of ouput
}

type RecordOfExecutionV4 struct {
	Prediction interface{} // The ouput only
	Explanation string // hash of explanation
	Model string // hash of model
	InputArchive string // hash of input
	OutputArchive string // hash of ouput
	ExecutionArchive string // hash of execution state
}

type RecordOfExecutionV5 struct {
	Prediction interface{} // The ouput only
	Explanation string // hash of explanation
	Model string // hash of model
	InputArchive string // hash of input
	OutputArchive string // hash of ouput
	ExecutionArchive string // hash of execution state
	Merkele *merkele.Merkele //merkel tree of the elements
}

type RecordOfExecution struct {
	Prediction interface{} // The ouput only
	Explanation string // hash of explanation
	Model string // hash of model
	InputArchive string // hash of input
	OutputArchive string // hash of ouput
	ExecutionArchive string // hash of execution state
	Merkele *merkele.Merkele //merkel tree of the elements
	Signature *sig.LamportSignature //signature and public key 
}