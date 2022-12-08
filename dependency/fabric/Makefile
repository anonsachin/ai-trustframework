full-up:
	(cd vault/ ; make up-with-pki)
	sleep 6
	(cd network/ ; make full-clean-start)

full-down:
	(cd network/ ; make teardown)
	(cd vault/ ; make vault-down)