test: test_neural test_persist test_learn test_engine

test_persist:
	@( go test ./persist/ )

test_neural:
	@( go test )

test_learn:
	@( go test ./learn/ )

test_engine:
	@( go test ./engine/ )

goget:
	@( \
		go get github.com/NOX73/go-neural; \
		go get github.com/NOX73/go-neural/persist; \
		go get github.com/NOX73/go-neural/learn; \
		go get github.com/NOX73/go-neural/engine; \
	)

gogetu:
	@( \
		go get -u github.com/NOX73/go-neural; \
		go get -u github.com/NOX73/go-neural/persist; \
		go get -u github.com/NOX73/go-neural/learn; \
		go get -u github.com/NOX73/go-neural/engine; \
	)
