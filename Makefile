GO=go1.19.1
DIR=./demo/cmd/main.go
BIN_OUT=./demo/cmd/demo

run:
	${GO} run ${DIR}

build: 
	${GO} build -o ${BIN_OUT} ${DIR}