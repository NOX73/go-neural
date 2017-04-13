package evaluation

import (
	"fmt"
	"math"

	neural "github.com/flezzfx/gopher-neural"
)

// TODO (abresk) confusion matrix handling
// TODO (abresk) write tests for evaluation

// Evaluation contains all the structures necessary for the evaluation
type Evaluation struct {
	Confusion map[string]map[string]int
}

// NewEvaluation creates a new evaluation object
func NewEvaluation(classes []string) *Evaluation {
	evaluation := &Evaluation{
		Confusion: make(map[string]map[string]int),
	}
	for i := range classes {
		for j := range classes {
			evaluation.Confusion[classes[i]][classes[j]] = 0
		}
	}
	return evaluation
}

func (e *Evaluation) add(labeledClass, predictedClass string) {
	if _, ok := e.Confusion[labeledClass][predictedClass]; ok {
		e.Confusion[labeledClass][predictedClass]++
	} else {
		e.Confusion[labeledClass][predictedClass] = 1
	}
}

func (e *Evaluation) getTruePositives(label string) int {
	return e.Confusion[label][label]
}

func (e *Evaluation) getFalsePositives(label string) int {
	s := 0
	for l := range e.Confusion {
		if l != label {
			s += e.Confusion[l][label]
		}
	}
	return s
}

func (e *Evaluation) getTrueNegatives(label string) int {
	s := 0
	for la := range e.Confusion {
		if la != label {
			for l := range e.Confusion[la] {
				if l != label {
					s += e.Confusion[la][l]
				}
			}
		}
	}
	return s
}

func (e *Evaluation) getFalseNegatives(label string) int {
	s := 0
	for la := range e.Confusion[label] {
		for l := range e.Confusion[la] {
			if l != label {
				s += e.Confusion[la][l]
			}
		}
	}
	return s
}

// TP + FN
func (e *Evaluation) getPositives(label string) int {
	return e.getTruePositives(label) + e.getFalseNegatives(label)
}

// FP + TN
func (e *Evaluation) getNegatives(label string) int {
	return e.getFalsePositives(label) + e.getTrueNegatives(label)
}

// (TP+TN) / (P+N)
func (e *Evaluation) getAccuray(label string) float64 {
	return float64(e.getTruePositives(label)+e.getTrueNegatives(label)) / float64(e.getPositives(label)+e.getNegatives(label))
}

// TP/P, TP/(TP + FN)
func (e *Evaluation) getRecall(label string) float64 {
	return float64(e.getTruePositives(label)) / float64(e.getPositives(label))
}

// like recall
func (e *Evaluation) getSensitivity(label string) float64 {
	return e.getRecall(label)
}

// TN / N, TN/(FP+TN)
func (e *Evaluation) getSpecificity(label string) float64 {
	return float64(e.getTrueNegatives(label)) / float64(e.getNegatives(label))
}

// TP/(TP+FP)
func (e *Evaluation) getPrecision(label string) float64 {
	return float64(e.getTruePositives(label)) / float64(e.getTruePositives(label)+e.getFalsePositives(label))
}

// FP / N
func (e *Evaluation) getFallout(label string) float64 {
	return float64(e.getFalsePositives(label)) / float64(e.getNegatives(label))
}

// same as fallout
func (e *Evaluation) getFalsePositiveRate(label string) float64 {
	return e.getFallout(label)
}

// FP / (FP+TP)
func (e *Evaluation) getFalseDiscoveryRate(label string) float64 {
	return float64(e.getFalsePositives(label)) / float64(e.getFalsePositives(label)+e.getTruePositives(label))
}

// TN/(TN+FN)
func (e *Evaluation) getNegativePredictionValue(label string) float64 {
	return float64(e.getTrueNegatives(label)) / float64(e.getTrueNegatives(label)+e.getFalseNegatives(label))
}

// 2TP/(2TP+FP+FN)
func (e *Evaluation) getFMeasure(label string) float64 {
	return 2.0 * float64(e.getTruePositives(label)) / float64(2*e.getTruePositives(label)+e.getFalsePositives(label)+e.getFalseNegatives(label))
}

// (TP/P + TN/N) / 2
func (e *Evaluation) getBalancedAccuracy(label string) float64 {
	positives := float64(e.getTruePositives(label)) / float64(e.getPositives(label))
	negatives := float64(e.getTrueNegatives(label)) / float64(e.getNegatives(label))
	return (positives + negatives) / 2.0
}

// Informedness = Sensitivity + Specificity − 1
func (e *Evaluation) getInformedness(label string) float64 {
	return e.getSensitivity(label) + e.getSpecificity(label) - 1.0
}

// Markedness = Precision + NegativePredictionValue − 1
func (e *Evaluation) getMarkedness(label string) float64 {
	return e.getPrecision(label) + e.getNegativePredictionValue(label) - 1
}

// Math Evaluation with Least squares method.
func shortEvaluation(n *neural.Network, in, ideal []float64) float64 {
	// This function was part of the former go-neural and moved to this package.
	out := n.Calculate(in)
	var e float64
	for i := range out {
		e += math.Pow(out[i]-ideal[i], 2)
	}
	return e / 2
}

func (e *Evaluation) getSummary(label string) {
	fmt.Printf("summary for class %v\n", label)
	fmt.Printf(" * TP: %v TN: %v FP: %v FN: %v\n", e.getTruePositives(label), e.getTrueNegatives(label), e.getFalsePositives(label), e.getFalseNegatives(label))
	fmt.Printf(" * Recall/Sensitivity: %v\n", e.getRecall(label))
	fmt.Printf(" * Precision: %v\n", e.getPrecision(label))
	fmt.Printf(" * Fallout/FalsePosRate: %v\n", e.getFallout(label))
	fmt.Printf(" * False Discovey Rate: %v\n", e.getFalseDiscoveryRate(label))
	fmt.Printf(" * Negative Prediction Rate: %v\n", e.getNegativePredictionValue(label))
	fmt.Println("--")
	fmt.Printf(" * Accuracy: %v\n", e.getAccuray(label))
	fmt.Printf(" * F-Measure: %v\n", e.getFMeasure(label))
	fmt.Printf(" * Balanced Accuracy: %v\n", e.getBalancedAccuracy(label))
	fmt.Printf(" * Informedness: %v\n", e.getInformedness(label))
	fmt.Printf(" * Markedness: %v\n", e.getMarkedness(label))

}
