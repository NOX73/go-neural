package learn

import (
	"math/rand"
)

type Set struct {
	Samples      []Sample
	ClassToLabel map[int]string
}

type Sample struct {
	Vector      []float64
	Output      []float64
	Label       string
	ClassNumber int
}

func (s *Set) add(vector, output []float64, label string, classNumber int) {
	var sample Sample
	sample.Vector = vector
	sample.Output = output
	sample.Label = label
	sample.ClassNumber = classNumber
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
