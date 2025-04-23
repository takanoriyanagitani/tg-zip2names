#!/bin/sh

tinygo \
	build \
	-o ./zip2names.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./zip2names.go
