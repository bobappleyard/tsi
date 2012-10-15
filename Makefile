# This makefile is for locally developing the ts package in tandem with this
# one.

export GOPATH=${PWD}
export TSROOT=${PWD}/src/github.com/bobappleyard/ts

all:
	@ go build

run:
	@ ./tsi ${ARGS}

libs:
	@ go install github.com/bobappleyard/ts/ext

test:
	@ ./tsi test.bs

