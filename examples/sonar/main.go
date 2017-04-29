package main

import (
	"fmt"

	"github.com/flezzfx/gopher-neural/engine"
	"github.com/flezzfx/gopher-neural/learn"
	"github.com/flezzfx/gopher-neural/persist"
)

const (
	dataFile      = "data.csv"
	networkFile   = "network.json"
	tries         = 1
	epochs        = 1 //100
	trainingSplit = 0.7
	learningRate  = 0.4
	decay         = 0.005
)

func main() {
	data := learn.NewSet()
	ok, err := data.LoadFromCSV(dataFile)
	if !ok || nil != err {
		fmt.Printf("something went wrong -> %v", err)
	}
	e := engine.NewEngine([]int{100}, data)
	e.SetVerbose(true)
	e.Start(engine.CriterionDistance, tries, epochs, trainingSplit, learningRate, decay)
	network, evaluation := e.GetWinner()

	evaluation.GetSummary("R")
	fmt.Println()
	evaluation.GetSummary("M")

	err = persist.ToFile(networkFile, network)
	if err != nil {
		fmt.Printf("error while saving network: %v\n", err)
	}

	network2, err := persist.FromFile(networkFile)
	if err != nil {
		fmt.Printf("error while loading network: %v\n", err)
	}

	w := network2.CalculateWinnerLabel(data.Samples[0].Vector)
	fmt.Printf("%v -> %v\n", data.Samples[0].Label, w)
	w = network2.CalculateWinnerLabel(data.Samples[70].Vector)
        fmt.Printf("%v -> %v\n", data.Samples[70].Label, w)
	w = network2.CalculateWinnerLabel(data.Samples[120].Vector)
        fmt.Printf("%v -> %v\n", data.Samples[120].Label, w)

	// print confusion matrix
	fmt.Println(" * Confusion Matrix *")
	evaluation.PrintConfusionMatrix()
}
