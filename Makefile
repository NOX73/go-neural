test: test_neural test_persist test_lern

test_persist:
	@( cd persist && go test )

test_neural:
	@( go test )


test_lern:
	@( cd lern go test )
