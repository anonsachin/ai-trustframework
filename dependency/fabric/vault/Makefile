pki:
	docker run --rm --network host -v ${PATH_TO_DEPENDENCY}/dependency/fabric/vault/collection:/etc/postman/ \
	-t postman/newman run /etc/postman/vault-setup-Simple-Network.postman_collection.json

vault:
	docker-compose up -d

vault-down:
	docker-compose down -v

up-with-pki:
	(make vault)
	sleep 6
	(make pki)