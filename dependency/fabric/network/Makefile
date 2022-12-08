NETWORK_FILE ?= docker-network.yml
BUILD_FILE ?= docker-build.yml
CHAIN_CODE_STATE ?= "run"

# Pre Start Setup
#----------------
issue-certs:
	./certs.sh

base-setup:
	mkdir channel-artifacts
	mkdir certs

artifacts:
	docker-compose -f docker-build.yml up

artifacts-down:
	docker-compose -f docker-build.yml down -v --remove-orphans

#----------------

# Network setup
#--------------
net-up:
	docker-compose -f ${NETWORK_FILE} up -d

net-down:
	docker-compose -f ${NETWORK_FILE} down -v 

install-channel:
	docker exec cli ./install.sh

init-chaincode:
	docker exec cli ./init.sh

start-chaincode:
	./start-chaincode.sh

stop-chaincode:
	docker rm -f trust

#--------------

# Clean up
#---------
clean-certs:
	sudo rm -r certs/

clean-artifacts:
	sudo rm -r channel-artifacts/

clean: clean-artifacts clean-certs
	(cd scripts/ ; ./cleanup.sh)

# Complete
#---------

full-clean-start: base-setup issue-certs sleep-6 artifacts sleep-2 net-up
	sleep 6
	(make install-channel)
	sleep 6 
	(make start-chaincode)
	sleep 6
	(make init-chaincode)

teardown: artifacts-down stop-chaincode net-down clean

sleep-6:
	sleep 6

sleep-2:
	sleep 2