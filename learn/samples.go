package learn

// TODO (abresk) write tests for samples

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

// Set holds the samples and the output labels
type Set struct {
	Samples      []Sample
	ClassToLabel map[int]string
}

// Sample holds the sample data, value is just used for regression annotation
type Sample struct {
	Vector      []float64
	Output      []float64
	Label       string
	ClassNumber int
	Value       float64
}

// NewSet creates a new set of empty data samples
func NewSet() *Set {
	return &Set{
		Samples:      make([]Sample, 0),
		ClassToLabel: make(map[int]string),
	}
}

func (s *Set) add(vector, output []float64, label string, classNumber int, value float64) {
	var sample Sample
	sample.Vector = vector
	sample.Output = output
	sample.Label = label
	sample.ClassNumber = classNumber
	sample.Value = value
	s.Samples = append(s.Samples, sample)
}

func (s *Set) getLabelFromClass(number int) (string, bool) {
	if val, ok := s.ClassToLabel[number]; ok {
		return val, true
	}
	return "", false
}

func (s *Set) getClassFromLabel(label string) (int, bool) {
	for k, v := range s.ClassToLabel {
		if v == label {
			return k, true
		}
	}
	return -1, false
}

func splitSamples(set *Set, ratio float64) (Set, Set) {
	normalizedRatio := int(ratio * 100.0)
	firstSet := Set{
		Samples:      make([]Sample, 0),
		ClassToLabel: set.ClassToLabel,
	}
	secondSet := Set{
		Samples:      make([]Sample, 0),
		ClassToLabel: set.ClassToLabel,
	}
	for i := range set.Samples {
		if rand.Intn(100) <= normalizedRatio {
			firstSet.Samples = append(firstSet.Samples, set.Samples[i])
		} else {
			secondSet.Samples = append(secondSet.Samples, set.Samples[i])
		}
	}
	return firstSet, secondSet
}

func (s *Set) distributionByLabel(label string) map[string]int {
	dist := make(map[string]int)
	for sample := range s.Samples {
		c := s.Samples[sample].Label
		if _, ok := dist[c]; ok {
			dist[c]++
		} else {
			dist[c] = 1
		}
	}
	return dist
}

func (s *Set) distributionByClassNumber(number int) map[int]int {
	dist := make(map[int]int)
	for sample := range s.Samples {
		c := s.Samples[sample].ClassNumber
		if _, ok := dist[c]; ok {
			dist[c]++
		} else {
			dist[c] = 1
		}
	}
	return dist
}

// where the last dimension is the label
func (s *Set) loadFromCSV(path string) (bool, error) {
	classNumbers := make(map[string]int)
	classNumber := 0
	f, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("error while open file: %v", path)
	}
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		l := len(record)
		var sample Sample
		sample.Vector = make([]float64, l-1)
		sample.Label = record[l-1]
		regression, err := strconv.ParseFloat(record[l-1], 64)
		if err != nil {
			sample.Value = regression
		}
		if _, ok := classNumbers[sample.Label]; !ok {
			classNumbers[sample.Label] = classNumber
			classNumber++
		}
		for value := range record {
			if value < l-1 {
				f, err := strconv.ParseFloat(record[value], 64)
				if err != nil {
					return false, fmt.Errorf("failed to parse float %v with error: %v", record[value], err)
				}
				sample.Vector[value] = f
			}
		}
		s.Samples = append(s.Samples, sample)
	}
	s.createClassToLabel(classNumbers)
	return true, nil
}

func (s *Set) createClassToLabel(mapping map[string]int) {
	s.ClassToLabel = make(map[int]string)
	for k, v := range mapping {
		s.ClassToLabel[v] = k
	}
	for i := range s.Samples {
		s.Samples[i].ClassNumber = mapping[s.Samples[i].Label]
	}
}
