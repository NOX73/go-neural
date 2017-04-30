package main

import (
	"fmt"

	neural "github.com/flezzfx/gopher-neural"
	"github.com/flezzfx/gopher-neural/engine"
	"github.com/flezzfx/gopher-neural/learn"
	"github.com/flezzfx/gopher-neural/persist"
)

const (
	dataFile            = "winequality-white.csv"
	networkFile         = "network.json"
	tries               = 1
	epochs              = 100
	trainingSplit       = 0.7
	learningRate        = 0.9
	decay               = 0.001
	hiddenNeurons       = 100
	regressionThreshold = 0.005 // helps evaluation to define between wrong or right
)

func main() {
	data := learn.NewSet(neural.Regression)
	ok, err := data.LoadFromCSV(dataFile)
	if !ok || nil != err {
		fmt.Printf("something went wrong -> %v", err)
	}
	e := engine.NewEngine(neural.Regression, []int{hiddenNeurons}, data)
	e.SetRegressionThreshold(regressionThreshold)
	e.SetVerbose(true)
	// here we ware choosing CriterionSimple because we want the regressor that produces the best examples
	e.Start(engine.CriterionSimple, tries, epochs, trainingSplit, learningRate, decay)
	network, evaluation := e.GetWinner()

	// regression evaluation
	evaluation.GetRegressionSummary()

	err = persist.ToFile(networkFile, network)
	if err != nil {
		fmt.Printf("error while saving network: %v\n", err)
	}
	// persisted network
	network2, err := persist.FromFile(networkFile)
	if err != nil {
		fmt.Printf("error while loading network: %v\n", err)
	}

	// some examples
	w := network2.Calculate(data.Samples[0].Vector)
	fmt.Printf("%v -> %v\n", data.Samples[0].Value, w)
	w = network2.Calculate(data.Samples[52].Vector)
	fmt.Printf("%v -> %v\n", data.Samples[52].Value, w)
	w = network2.Calculate(data.Samples[180].Vector)
	fmt.Printf("%v -> %v\n", data.Samples[189].Value, w)
}
