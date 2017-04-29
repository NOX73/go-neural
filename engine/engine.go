package engine

import (
	"fmt"
	"math/rand"

	neural "github.com/flezzfx/gopher-neural"
	"github.com/flezzfx/gopher-neural/evaluation"
	"github.com/flezzfx/gopher-neural/learn"
	persist "github.com/flezzfx/gopher-neural/persist"
)

const (
	// Classification engine used for classification
	Classification = 0
	// Regresssion engine used for regression
	Regresssion = 1
	// CriterionAccuracy decides evaluation by accuracy
	CriterionAccuracy = 0
	// CriterionBalancedAccuracy decides evaluation by balanced accuracy
	CriterionBalancedAccuracy = 1
	// CriterionFMeasure decides evaluation by f-measure
	CriterionFMeasure = 2
	// CriterionSimple decides on simple wrong/correct ratio
	CriterionSimple = 3
	// CriterionDistance decieds evaluation by distance to ideal output
	CriterionDistance = 4
	// some output tokens
	runToken   = ","
	epochToken = "."
	tryToken   = "*"
)

// Engine contains every necessary for starting the engine
type Engine struct {
	NetworkInput     int
	NetworkLayer     []int
	NetworkOutput    int
	Data             *learn.Set
	WinnerNetwork    *neural.Network
	WinnerEvaluation evaluation.Evaluation
	Verbose          bool
	Type             int
}

// NewEngine creates a new Engine object
func NewEngine(usage int, hiddenLayer []int, data *learn.Set) *Engine {
	return &Engine{
		NetworkInput:     len(data.Samples[0].Vector),
		NetworkOutput:    len(data.Samples[0].Output),
		NetworkLayer:     hiddenLayer,
		Data:             data,
		WinnerNetwork:    build(len(data.Samples[0].Vector), hiddenLayer, data.ClassToLabel),
		WinnerEvaluation: *evaluation.NewEvaluation(data.GetClasses()),
		Verbose:          false,
		Type:             usage,
	}
}

// SetVerbose set verbose mode default = false
func (e *Engine) SetVerbose(verbose bool) {
	e.Verbose = verbose
}

// GetWinner returns the winner network from training
func (e *Engine) GetWinner() (*neural.Network, *evaluation.Evaluation) {
	return e.WinnerNetwork, &e.WinnerEvaluation
}

// Start takes the paramter to start the engine and run it
func (e *Engine) Start(criterion, tries, epochs int, trainingSplit, startLearning, decay float64) {
	network := build(e.NetworkInput, e.NetworkLayer, e.Data.ClassToLabel)
	training, validation := split(e.Data, trainingSplit)
	for try := 0; try < tries; try++ {
		learning := startLearning
		if e.Verbose {
			fmt.Printf("\n> start try %v. training / test: %v / %v (%v)\n", (try + 1), len(training.Samples), len(validation.Samples), trainingSplit)
		}
		for ; learning > 0.0; learning -= decay {
			train(network, training, learning, epochs)
			evaluation := evaluate(network, validation, training)
			if compare(criterion, &e.WinnerEvaluation, evaluation) {
				e.WinnerNetwork = copy(network)
				e.WinnerEvaluation = *evaluation
				if e.Verbose {
					print(&e.WinnerEvaluation)
				}
			}
		}
		if e.Verbose {
			fmt.Print(tryToken + "\n")
		}
	}
}

func print(e *evaluation.Evaluation) {
	fmt.Printf("\n [Best] acc: %v  / bacc: %v / f1: %v / correct: %v / distance: %v \n", e.GetOverallAccuracy(), e.GetOverallBalancedAccuracy(), e.GetOverallFMeasure(), e.GetCorrectRatio(), e.GetDistance())
}

func build(input int, hidden []int, labels map[int]string) *neural.Network {
	hidden = append(hidden, len(labels))
	network := neural.NewNetwork(input, hidden, labels)
	network.RandomizeSynapses()
	return network
}

func split(set *learn.Set, ratio float64) (*learn.Set, *learn.Set) {
	multiplier := 100
	normalizedRatio := int(ratio * float64(multiplier))
	var training, evaluation learn.Set
	training.ClassToLabel = set.ClassToLabel
	evaluation.ClassToLabel = set.ClassToLabel
	for i := range set.Samples {
		if rand.Intn(multiplier) <= normalizedRatio {
			training.Samples = append(training.Samples, set.Samples[i])
		} else {
			evaluation.Samples = append(evaluation.Samples, set.Samples[i])
		}
	}
	return &training, &evaluation
}

func train(network *neural.Network, data *learn.Set, learning float64, epochs int) {
	for e := 0; e < epochs; e++ {
		for sample := range data.Samples {
			learn.Learn(network, data.Samples[sample].Vector, data.Samples[sample].Output, learning)
		}
		fmt.Print(epochToken)
	}
	fmt.Print(runToken)
}

func evaluate(network *neural.Network, test *learn.Set, train *learn.Set) *evaluation.Evaluation {
	evaluation := evaluation.NewEvaluation(train.GetClasses())
	for sample := range test.Samples {
		evaluation.AddDistance(network, test.Samples[sample].Vector, test.Samples[sample].Output)
		winner := network.CalculateWinnerLabel(test.Samples[sample].Vector)
		evaluation.Add(test.Samples[sample].Label, winner)
	}
	return evaluation
}

func compare(criterion int, current *evaluation.Evaluation, try *evaluation.Evaluation) bool {
	if current.Correct+current.Wrong == 0 {
		return true
	}
	switch criterion {
	case CriterionAccuracy:
		if current.GetOverallAccuracy() < try.GetOverallAccuracy() {
			return true
		}
	case CriterionBalancedAccuracy:
		if current.GetOverallBalancedAccuracy() < try.GetOverallBalancedAccuracy() {
			return true
		}
	case CriterionFMeasure:
		if current.GetOverallFMeasure() < try.GetOverallFMeasure() {
			return true
		}
	case CriterionSimple:
		if current.GetCorrectRatio() < try.GetCorrectRatio() {
			return true
		}
	case CriterionDistance:
		if current.GetDistance() > try.GetDistance() {
			return true
		}
	}
	return false
}

func copy(from *neural.Network) *neural.Network {
	return persist.FromDump(persist.ToDump(from))
}
