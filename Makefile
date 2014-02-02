test: test_neural test_persist test_lern test_engine

test_persist:
	@( go test ./persist/ )

test_neural:
	@( go test )

test_lern:
	@( go test ./lern/ )

test_engine:
	@( go test ./engine/ )

goget:
	@( \
		go get github.com/NOX73/go-neural; \
		go get github.com/NOX73/go-neural/persist; \
		go get github.com/NOX73/go-neural/lern; \
		go get github.com/NOX73/go-neural/engine; \
	)

gogetu:
	@( \
		go get -u github.com/NOX73/go-neural; \
		go get -u github.com/NOX73/go-neural/persist; \
		go get -u github.com/NOX73/go-neural/lern; \
		go get -u github.com/NOX73/go-neural/engine; \
	)
