go-neural
==============

# Install

```
  go get github.com/flezzfx/gopher-neural
  go get github.com/flezzfx/gopher-neural/persist
  go get github.com/flezzfx/gopher-neural/learn
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

# Persist network

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

# Learning

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

# Engine 

For concurrent learn, calculate & dump neural network.

```go
	network := neural.NewNetwork(2, []int{2, 2})
	engine := New(network)
	engine.Start()

	engine.Learn([]float64{1, 2}, []float64{3, 3}, 0.1)

	out := engine.Calculate([]float64{1, 2})
```

# Roadmap

After forking the repository from github.com/NOX73/go-neural this will be the roadmap so far: 
* Establish a learning framework (using epochs, decays, interraters)
* Change panic() handling
* Simple data I/O for training / testing
* Provide another page using example projects including data
* Implement rich measurements for the evaluation of the classifier

# Live example

Dirty live example: [https://github.com/NOX73/go-neural-play]

