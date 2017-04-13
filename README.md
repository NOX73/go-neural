gopher-neural
==============

# Preface
This code was partly taken from github.com/NOX73/go-neural. For the implementation of the core algorithmn all credits belong to NOX73. The fork to gopher-neural was made to pursue the following goals:
* Build a training / testing framework around this algorithm
* Build rich measurement mechanisms to control the training
* Improved I/O functionality for training
* Provide examples for the usage of the library

### Roadmap

After forking the repository from github.com/NOX73/go-neural this will be the roadmap so far (ordered by priority desc):

* Implement rich measurements for the evaluation of the classifier
* Simple data I/O for training / testing and libSVM and csv format
* Establish a learning framework as engine package (using epochs, decays, interraters)
* Provide another page using example projects including data
* Pipelined learning in channels to find the optimum
* Online learning with online evaluation

### Done so far

* Changed I/O handling for JSON models
* Added Sample and Set structure for handling of data sets

# Install

```
  go get github.com/flezzfx/gopher-neural
  go get github.com/flezzfx/gopher-neural/persist
  go get github.com/flezzfx/gopher-neural/learn
  go get github.com/flezzfx/gopher-neural/engine
  go get github.com/flezzfx/gopher-neural/evaluation
```

# Neural Network

Create new network:

```go

  import "github.com/flezzfx/gopher-neural"

  //...

  // Network has 9 enters and 3 layers
  // ( 9 neurons, 9 neurons and 4 neurons).
  // Last layer is network output.
  n := neural.NewNetwork(9, []int{9,9,4})
  // Randomize sypaseses weights
  n.RandomizeSynapses()

  result := n.Calculate([]float64{0,1,0,1,1,1,0,1,0})

```

# Persist network (deprecated)

Save to file:

```go
  import "github.com/flezzfx/gopher-neural/persist"

  persist.ToFile("/path/to/file.json", network)
```

Load from file:

```go
  import "github.com/flezzfx/gopher-neural/persist"

  network := persist.FromFile("/path/to/file.json")
```

# Learning (deprecated)

```go
  import "github.com/flezzfx/gopher-neural/learn"

  var input, idealOutput []float64
  // Learning speed [0..1]
  var speed float64

  learn.Learn(network, in, idealOut, speed)
```

You can get estimate of learning quality:

```go
  e := learn.Evaluation(network, in, idealOut)
```

# Engine  (deprecated)

For concurrent learn, calculate & dump neural network.

```go
	network := neural.NewNetwork(2, []int{2, 2})
	engine := New(network)
	engine.Start()

	engine.Learn([]float64{1, 2}, []float64{3, 3}, 0.1)

	out := engine.Calculate([]float64{1, 2})
```





# Live example (deprecated)

Dirty live example: [https://github.com/NOX73/go-neural-play]
