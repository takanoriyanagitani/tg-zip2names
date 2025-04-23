#!/bin/sh

inputz="./sample.d/input.zip"

geninput(){
	echo generating input zip file...

	mkdir -p sample.d

	echo hw1 > ./sample.d/hw1.txt
	echo hw2 > ./sample.d/hw2.txt

	ls ./sample.d/*.txt |
		zip \
			-0 \
			-@ \
			-T \
			-v \
			-o \
			"${inputz}"
}

test -f "${inputz}" || geninput

wazero \
	run \
	-mount "${PWD}/sample.d:/guest.d:ro" \
	-env ENV_ZIP_FILENAME=/guest.d/input.zip \
	./zip2names.wasm
