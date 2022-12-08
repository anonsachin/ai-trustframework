# Depedency setup
PATH_TO_DEPENDENCY ?= ${PWD}
VER ?= 0.1.0
network-up:
	(cd dependency/fabric ; PATH_TO_DEPENDENCY=${PATH_TO_DEPENDENCY} make full-up)

network-down:
	(cd dependency/fabric ; PATH_TO_DEPENDENCY=${PATH_TO_DEPENDENCY} make full-down)


# Build cc
build-trust-wasm:
	docker build -t ai-trustframework/trustcc-wasm:${VER} --file=docker-build/trust-wasmcc .