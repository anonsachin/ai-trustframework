# Depedency setup
PATH_TO_DEPENDENCY ?= ${PWD}
network-up:
	(cd dependency/fabric ; PATH_TO_DEPENDENCY=${PATH_TO_DEPENDENCY} make full-up)

network-down:
	(cd dependency/fabric ; PATH_TO_DEPENDENCY=${PATH_TO_DEPENDENCY} make full-down)
