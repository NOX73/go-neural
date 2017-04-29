gopher-neural
==============
![gopher-neural-logo](http://alexander.bre.sk/x/gopher-neural-small.png " The Gopher Neural logo ")

# Quickstart
* See examples here: https://github.com/flezzfx/gopher-neural-examples
* Version 1 roadmap: https://github.com/flezzfx/gopher-neural/projects/1
* Vversion 1.1 roadmap: https://github.com/flezzfx/gopher-neural/projects/2 

# Preface
This code was partly taken from github.com/NOX73/go-neural. For the implementation of the core algorithm all credits belong to NOX73. The fork to gopher-neural was made to pursue the following goals:
* Build a training / testing framework around this algorithm
* Build rich measurement mechanisms to control the training
* Improved I/O functionality for training
* Provide examples for the usage of the library

### Done so far
* Changed I/O handling for JSON models
* Added Sample and Set structure for handling of data sets
* Implement rich measurements for the evaluation of the classifier
* Simple data I/O for training / testing and libSVM and csv format
* Added labels to output neurons in network and persist
* Just output label of neuron with most confidence
* Establish a learning framework as engine package (using epochs, decays, interraters)
* Provide another repository using example projects including data
* Confusion matrix handling

### Roadmap
* Improve the split data set handling by classes
* Implement rich measurements for the evaluation of regressors
* Pipelined learning in channels to find the optimum
* Online learning with online evaluation
* Feature normalizer (auto encoder also for alphanumerical features)

### Future ReadMe contents
* Fast training and storage of network with csv and svm format i/o
* Explain the algorithm, the engine and the terms
* Explain the evaluation in short

# Install

```
  go get github.com/flezzfx/gopher-neural
  go get github.com/flezzfx/gopher-neural/persist
  go get github.com/flezzfx/gopher-neural/learn
  go get github.com/flezzfx/gopher-neural/engine
  go get github.com/flezzfx/gopher-neural/evaluation
```

# Explaining engine

one epoch = one forward pass and one backward pass of all the training examples
batch size = the number of training examples in one forward/backward pass. The higher the batch size, the more memory space you'll need.
number of iterations = number of passes, each pass using [batch size] number of examples. To be clear, one pass = one forward pass + one backward pass (we do not count the forward pass and backward pass as two different passes).


learningRate = <number>
n x epoch
	then learningRate - decay

  epochs per learning-decay


# Neural Network

## Create new network

```go

  import "github.com/flezzfx/gopher-neural"
  // Network has 9 enters and 3 layers
  // ( 9 neurons, 9 neurons and 2 neurons).
  // Last layer is network output (2 neurons).
  // For these last neurons we need labels (like: spam, nospam, positive, negative)
  labels := make(map[int]string)
  labels[0] = "positive"
  labels[1] = "negative"
  n := neural.NewNetwork(9, []int{9,9,2}, map[int])
  // Randomize sypaseses weights
  n.RandomizeSynapses()

  // now you can calculate on this network (of course it is not trained yet)
  // (for the training you can use then engine)
  result := n.Calculate([]float64{0,1,0,1,1,1,0,1,0})

```

## Persist network (deprecated)

### Save to file

```go
  import "github.com/flezzfx/gopher-neural/persist"

  err := persist.ToFile("/path/to/file.json", network)
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
```

### Load from file

```go
  import "github.com/flezzfx/gopher-neural/persist"

  network, err := persist.FromFile("/path/to/file.json")
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
```

# Learning

If you want to train the network on your own, you can use the Learn function. This fuction helps you with one learning step through the network. If you want to learn a network
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
