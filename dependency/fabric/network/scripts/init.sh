#!/bin/bash

peer chaincode invoke -o orderer.testnetwork.com:7050 -C testchannel -n wasmcc -c '{"function":"Test","Args":[]}' --tls --cafile $ORDERER_CA_CERT --ordererTLSHostnameOverride orderer.testnetwork.com