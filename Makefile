test: test_neural test_persist test_learn test_engine test_evaluation

test_persist:
	@( go test ./persist/ )

test_neural:
	@( go test )

test_learn:
	@( go test ./learn/ )

test_engine:
	@( go test ./engine/ )

test_engine:
	@( go test ./evaluation/ )

goget:
	@( \
		go get github.com/flezzfx/gopher-neural; \
		go get github.com/flezzfx/gopher-neural/persist; \
		go get github.com/flezzfx/gopher-neural/learn; \
		go get github.com/flezzfx/gopher-neural/engine; \
		go get github.com/flezzfx/gopher-neural/evaluation; \
	)

gogetu:
	@( \
		go get github.com/flezzfx/gopher-neural; \
		go get github.com/flezzfx/gopher-neural/persist; \
		go get github.com/flezzfx/gopher-neural/learn; \
		go get github.com/flezzfx/gopher-neural/engine; \
		go get github.com/flezzfx/gopher-neural/evaluation; \
	)
