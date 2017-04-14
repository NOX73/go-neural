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
	"strings"
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

// GetClasses returns the classes in the set
func (s *Set) GetClasses() []string {
	classes := make([]string, len(s.ClassToLabel))
	for k, v := range s.ClassToLabel {
		classes[k] = v
	}
	return classes
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
func (s *Set) LoadFromCSV(path string) (bool, error) {
	classNumbers := make(map[string]int)
	classNumber := 0
	f, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("error while open file: %v", path)
	}
	defer f.Close()
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

func (s *Set) LoadFromSVMFile(path string) (bool, error) {
	classNumbers := make(map[string]int)
	classNumber := 0
	highestIndex := scanSamples(path)
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("error while opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m, label, err := problemToMap(line)
		if err != nil {
			return false, fmt.Errorf("error while scanning files: %v", err)
		}
		var sample Sample
		sample.Vector = make([]float64, highestIndex)
		sample.Label = label
		regression, err := strconv.ParseFloat(label, 64)
		if err != nil {
			sample.Value = regression
		}
		if _, ok := classNumbers[sample.Label]; !ok {
			classNumbers[sample.Label] = classNumber
			classNumber++
		}
		for i := 0; i < highestIndex; i++ {
			if val, ok := m[i]; ok {
				sample.Vector[i] = val
			} else {
				sample.Vector[i] = 0.0
			}
		}
		s.Samples = append(s.Samples, sample)
	}
	return true, nil
}

func problemToMap(problem string) (map[int]float64, string, error) {
	sliced := strings.Split(problem, " ")
	m := make(map[int]float64)
	label := sliced[0]
	features := sliced[1:len(sliced)]
	for feature := range features {
		if features[feature] == "" {
			continue
		}
		splitted := strings.Split(features[feature], ":")
		idx, errIdx := strconv.Atoi(splitted[0])
		value, errVal := strconv.ParseFloat(splitted[1], 64)
		if errIdx == nil && errVal == nil {
			m[idx] = value
		}
	}
	return m, label, nil
}

// this function returns the highest index found
func scanSamples(path string) int {
	highest := 0
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error while opening file")
		os.Exit(-1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m, _, err := problemToMap(scanner.Text())
		if err != nil {
			fmt.Printf("error while scanning files: %v", err)
			os.Exit(-1)
		}
		for k := range m {
			if k > highest {
				highest = k
			}
		}
	}
	return highest
}
