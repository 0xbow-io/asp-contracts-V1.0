.PHONY: compile
compile:
	forge b

.PHONY: test
test:
	forge test -vvvv

.PHONY: bindings
bindings: compile
	./bin/compile-bindings.sh

.PHONY: all
all: compile bindings
