go-neural
==============

# Install

```
  go get github.com/NOX73/go-neural
  go get github.com/NOX73/go-neural/persist
  go get github.com/NOX73/go-neural/lern
```

# Neural Network

Create new network:

```go

  import "github.com/NOX73/go-neural"

  //...

  // Network has 9 enters and 3 layers ( 9 neurons, 9 neurons and 4 neurons). Last layer is network output.
  n := neural.NewNetwork(9, []int{9,9,4})
  // Randomize sypaseses weights
  n.RandomizeSynapses()
  
  result := n.Calucate([]float64{0,1,0,1,1,1,0,1,0})
  
```

# Persist network

Save to file:

```go
  import "github.com/NOX73/gr-neural-persist"

  persist.ToFile("/path/to/file.json", network)
```

Load from file:

```go
  import "github.com/NOX73/gr-neural-persist"

  network := persist.FromFile("/path/to/file.json")
```

# Lerning

```go
  import "github.com/NOX73/gr-neural-lern"

  var input, idealOutput []float64
  // Lerning speed [0..1]
  var speed float64

  lern.Lern(network, in, idealOut, speed)
```

You can get estimate of lerning quality:

```go
  e := lern.Evaluation(network, in, idealOut)
```

# Live example

Dirty live example: [https://github.com/NOX73/go-neural-play]

