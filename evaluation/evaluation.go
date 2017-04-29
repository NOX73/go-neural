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
	Confusion       map[string]map[string]int
	Correct         int
	Wrong           int
	OverallDistance float64
}

// NewEvaluation creates a new evaluation object
func NewEvaluation(classes []string) *Evaluation {
	evaluation := &Evaluation{
		Confusion: make(map[string]map[string]int),
	}
	for i := range classes {
		evaluation.Confusion[classes[i]] = make(map[string]int)
		for j := range classes {
			evaluation.Confusion[classes[i]][classes[j]] = 0
		}
	}
	return evaluation
}

// Add adds a new data point to the evaluation
func (e *Evaluation) Add(labeledClass, predictedClass string) {
	if _, ok := e.Confusion[labeledClass][predictedClass]; ok {
		e.Confusion[labeledClass][predictedClass]++
	} else {
		e.Confusion[labeledClass][predictedClass] = 1
	}
	if labeledClass == predictedClass {
		e.Correct++
	} else {
		e.Wrong++
	}
}

// GetTruePositives returns TP
func (e *Evaluation) GetTruePositives(label string) int {
	return e.Confusion[label][label]
}

// GetFalsePositives returns FP
func (e *Evaluation) GetFalsePositives(label string) int {
	s := 0
	for l := range e.Confusion {
		if l != label {
			s += e.Confusion[l][label]
		}
	}
	return s
}

// GetTrueNegatives returns TN
func (e *Evaluation) GetTrueNegatives(label string) int {
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

// GetFalseNegatives returns FNs
func (e *Evaluation) GetFalseNegatives(label string) int {
	s := 0
	for la := range e.Confusion[label] {
		for l := range e.Confusion[la] {
			if l != label && la == label {
				s += e.Confusion[la][l]
			}
		}
	}
	return s
}

// GetPositives TP + FN
func (e *Evaluation) GetPositives(label string) int {
	return e.GetTruePositives(label) + e.GetFalseNegatives(label)
}

// GetNegatives FP + TN
func (e *Evaluation) GetNegatives(label string) int {
	return e.GetFalsePositives(label) + e.GetTrueNegatives(label)
}

// GetAccuracy (TP+TN) / (P+N)
func (e *Evaluation) GetAccuracy(label string) float64 {
	if float64(e.GetPositives(label)+e.GetNegatives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetTruePositives(label)+e.GetTrueNegatives(label)) / float64(e.GetPositives(label)+e.GetNegatives(label))
}

// GetRecall TP/P, TP/(TP + FN)
func (e *Evaluation) GetRecall(label string) float64 {
	if float64(e.GetPositives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetTruePositives(label)) / float64(e.GetPositives(label))
}

// GetSensitivity like recall
func (e *Evaluation) GetSensitivity(label string) float64 {
	return e.GetRecall(label)
}

// GetSpecificity TN / N, TN/(FP+TN)
func (e *Evaluation) GetSpecificity(label string) float64 {
	if float64(e.GetNegatives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetTrueNegatives(label)) / float64(e.GetNegatives(label))
}

// GetPrecision TP/(TP+FP)
func (e *Evaluation) GetPrecision(label string) float64 {
	if float64(e.GetTruePositives(label)+e.GetFalsePositives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetTruePositives(label)) / float64(e.GetTruePositives(label)+e.GetFalsePositives(label))
}

// GetFallout FP / N
func (e *Evaluation) GetFallout(label string) float64 {
	if float64(e.GetNegatives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetFalsePositives(label)) / float64(e.GetNegatives(label))
}

// GetFalsePositiveRate same as fallout
func (e *Evaluation) GetFalsePositiveRate(label string) float64 {
	return e.GetFallout(label)
}

// GetFalseDiscoveryRate FP / (FP+TP)
func (e *Evaluation) GetFalseDiscoveryRate(label string) float64 {
	if float64(e.GetFalsePositives(label)+e.GetTruePositives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetFalsePositives(label)) / float64(e.GetFalsePositives(label)+e.GetTruePositives(label))
}

// GetNegativePredictionValue TN/(TN+FN)
func (e *Evaluation) GetNegativePredictionValue(label string) float64 {
	if float64(e.GetTrueNegatives(label)+e.GetFalseNegatives(label)) == 0.0 {
		return 0.0
	}
	return float64(e.GetTrueNegatives(label)) / float64(e.GetTrueNegatives(label)+e.GetFalseNegatives(label))
}

// GetFMeasure 2TP/(2TP+FP+FN)
func (e *Evaluation) GetFMeasure(label string) float64 {
	if float64(2*e.GetTruePositives(label)+e.GetFalsePositives(label)+e.GetFalseNegatives(label)) == 0.0 {
		return 0.0
	}
	return 2.0 * float64(e.GetTruePositives(label)) / float64(2*e.GetTruePositives(label)+e.GetFalsePositives(label)+e.GetFalseNegatives(label))
}

// GetBalancedAccuracy (TP/P + TN/N) / 2
func (e *Evaluation) GetBalancedAccuracy(label string) float64 {
	var positives, negatives float64
	if float64(e.GetPositives(label)) == 0.0 {
		positives = 0.0
	} else {
		positives = float64(e.GetTruePositives(label)) / float64(e.GetPositives(label))
	}
	if float64(e.GetNegatives(label)) == 0.0 {
		negatives = 0.0
	} else {
		negatives = float64(e.GetTrueNegatives(label)) / float64(e.GetNegatives(label))
	}
	return (positives + negatives) / 2.0
}

// GetOverallBalancedAccuracy calculates for the training evaluation
func (e *Evaluation) GetOverallBalancedAccuracy() float64 {
	classes := float64(len(e.Confusion))
	sum := 0.0
	for k := range e.Confusion {
		sum += e.GetBalancedAccuracy(k)
	}
	return sum / classes
}

// GetOverallAccuracy calculates for the training evaluation
func (e *Evaluation) GetOverallAccuracy() float64 {
	classes := float64(len(e.Confusion))
	sum := 0.0
	for k := range e.Confusion {
		sum += e.GetAccuracy(k)
	}
	return sum / classes
}

// GetOverallFMeasure calculates for the training evaluation
func (e *Evaluation) GetOverallFMeasure() float64 {
	classes := float64(len(e.Confusion))
	sum := 0.0
	for k := range e.Confusion {
		sum += e.GetFMeasure(k)
	}
	return sum / classes
}

// GetInformedness  = Sensitivity + Specificity − 1
func (e *Evaluation) GetInformedness(label string) float64 {
	return e.GetSensitivity(label) + e.GetSpecificity(label) - 1.0
}

// GetMarkedness  = Precision + NegativePredictionValue − 1
func (e *Evaluation) GetMarkedness(label string) float64 {
	return e.GetPrecision(label) + e.GetNegativePredictionValue(label) - 1
}

// AddDistance adds distance between ideal output and output of the network
func (e *Evaluation) AddDistance(n *neural.Network, in, ideal []float64) float64 {
	// This function was part of the former go-neural and moved to this package.
	out := n.Calculate(in)
	var d float64
	for i := range out {
		d += math.Pow(out[i]-ideal[i], 2)
	}
	e.OverallDistance += d / 2.0
	return d / 2.0
}

// GetDistance returns the distance from the evaluation
func (e *Evaluation) GetDistance() float64 {
	return e.OverallDistance / float64(e.Wrong+e.Correct)
}

// GetCorrectRatio returns correct classified samples ratio
func (e *Evaluation) GetCorrectRatio() float64 {
	return float64(e.Correct) / float64(e.Wrong+e.Correct)
}

// PrintConfusionMatrix prints the confusion matrix of the evaluation
func (e *Evaluation) PrintConfusionMatrix() {
	fmt.Printf("\t|")
	for k := range e.Confusion {
		fmt.Printf("%v\t|", k)
	}
	fmt.Print("\n")
	for cl := range e.Confusion {
		fmt.Printf("%v\t|", cl)
		for c := range e.Confusion[cl] {
			fmt.Printf("%v\t|", e.Confusion[cl][c])
		}
		fmt.Printf("\n")
	}

}

// GetSummary returns a summary
func (e *Evaluation) GetSummary(label string) {
	fmt.Printf("summary for class %v\n", label)
	fmt.Printf(" * TP: %v TN: %v FP: %v FN: %v\n", e.GetTruePositives(label), e.GetTrueNegatives(label), e.GetFalsePositives(label), e.GetFalseNegatives(label))
	fmt.Printf(" * Recall/Sensitivity: %v\n", e.GetRecall(label))
	fmt.Printf(" * Precision: %v\n", e.GetPrecision(label))
	fmt.Printf(" * Fallout/FalsePosRate: %v\n", e.GetFallout(label))
	fmt.Printf(" * False Discovey Rate: %v\n", e.GetFalseDiscoveryRate(label))
	fmt.Printf(" * Negative Prediction Rate: %v\n", e.GetNegativePredictionValue(label))
	fmt.Println("--")
	fmt.Printf(" * Accuracy: %v\n", e.GetAccuracy(label))
	fmt.Printf(" * F-Measure: %v\n", e.GetFMeasure(label))
	fmt.Printf(" * Balanced Accuracy: %v\n", e.GetBalancedAccuracy(label))
	fmt.Printf(" * Informedness: %v\n", e.GetInformedness(label))
	fmt.Printf(" * Markedness: %v\n", e.GetMarkedness(label))

}
