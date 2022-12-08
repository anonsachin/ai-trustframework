#!/bin/bash

. $PWD/channel.sh

installChannel "testchannel"
anchorPeerUpdate orgMSP

. $PWD/chaincode.sh

prepareChaincode "trust:7054"
installAndApprove
commitChanicode