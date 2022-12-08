#!/bin/bash

# peer chaincode invoke -o orderer:7050 -C testchannel -n cpu -c '{"function":"init","Args":[]}' --tls --cafile $ORDERER_CA_CERT --ordererTLSHostnameOverride orderer.testnetwork.com
peer chaincode invoke -o orderer.testnetwork.com:7050 -C testchannel -n cpu -c '{"function":"InitKeysRegistry","Args":[]}' --tls --cafile $ORDERER_CA_CERT --ordererTLSHostnameOverride orderer.testnetwork.com