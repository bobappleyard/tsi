# This makefile is for locally developing the ts package in tandem with this
# one.

export GOPATH=${PWD}
export TSROOT=${PWD}/src/github.com/bobappleyard/ts

all:
	@ go build

run:
	@ ./tsi

test:
	@ ./tsi test.bs

