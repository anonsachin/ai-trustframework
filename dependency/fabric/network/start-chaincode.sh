#!/bin/bash
export CHAINCODE_ID=$(docker exec cli sh -c 'peer lifecycle chaincode queryinstalled | grep Package | cut -d \  -f 3 | cut -d , -f 1')
export NETWORK_NAME="network_test"
export VER=0.1.3
# if [ "$CHAIN_CODE_STATE" == "run" ]
# then
    echo "The chaioncode state ${CHAIN_CODE_STATE}"
    docker run -d --rm  --network ${NETWORK_NAME}  -e CHAINCODE_SERVER_ADDRESS=trust:7054 -e CHAINCODE_ID=${CHAINCODE_ID} --hostname trust --name trust  ai-trustframework/trustcc-wasm:${VER}
# fi