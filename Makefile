export GOPATH := $(GOPATH):$(PWD)

test: test_neural test_persist

test_persist:
	@( go test go-neural-persist )

test_neural:
	@( go test go-neural )

vim:
	@vim .
