gopher-neural
==============
![gopher-neural-logo](http://alexander.bre.sk/x/gopher-neural-small.png " The Gopher Neural logo ")

# Quickstart
* See examples here: https://github.com/flezzfx/gopher-neural/tree/master/examples
* Roadmap current version 1.0: https://github.com/flezzfx/gopher-neural/projects/1

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
* Implement rich measurements for the evaluation of regressors

### Roadmap
* Improve the split data set handling by classes (for classification and regression)
* Pipelined learning in channels to find the optimum
* Online learning with online evaluation
* Feature normalizer (auto encoder also for alphanumerical features)


# Install
```
  go get github.com/flezzfx/gopher-neural
  go get github.com/flezzfx/gopher-neural/persist
  go get github.com/flezzfx/gopher-neural/learn
  go get github.com/flezzfx/gopher-neural/engine
  go get github.com/flezzfx/gopher-neural/evaluation
```

# gophers engine
* number of #try (tries)
  * learningRate minus decay if not 0 continue
    * num of #epochs the network sees the training set

* one epoch = one forward pass and one backward pass of all the training examples


learningRate = <number>
n x epoch
	then learningRate - decay

  epochs per learning-decay

## Modes
Gopher-neural can be used to perform classification and regression. This sections helps to set up both modes. In general, you have to take care about the differences between both modes during these parts: read training data from file, start engine, use evaluation modes and perform in production.

### Classification
#### Read training data from file
```go
data := learn.NewSet(neural.Classification)
ok, err := data.LoadFromCSV(dataFile)
```
#### Start engine
```go
e := engine.NewEngine(neural.Classification, []int{hiddenNeurons}, data)
e.SetVerbose(true)
e.Start(engine.CriterionDistance, tries, epochs, trainingSplit, learningRate, decay)
```

#### Use evalation mode
```go
evaluation.GetSummary("name of class1")
evaluation.GetSummary("name of class2")
evaluation.PrintConfusionMatrix()
```

#### Perform in production
```go
x := net.CalculateWinnerLabel(vector)
```

### Regression
Important note: Use regression just with a target value between 0 and 1.

#### Read training data from file
```go
data := learn.NewSet(neural.Regression)
ok, err := data.LoadFromCSV(dataFile)
```
#### Start engine
```go
e := engine.NewEngine(neural.Regression, []int{hiddenNeurons}, data)
e.SetVerbose(true)
e.Start(engine.CriterionDistance, tries, epochs, trainingSplit, learningRate, decay)
```

#### Use evalation mode
```go
evaluation.GetRegressionSummary()
```

#### Perform in production
```go
x := net.Calculate(vector)
```

## Criterions
To let the engine decide for the best model, a few criterias were implemented. They are listed below together with a short regarding their application:

* **CriterionAccuracy** - uses simple accuracy calculation to decide the best model. Not suitable with unbalanced data sets.
* **CriterionBalancedAccuracy** - uses balanced accuracy. Suitable for unbalanced data sets.
* **CriterionFMeasure** - uses F1 score. Suitable for unbalanced data sets.
* **CriterionSimple** - uses simple correct classified divided by all classified samples. Suitable for regression with thresholding.
* **CriterionDistance** - uses distance between ideal output and current output. Suitable for regression.

```go
...
e := engine.NewEngine(neural.Classification, []int{100}, data)
e.Start(engine.CriterionDistance, tries, epochs, trainingSplit, learningRate, decay)
...
```


# Some more basics

## Train a network using engine
```go
import (
	"fmt"

	"github.com/flezzfx/gopher-neural"
	"github.com/flezzfx/gopher-neural/engine"
	"github.com/flezzfx/gopher-neural/learn"
	"github.com/flezzfx/gopher-neural/persist"
)

const (
	dataFile      = "data.csv"
	networkFile   = "network.json"
	tries         = 1
	epochs        = 100 //100
	trainingSplit = 0.7
	learningRate  = 0.6
	decay         = 0.005
  hiddenNeurons = 20
)

func main() {
	data := learn.NewSet(neural.Classification)
	ok, err := data.LoadFromCSV(dataFile)
	if !ok || nil != err {
		fmt.Printf("something went wrong -> %v", err)
	}
	e := engine.NewEngine(neural.Classification, []int{hiddenNeurons}, data)
	e.SetVerbose(true)
	e.Start(engine.CriterionDistance, tries, epochs, trainingSplit, learningRate, decay)
	network, evaluation := e.GetWinner()

	evaluation.GetSummary("name of class1")
	evaluation.GetSummary("name of class2")

	err = persist.ToFile(networkFile, network)
	if err != nil {
		fmt.Printf("error while saving network: %v\n", err)
	}
	network2, err := persist.FromFile(networkFile)
	if err != nil {
		fmt.Printf("error while loading network: %v\n", err)
	}
  // check the network with the first sample
	w := network2.CalculateWinnerLabel(data.Samples[0].Vector)
	fmt.Printf("%v -> %v\n", data.Samples[0].Label, w)

  fmt.Println(" * Confusion Matrix *")
	evaluation.PrintConfusionMatrix()
}

```


## Create simple network for classification

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

# Further ideas

## Rename and batching in learning
* Use term **batch** size = the number of training examples in one forward/backward pass.
* Use term **iterations** = number of passes, each pass using [batch size] number of examples.
* Random application of samples
